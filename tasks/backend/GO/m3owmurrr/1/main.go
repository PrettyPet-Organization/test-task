// Заказчик знает как пользоваться Docker. У него есть готовый front. Ему нужно от тебя API на GO.
// Он открыл маленький склад, занимающийся продажей товаров из Казахстана и Китая. Помоги человеку. Реализация на твоё усмотрение.
package main

import (
	"fmt"
	"log"
	"net/http"

	"products"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var store *products.ProductStore
var err error

const PORT = ":1234"

func main() {
	store, err = products.NewStore(10, 10, 30)
	if err != nil {
		log.Println(err)
		return
	}
	defer store.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", DefaultHandler)
	router.HandleFunc("/products", AddProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/products", GetAllProductsHandler).Methods(http.MethodGet)
	router.HandleFunc("/products/{id:[0-9]+}", GetProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", DeleteAllProductsHandler).Methods(http.MethodDelete)
	router.HandleFunc("/products/{id:[0-9]+}", DeleteProductHandler).Methods(http.MethodDelete)
	router.HandleFunc("/products/{id:[0-9]+}", ChangeProductHandler).Methods(http.MethodPut)
	router.HandleFunc("/products/{id:[0-9]+}/purchase", BuyProductHandler).Methods(http.MethodPut)

	fmt.Println("Ready to serve at", PORT)
	err = http.ListenAndServe("localhost"+PORT, router)
	if err != nil {
		log.Println(err)
		return
	}
}
