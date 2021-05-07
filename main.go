package main

import "github.com/Damillora/Yuriko/pkg/app"

func main() {
	app := app.App{}
	app.Initialize()
	app.StartApp()
}
