package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/flame/view"
	"github.com/flame/managers/mail"
)

type WelcomeController struct {

}

func (WelcomeController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	go mail.Compose().To("mitchelldennett2@gmail.com").Subject("Test Email").Send("This message")
	view.Render("welcome.html", nil)
}