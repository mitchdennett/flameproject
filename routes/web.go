package routes

import (
	flameroutes "github.com/mitchdennett/flameframework/routes"
	"github.com/mitchdennett/flameproject/controllers"
)

var Routes = []flameroutes.Route{
	flameroutes.Get().Define("/auth/:id", controllers.Show),
	flameroutes.Get().Define("/", controllers.ShowRoot),
}
