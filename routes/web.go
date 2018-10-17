package routes

import (
	flameroutes "github.com/flame/routes"
	"github.com/app/controllers"
)

var Routes = []flameroutes.Route{
 	flameroutes.Get().Define("/auth/:id", controllers.Show),
	flameroutes.Get().Define("/", controllers.ShowRoot),
}