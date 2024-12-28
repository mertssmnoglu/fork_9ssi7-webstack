package models

type Transaction struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type Account struct {
	ID                 string        `json:"id"`
	Balance            float64       `json:"balance"`
	RecentTransactions []Transaction `json:"recentTransactions"`
}
