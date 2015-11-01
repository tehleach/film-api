package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tehleach/film-api/db"
)

var conn *db.Conn

func main() {
	var err error
	conn, err = db.NewConn("localhost")
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/films", Films)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index is the main server index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//Films is the film list
func Films(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var results []Film
	err := conn.Session.DB("film-api").C("films").Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
