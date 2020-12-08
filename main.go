
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO API! test")
}


func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)


	log.Fatal(http.ListenAndServe(":3000", router))
}