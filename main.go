package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/config"
	"github.com/goal-web/console"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/events"
	"github.com/goal-web/http"
	"github.com/goal-web/redis"
	console2 "goal-test/app/console"
	"goal-test/app/exceptions"
	config2 "goal-test/config"
	"goal-test/database/migrations"
	"goal-test/routes"
	"os"
)

func main() {
	app := application.Singleton()
	path, _ := os.Getwd()
	app.Instance("path", path)

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})
	app.RegisterServices(
		config.Service(os.Getenv("env"), path, config2.Configs()),
		events.ServiceProvider{},
		redis.ServiceProvider{},
		&console.ServiceProvider{
			ConsoleProvider: console2.NewKernel,
		},
		database.Service(migrations.Migrations),
		&http.ServiceProvider{RouteCollectors: []interface{}{
			// 路由收集器
			routes.ApiRoutes,
		}},
	)

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(input)
	})
}
