package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/necrophonic/pony/internal/db"
	"github.com/necrophonic/pony/internal/pony"
)

func main() {

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Fatal("Didn't get a PORT env var!")
	}

	http.HandleFunc("/pony", ponyHandler)
	http.HandleFunc("/healthcheck", healthcheckHandler)

	http.ListenAndServe(":"+port, nil)
}

func ponyHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		name := r.URL.Query()["name"]
		if len(name) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "No 'name' supplied!")
			return
		}

		// Attempt to fetch the pony from the DB. Return a "not found"
		// to the client if it wasn't found in the database
		pony, err := db.GetPony(name[0])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Pony not found")
			return
		}

		// Now we have a pony, let's marshal that information into json
		// and return it to our client
		json, err := json.Marshal(pony)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Failed to marshal pony data")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	if r.Method == http.MethodPost {

		// Grab the body from the request - for this example we're going to
		// assume the data is all going to be small enough to safely read
		// without blowing any memory limits in one go.
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "No pony data supplied!")
			return
		}

		log.Println(body)

		// Unmarshal the request into a pony
		var pony pony.Pony
		err = json.Unmarshal(body, &pony)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Couldn't unmarshal pony data: %v", err)
			return
		}

		if err = db.SetPony(pony.Name, pony.Element); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to create pony: %v", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return

	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"friendship":"magic!"}`)
}
