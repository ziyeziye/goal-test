package routes

import (
	"github.com/goal-web/contracts"
	"goal-test/app/api"
)

func ApiRoutes(router contracts.Router) {
	router.Get("/users", api.GetUsers)
}
