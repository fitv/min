package main

import (
	"embed"

	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/service"
	"github.com/fitv/min/global"
)

//go:embed static templates
var fs embed.FS

func main() {
	global.App = app.NewApplication(fs)
	defer global.App.Close()

	global.App.AddService(
		&service.Logger{},
		&service.Cache{},
		&service.Redis{},
		&service.Database{},
		&service.Validator{},
		&service.Translator{},
		&service.Gin{},
		&service.Router{},
	)

	global.App.Run()
}
