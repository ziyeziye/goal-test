package api

import (
	"github.com/goal-web/contracts"
	"goal-test/app/models"
)

func GetUsers(request contracts.HttpRequest) interface{} {
	return contracts.Fields{
		"name":   request.StringOption("name", "user hello"),
		"models": models.UserQuery().Select("name", "address", "age").Where("age", ">", 0).Get().ToArrayFields(),
	}
}
