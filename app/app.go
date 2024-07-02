package app

import (
	"client-ccs/app/client"
	"client-ccs/app/config"
	"client-ccs/app/service"
)

type Application struct {
	conf *config.Config
}

func NewApplication(conf *config.Config) *Application {
	return &Application{conf}
}

func (app *Application) Start() {
	loadService := service.NewLoadService(
		app.conf.Rps,
		client.NewHttpClient(),
		app.conf.CcsUrl,
		app.conf.CcsClientToken,
	)
	loadService.StartLoadTest()
}
