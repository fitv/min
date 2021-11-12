package service

import (
	"fmt"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/db"
)

type Database struct {
	Service
}

func (Database) Register(app *app.Application) {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&timeout=%s&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
		config.Database.Charset,
		config.Database.Collation,
		config.Database.Timeout,
	)

	switch config.Database.Driver {
	case "mysql":
		db, err := db.New(&db.Option{
			Dns:    dns,
			Driver: config.Database.Driver,
			Debug:  config.Database.Debug,
		})
		if err != nil {
			panic(fmt.Errorf("database init error: %w", err))
		}
		app.DB = db

		app.AddClose(func() {
			app.DB.Close()
		})
	default:
		panic(fmt.Errorf("unsupported database driver: %s", config.Database.Driver))
	}
}
