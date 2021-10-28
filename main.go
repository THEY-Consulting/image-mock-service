package main

import (
	"github.com/they-consulting/image-mock-service/api"
	"github.com/they-consulting/image-mock-service/application"
	"github.com/they-consulting/image-mock-service/config"
)

func main() {
	config.New()
	api.InitRouter(application.Get().Config.Port)
}
