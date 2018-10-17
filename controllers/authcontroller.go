package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/flame/view"
)

type AuthController struct {

}

func (AuthController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view.Render("auth.html", nil)
}