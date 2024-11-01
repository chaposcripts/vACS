package main

// TODO сука а че такое БД

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var PORT string = "9191"
var BASE_URL string = "/api/v1/vAcs/library"

type Item struct {
	Name      string     `json:"name"`
	Author    string     `json:"author"`
	SecretKey string     `json:"secretKey"`
	Model     uint16     `json:"model"`
	Bone      uint8      `json:"bone"`
	Position  [3]float64 `json:"position"`
	Rotation  [3]float64 `json:"rotation"`
	Scale     [3]float64 `json:"scale"`
}

type User struct {
	Name     string
	Password string
	Token    string
}

var items []Item
var testItem Item = Item{
	Name:      "Test item",
	Author:    "chapo",
	SecretKey: "123123qwe",
	Model:     321,
	Bone:      0,
	Position:  [3]float64{0, 0, 0},
	Rotation:  [3]float64{1, 1, 1},
	Scale:     [3]float64{2, 2, 2},
}

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var item Item
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	fmt.Println(item)
	items = append(items, item)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("name")
	secretKey := r.URL.Query().Get("key")
	if itemName == "" || secretKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for index, item := range items {
		if item.Name == itemName && item.SecretKey == secretKey {
			items = append(items[:index], items[index+1:]...)
			w.Write([]byte(string("OK")))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func listRequestHandler(w http.ResponseWriter, r *http.Request) {
	jsonString, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(jsonString)))
}

func main() {
	items = append(items, testItem)
	http.HandleFunc(BASE_URL+"/add", addItemHandler)
	http.HandleFunc(BASE_URL+"/delete", deleteItemHandler)
	http.HandleFunc(BASE_URL+"/get-list", listRequestHandler)
	http.ListenAndServe(":"+PORT, nil)
	fmt.Println("Server started!")
}
