package main

import (
	"awesomeProject/db"
	"awesomeProject/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db.DbConnection()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Home)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
		return
	}

}
