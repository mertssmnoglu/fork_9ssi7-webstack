package rest

import (
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
	api := app.Group("/api")
	api.Get("/accounts", s.GetAccounts)
	api.Get("/transactions", s.GetTransactions)
}

func (s *srv) GetAccounts(c *fiber.Ctx) error {
	account := getDummyAccount()
	return c.JSON(account)
}

func (s *srv) GetTransactions(c *fiber.Ctx) error {
	transactions := getDummyTransactions()
	return c.JSON(transactions)
}

func getDummyAccount() models.Account {
	return models.Account{
		ID:                 "1",
		Balance:            1500.50,
		RecentTransactions: getDummyTransactions(),
	}
}

func getDummyTransactions() []models.Transaction {
	return []models.Transaction{
		{ID: "1", Amount: -50.00, Description: "Market Alışverişi"},
		{ID: "2", Amount: 2500.00, Description: "Maaş"},
		{ID: "3", Amount: -150.00, Description: "Elektrik Faturası"},
	}
}
