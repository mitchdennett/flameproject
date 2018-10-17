package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/flame/view"
)

//routemeta: {"method":"get", "route":"/auth/:id", "middleware":"auth"}
func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view.Render("auth.html", nil)
}

//routemeta: {"method":"get", "route":"/"}
func ShowRoot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view.Render("welcome.html", nil)
}