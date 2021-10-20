package application

import (
	"github.com/they-consulting/gogogo/config"
)

type Application struct {
	Config *config.Config
}

var app Application

func init() {
	app = Application{}
	app.Config = config.Get()
}

func Get() Application {
	return app
}
