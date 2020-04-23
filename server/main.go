package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func welcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to FuzzMe\n")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomeMessage)
	log.Fatal(http.ListenAndServe(":8080", router))
}
