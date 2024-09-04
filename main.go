package main

import (
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/di"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/router"
)

func main() {

	di.Init()

	server := router.Initialize()

	server.Run(":8080")
}
