package main

import (
	"github.com/they-consulting/gogogo/api"
	"github.com/they-consulting/gogogo/application"
	"github.com/they-consulting/gogogo/config"
)

func main() {
	config.New()
	api.InitRouter(application.Get().Config.Port)
}
