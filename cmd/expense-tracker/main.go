package main

import (
	"fmt"
	"os"

	"github.com/yasefha/expense-tracker/internal/app"
	csvrepo "github.com/yasefha/expense-tracker/internal/infra/storage/csv-repo"
)

func main() {
	repo := csvrepo.NewCSVExpenseRepository("data/expenses.csv")
	service := app.NewExpenseService(repo)

	if len(os.Args) == 1 {
		fmt.Println(`Expense Tracker Program.
Run "expense-tracker add --description <string> --amount <integer> to add new expense.`)
		return
	}

	switch os.Args[1] {
	case "add":
		RunAddCommand(os.Args[2:], service)
	case "list":
		RunListCommand(os.Args[2:], service)
	case "summary":
		RunSummaryComand(os.Args[2:], service)
	case "update":
		RunUpdateCommand(os.Args[2:], service)
	case "delete":
		RunDeleteCommand(os.Args[2:], service)
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
}
