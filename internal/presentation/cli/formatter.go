package cli

import (
	"fmt"

	"github.com/yasefha/expense-tracker/internal/domain"
)

type CLIFormatter struct{}

func NewCLIFormatter() *CLIFormatter {
	return &CLIFormatter{}
}

func (f *CLIFormatter) PrintExpenseTable(expenses []domain.Expense) {
	fmt.Printf("%-3s %-10s %-15s %s \n", "ID", "Date", "Description", "Amount")
	for _, e := range expenses {
		fmt.Printf("%-3d %-10s %-15s $%d \n", e.ID, e.Date.Format("02/01/2006"), e.Description, e.Amount)
	}
}
