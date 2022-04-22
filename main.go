package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Item struct {
	UID string `json:"uid"`
	Name string `json:"name"`
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is home page.")
}

func getInventory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("This is getInventory() endpoint")
	json.NewEncoder(w).Encode(inventory)
}

func createInventory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	inventory = append(inventory, item)
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_deleteItemAtUID(params["uid"])
	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUID(uid string){
	for index, item := range inventory{
		if item.UID == uid{
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)
	//deletes the item at uid
	_deleteItemAtUID(params["uid"])
	//then updates that item with new data
	inventory = append(inventory, item)
	json.NewEncoder(w).Encode(inventory)
}

func handleRequests(){
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/inventory", getInventory).Methods("GET")
	r.HandleFunc("/inventory", createInventory).Methods("POST")
	r.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	r.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe("", r))
}

func main(){
	inventory = append(inventory, Item{
		UID: "0",
		Name: "Bread",
		Desc: "Fresh Bread",
		Price: 1.23,
	})

	inventory = append(inventory, Item{
		UID: "1",
		Name: "Chocolate",
		Desc: "A bar of black chocolate with nuts",
		Price: 2.0,
	})
	handleRequests()
}
