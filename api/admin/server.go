package admin

import (
	"embed"

	"github.com/9ssi7/webstack/pkg/server"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

//go:embed dist/*
var distFS embed.FS

type srv struct{}

func New() server.Rest {
	return &srv{}
}

func (s *srv) Register(app *fiber.App) {
	app.Get("/admin*", static.New("", static.Config{
		FS:            distFS,
		Browse:        true,
		IndexNames:    []string{"dist/index.html"},
		CacheDuration: 0,
		MaxAge:        0,
	}))
}
