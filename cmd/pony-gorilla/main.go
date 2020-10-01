package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/necrophonic/pony/internal/db"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("No PORT supplied")
	}

	r := mux.NewRouter()

	r.HandleFunc("/pony", ponyListHandler).Methods("GET")
	r.HandleFunc("/pony/{name}", ponyHandler).Methods("GET")
	r.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}

func ponyListHandler(w http.ResponseWriter, r *http.Request) {
	ponies := db.AllPonies()

	json, err := json.Marshal(&ponies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func ponyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name, ok := vars["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pony, err := db.GetPony(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json, err := json.Marshal(&pony)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"friendship":"magic"}`)
}
