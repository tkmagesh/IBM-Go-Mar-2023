package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
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
	reqCount int
	handlers map[string]func(http.ResponseWriter, *http.Request)
	sync.Mutex
}

func (appServer *AppServer) Register(pattern string, handlerFn func(http.ResponseWriter, *http.Request)) {
	appServer.handlers[pattern] = handlerFn
}

/* http.Handler interface implementation */
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	appServer.Lock()
	{
		appServer.reqCount++
	}
	appServer.Unlock()
	// fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
	if handlerFn, exists := appServer.handlers[r.URL.Path]; exists {
		handlerFn(w, r)
		return
	}
	http.Error(w, "Resource not found", http.StatusNotFound)
}

func NewAppServer() *AppServer {
	return &AppServer{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func main() {
	appServer := NewAppServer()
	go func() {
		for {
			fmt.Scanln()
			fmt.Println("server req count :", appServer.reqCount)
		}
	}()

	appServer.Register("/", indexHandler)
	appServer.Register("/products", productsHandler)
	http.ListenAndServe(":8080", appServer)
}
