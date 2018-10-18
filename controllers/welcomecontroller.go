package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchdennett/flameframework/view"
	"github.com/mitchdennett/flameframework/routes"
)

func init(){
	routes.Get("/", "AuthMiddleware", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		view.Render("welcome.html", nil)
	})
}