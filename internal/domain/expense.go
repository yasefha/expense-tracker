package domain

import (
	"time"
)

type ExpenseRepository interface {
	Save(expense Expense) (int, error)
	FindAll() ([]Expense, error)
	DeleteByID(id int) error
	UpdateByID(expense Expense) error
}

type Expense struct {
	ID          int
	Date        time.Time
	Description string
	Amount      int
}
