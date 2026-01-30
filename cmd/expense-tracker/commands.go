package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/yasefha/expense-tracker/internal/app"
	"github.com/yasefha/expense-tracker/internal/presentation/cli"
)

func RunAddCommand(args []string, s *app.ExpenseService) {
	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)

	desc := addCmd.String("description", "", "expense description")
	amount := addCmd.Int("amount", 0, "expense amount")

	if err := addCmd.Parse(args); err != nil {
		fmt.Println("failed to parse add command:", err)
		return
	}

	if *desc == "" {
		fmt.Println("description is required")
		return
	}

	if *amount <= 0 {
		fmt.Println("amount must be greater than 0")
		return
	}

	id, err := s.AddExpense(*desc, *amount)
	if err != nil {
		fmt.Println("failed to add expense:", err)
		return
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", id)
}

func RunListCommand(args []string, s *app.ExpenseService) {
	if len(args) > 0 {
		fmt.Println("list command does not accept any arguments")
		return
	}

	expenses, err := s.ListExpenses()
	if err != nil {
		fmt.Println("failed to list expenses:", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	formatter := cli.NewCLIFormatter()
	formatter.PrintExpenseTable(expenses)
}

func RunSummaryComand(args []string, s *app.ExpenseService) {
	summaryCmd := flag.NewFlagSet("summary", flag.ContinueOnError)
	month := summaryCmd.Int("month", 0, "chosen month")

	if err := summaryCmd.Parse(args); err != nil {
		fmt.Println("failed to parse summary command:", err)
		return
	}

	if *month == 0 {
		total, err := s.GetTotalAmount()
		if err != nil {
			fmt.Println("failed to calculate summary:", err)
			return
		}

		fmt.Printf("Total expenses: $%d\n", total)
		return
	}

	total, err := s.GetTotalAmountByMonth(*month)
	if err != nil {
		fmt.Println("failed to calculate summary:", err)
		return
	}

	monthName := time.Month(*month).String()
	fmt.Printf("Total expenses for %s: $%d\n", monthName, total)
}

func RunDeleteCommand(args []string, s *app.ExpenseService) {
	deleteCmd := flag.NewFlagSet("delete", flag.ContinueOnError)
	id := deleteCmd.Int("id", 0, "expense id")

	if err := deleteCmd.Parse(args); err != nil {
		fmt.Println("failed to parse delete command:", err)
		return
	}

	if *id == 0 {
		fmt.Println("please provide --id")
		return
	}

	if err := s.DeleteExpense(*id); err != nil {
		fmt.Println("failed to delete expense:", err)
		return
	}

	fmt.Printf("Expense deleted successfully (ID: %d)\n", *id)
}

func RunUpdateCommand(args []string, s *app.ExpenseService) {
	updateCmd := flag.NewFlagSet("update", flag.ContinueOnError)

	id := updateCmd.Int("id", 0, "expense ID")
	dateStr := updateCmd.String("date", "", "date (DD/MM/YYYY)")
	desc := updateCmd.String("description", "", "new description")
	amount := updateCmd.Int("amount", 0, "new amount")

	if err := updateCmd.Parse(args); err != nil {
		fmt.Println("failed to parse update command:", err)
		return
	}

	if *id == 0 {
		fmt.Println("please provide --id")
		return
	}

	var (
		datePtr   *time.Time
		descPtr   *string
		amountPtr *int
	)

	if *dateStr != "" {
		parsedDate, err := time.Parse("02/01/2006", *dateStr)
		if err != nil {
			fmt.Println("invalid date format, use DD/MM/YYYY")
			return
		}
		datePtr = &parsedDate
	}

	if *desc != "" {
		descPtr = desc
	}

	if *amount != 0 {
		if *amount <= 0 {
			fmt.Println("amount must be greater than 0")
			return
		}
		amountPtr = amount
	}

	if datePtr == nil && descPtr == nil && amountPtr == nil {
		fmt.Println("nothing to update")
		return
	}

	updatedID, err := s.UpdateExpense(*id, datePtr, descPtr, amountPtr)
	if err != nil {
		fmt.Println("failed to update expense:", err)
		return
	}

	fmt.Printf("Expense updated successfully (ID: %d)\n", updatedID)
}
