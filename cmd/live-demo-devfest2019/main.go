package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/necrophonic/pony/internal/db"
	"github.com/necrophonic/pony/internal/pony"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("no PORT env")
	}

	r := httprouter.New()
	setupRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}

func setupRoutes(r *httprouter.Router) {
	r.GET("/pony/:name", ponyHandler)
	r.POST("/pony", createPonyHandler)

	r.HandlerFunc("GET", "/healthcheck", hcHandler)
	r.HandlerFunc("GET", "/pony", listHandler)
}

func createPonyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var p pony.Pony
	err = json.Unmarshal(body, &p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.SetPony(p.Name, p.Element)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func ponyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	name := ps.ByName("name")

	pony, err := db.GetPony(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if pony == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json, err := json.Marshal(&pony)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-TYpe", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func listHandler(w http.ResponseWriter, r *http.Request) {

	list := db.AllPonies()

	json, err := json.Marshal(&list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-TYpe", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func hcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-TYpe", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"alive":"yes"}`)
}
