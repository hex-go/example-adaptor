package main

import (
	_ "example/adaptor/standard"
	"example/api"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	engine := api.Routers()
	_ = engine.Run(":8887")

}
