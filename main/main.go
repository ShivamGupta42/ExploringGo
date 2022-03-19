package main

import (
	"ExploringGo/concurrency"
	"fmt"
	"net/http"
)

func main() {

	////Register Routes
	//http.HandleFunc("/Route1", Route1)
	//http.HandleFunc("/Route2", Route2)
	//http.HandleFunc("/Route3", Route3)
	//
	//log.Fatal(http.ListenAndServe(":2000", nil))
	concurrency.ConcurMultiplication()

}

func Route1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Route1")
}

func Route2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Route2")
}
func Route3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Route3")
}
