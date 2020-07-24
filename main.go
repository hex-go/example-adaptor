package main

import (
	_ "example/adaptor/standard"
	"example/api"
)

func main() {
	engine := api.Routers()
	_ = engine.Run(":8887")
}
