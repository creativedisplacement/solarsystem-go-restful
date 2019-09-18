package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/solarsystem-go-restful/data"
	"github.com/solarsystem-go-restful/planet"
	"github.com/solarsystem-go-restful/planets"
)

var DB *sqlx.DB

func main() {
	var err error
	DB, err = sqlx.Connect("sqlite3", "../solarsystem.db")
	if err != nil {
		log.Println("Problem creating the database")
	}
	defer DB.Close()
	planet.DB = DB
	planets.DB = DB
	data.DB = DB
	data.Initialize()

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	p := planet.Planet{}
	p.Register(wsContainer)
	ps := planets.Planets{}
	ps.Register(wsContainer)

	log.Printf("start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
