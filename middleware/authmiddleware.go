package middleware

import (
	"reflect"
	"net/http"
)

type AuthMiddlewareStruct struct {}

func AuthMiddleware() (reflect.Type) {
	return reflect.TypeOf(AuthMiddlewareStruct{})
}

func (AuthMiddlewareStruct) Before(w http.ResponseWriter, r *http.Request) bool {
	if false {
		http.Redirect(w, r, "/", 307)
		return true
	}

	return false
}

func (AuthMiddlewareStruct) After(w http.ResponseWriter, r *http.Request) {
	
}
