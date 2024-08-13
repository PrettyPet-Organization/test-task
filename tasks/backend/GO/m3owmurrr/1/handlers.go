package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func createJSON(w *http.ResponseWriter, v any) {
	JSON, err := json.Marshal(v)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(JSON)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Thanks for the request, but there's nothing here!\n\n")
	t := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "The current time is: %s\n", t)
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: POST", r.URL.Path, "from", r.Host)

	type reqProduct struct {
		Name    string  `json:"name"`
		Price   float64 `json:"price"`
		Country string  `json:"country"`
		Count   int     `json:"count"`
	}

	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		log.Println("expected application/json")
		http.Error(w, "expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	var rp reqProduct
	d := json.NewDecoder(r.Body)
	err = d.Decode(&rp)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data", http.StatusBadRequest)
		return
	}

	id, err := store.AddProduct(rp.Name, rp.Country, rp.Price, rp.Count)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	type ID struct {
		id int
	}

	w.WriteHeader(http.StatusCreated)
	createJSON(&w, ID{id})
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: GET", r.URL.Path, "from", r.Host)

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	product, err := store.GetProduct(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	createJSON(&w, product)
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: GET", r.URL.Path, "from", r.Host)

	products, err := store.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createJSON(&w, products)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: DELETE", r.URL.Path, "from", r.Host)

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := store.DeleteProduct(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("not found id:%d", id), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}

func DeleteAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: DELETE", r.URL.Path, "from", r.Host)
	err := store.DeleteAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ChangeProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: PUT", r.URL.Path, "from", r.Host)

	type reqProduct struct {
		Name    string  `json:"name"`
		Price   float64 `json:"price"`
		Country string  `json:"country"`
		Count   int     `json:"count"`
	}

	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		log.Println(err)
		http.Error(w, "expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	var rp reqProduct
	d := json.NewDecoder(r.Body)
	err = d.Decode(&rp)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err = store.ChangeProduct(id, rp.Name, rp.Country, rp.Price, rp.Count)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("not found id:%d", id), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}

func BuyProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: PUT", r.URL.Path, "from", r.Host)

	type purchase struct {
		Count int `json:"count"`
	}

	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		log.Println(err)
		http.Error(w, "expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	var p purchase
	d := json.NewDecoder(r.Body)
	err = d.Decode(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "wrong data", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err = store.BuyProduct(id, p.Count)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("not found id:%d", id), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}
