package application

import (
	"github.com/they-consulting/image-mock-service/config"
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
