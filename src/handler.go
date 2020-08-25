package main

import (
	"html/template"
	"net/http"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("./ui/gohtml/*")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		app.serverError(w, err)
	}
}
