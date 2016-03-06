package main

import (
	"log"
	"memorymaps-backend/db/postgres"
	"memorymaps-backend/routers"
	"runtime"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
}

func main() {

	log.Println("Create missing tables")
	db.CreateTablesIfNotExists()

	// Start the server
	router := mux.NewRouter().StrictSlash(true)
	router = routers.AddTextMemoryRoutes(router)
	n := negroni.Classic()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(router)
	n.Run(":80")

}
