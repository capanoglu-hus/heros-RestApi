package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Hero modeli oluşturduk
type Hero struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Hero yapısını herosa atadık
var heros []Hero

func main() {
	// routerı başlatmak için
	r := mux.NewRouter()

	// veri tabanında da yapılabilir
	heros = append(heros, Hero{ID: "1", Name: "Harry Potter", Title: "Harry Potter ve Felsefe Taşı", Author: &Author{Firstname: "J. K. ", Lastname: "Rowling"}})
	heros = append(heros, Hero{ID: "2", Name: "Percy Jackson", Title: "Şimşek Hırsızı", Author: &Author{Firstname: "Rick", Lastname: " Riordan"}})
	heros = append(heros, Hero{ID: "3", Name: "JON SNOW", Title: "Taht oyunları ", Author: &Author{Firstname: "George R.R.", Lastname: " Martin"}})
	heros = append(heros, Hero{ID: "4", Name: "Gandalf ", Title: "Yüzüklerin Efendisi", Author: &Author{Firstname: "J. R. R. ", Lastname: " Tolkien"}})

	r.HandleFunc("/", handle).Methods("GET")
	r.HandleFunc("/api/heros", getHeros).Methods("GET")
	r.HandleFunc("/api/heros/{id}", getHero).Methods("GET")
	r.HandleFunc("/api/heros", createHeros).Methods("POST")
	r.HandleFunc("/api/heros/{id}", updateHeros).Methods("PUT")
	r.HandleFunc("/api/heros/{id}", deleteHeros).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// cevabı işliyor
func handle(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(heros)
}

// get işlemi hepsini gösteriyor
func getHeros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(heros)
}

// id belirli olanı get işlemi için tek başına gösteriyor
func getHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r) // Get params
	fmt.Println(r)
	// Loop through the books and find with ID
	for _, item := range heros {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Hero{})
}

// post işlemi için
func createHeros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var book Hero
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // rastgele id ataması yapılıyor
	heros = append(heros, book)
	json.NewEncoder(w).Encode(book)
}

// id olarak güncelleme
func updateHeros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range heros {
		if item.ID == params["id"] {
			heros = append(heros[:index], heros[index+1:]...)
			var hero Hero
			_ = json.NewDecoder(r.Body).Decode(&hero)
			hero.ID = params["id"]
			heros = append(heros, hero)
			json.NewEncoder(w).Encode(hero)
			return
		}
	}

	json.NewEncoder(w).Encode(heros)
}

// delete işlemi id olarak
func deleteHeros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range heros {
		if item.ID == params["id"] {
			heros = append(heros[:index], heros[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(heros)
}
