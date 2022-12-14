package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"server-api/entity"
)

func GetProducts() []entity.Product {
	productsInByte, err := os.ReadFile("database/products.json")
	if err != nil {
		fmt.Println(err)

	}

	var products []entity.Product
	err = json.Unmarshal(productsInByte, &products)
	if err != nil {
		fmt.Println(err)
	}

	return products
}

// Middleware
func FilterRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(entity.Result{
				Message: "method not allowed",
				Result:  []entity.Product{},
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Products() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(entity.Result{
			Message: "success",
			Result:  GetProducts(),
		})
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/products", FilterRequest(Products()))

	fmt.Println("starting web server in http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
