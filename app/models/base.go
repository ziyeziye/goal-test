package models

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
)

func Query(name string) contracts.QueryBuilder {
	app := application.Singleton()
	return table.WithConnection(name, app.Get("db").(contracts.DBConnection))
}
