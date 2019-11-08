package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/necrophonic/pony/internal/db"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("No PORT supplied")
	}

	r := httprouter.New()

	r.GET("/pony", ponyListHandler)
	r.GET("/heathcheck", healthcheckHandler)

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}

func ponyListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func healthcheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"friendship":"magic"}`)
}
