package main

import(
	"flag"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() *mux.Router {
	var dir string

	flag.StringVar(&dir, "dir", "./ui/static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	mux := mux.NewRouter()
	mux.HandleFunc("/", app.index).Methods("GET")

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))
	return mux
}

