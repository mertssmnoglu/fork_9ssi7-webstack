package server

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Rest interface {
	Register(app *fiber.App)
}

type xhttp struct {
	cnf  XHTTPConfig
	app  *fiber.App
	apps []Rest
}

type XHTTPConfig struct {
	Host string
	Port uint16
}

func NewXHTTP(cnf XHTTPConfig, apps ...Rest) Listener {
	return &xhttp{
		app:  fiber.New(),
		apps: apps,
		cnf:  cnf,
	}
}

func (x *xhttp) Listen() error {
	for _, app := range x.apps {
		app.Register(x.app)
	}
	return x.app.Listen(x.cnf.Host + ":" + strconv.Itoa(int(x.cnf.Port)))
}

func (x *xhttp) Shutdown(ctx context.Context) error {
	return x.app.ShutdownWithContext(ctx)
}
