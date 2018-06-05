package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PROD OMIT
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// END OMIT

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, products)
}

// GET OMIT
func getProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	// database code goes here
	for _, item := range products {
		if item.ID == id {
			respondWithJSON(w, http.StatusOK, item)
			return
		}
	}

	respondWithJSON(w, http.StatusNotFound, "Product not found")
}

// END OMIT

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	for index, item := range products {
		if item.ID == id {
			products = append(products[:index], products[index+1:]...)
			respondWithJSON(w, http.StatusOK, item)
			return
		}
	}

	respondWithJSON(w, http.StatusNotFound, "Product not found")
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &p); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request. Error:"+err.Error())
		return
	}
	defer r.Body.Close()
	products = append(products, p)
	respondWithJSON(w, http.StatusCreated, p)
}

// RESPOND OMIT
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// END OMIT

func main() {
	products = append(products,
		Product{ID: "cof", Name: "Coffee", Price: 1.5},
		Product{ID: "tea", Name: "Tea", Price: 1.2},
	)
	// ROUTE OMIT
	r := mux.NewRouter()

	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/product/{id}", deleteProduct).Methods("DELETE")
	r.HandleFunc("/api/product/", addProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
	// END OMIT
}
