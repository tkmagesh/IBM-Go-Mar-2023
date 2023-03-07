package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products []Product = []Product{
	{Id: 101, Name: "Pen", Cost: 5},
	{Id: 102, Name: "Pencil", Cost: 2},
	{Id: 103, Name: "Marker", Cost: 15},
}

type AppServer struct {
}

/* http.Handler interface implementation */
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
