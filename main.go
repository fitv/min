package main

import (
	"embed"

	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/service"
	_ "github.com/fitv/min/ent/runtime"
	"github.com/fitv/min/global"
)

//go:embed static templates locales
var fs embed.FS

func main() {
	global.App = app.NewApplication(fs)
	defer global.App.Shutdown()

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
