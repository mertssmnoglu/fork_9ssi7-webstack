package web

import (
	"github.com/9ssi7/webstack/api/web/templates"
	"github.com/9ssi7/webstack/internal/models"
	"github.com/9ssi7/webstack/pkg/server"
	"github.com/gofiber/fiber/v2"
)

type srv struct {
}

func New() server.Rest {
	return &srv{}
}

func (s *srv) Register(app *fiber.App) {
	app.Static("/static", "./api/web/static/dist") // Changed to use working directory path

	// Register Web Routes
	app.Get("/", s.HandleHome)
	app.Get("/dashboard", s.HandleDashboard)
}

func (s *srv) HandleHome(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	component := templates.Home()
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func (s *srv) HandleDashboard(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	accountData := getDummyAccount()
	component := templates.Dashboard(accountData)
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func getDummyAccount() models.Account {
	// ...existing code from handlers.go...
	return models.Account{
		ID:                 "1",
		Balance:            1500.50,
		RecentTransactions: getDummyTransactions(),
	}
}

func getDummyTransactions() []models.Transaction {
	// ...existing code from handlers.go...
	return []models.Transaction{
		{ID: "1", Amount: -50.00, Description: "Market Alışverişi"},
		{ID: "2", Amount: 2500.00, Description: "Maaş"},
		{ID: "3", Amount: -150.00, Description: "Elektrik Faturası"},
	}
}
