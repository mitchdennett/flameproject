package routes

import (
	. "github.com/flame/routes" 
	. "github.com/app/controllers"
	. "github.com/app/middleware" 
)

var Routes = []Route{
	Get().Define("/", WelcomeController{}),
	Get().Define("/auth/mitch", AuthController{}).Middleware(AuthMiddleware),
}