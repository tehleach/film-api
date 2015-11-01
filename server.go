package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tehleach/film-api/db"
)

func main() {
	insertTestDbData()
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index is the main server index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//Film represents a film
type Film struct {
	Name     string
	Director string
}

func insertTestDbData() {
	dbConn, err := db.NewConn("localhost")
	if err != nil {
		panic(err)
	}
	c := dbConn.Session.DB("film-api").C("films")
	err = c.Insert(&Film{"Django Unchained", "Quentin Tarantino"})
	if err != nil {
		panic(err)
	}
}
