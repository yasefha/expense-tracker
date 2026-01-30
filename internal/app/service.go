package app

import (
	"errors"
	"time"

	"github.com/yasefha/expense-tracker/internal/domain"
)

type ExpenseService struct {
	repo domain.ExpenseRepository
}

func NewExpenseService(repo domain.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) AddExpense(desc string, amount int) (int, error) {
	if desc == "" {
		return 0, errors.New("description must not be empty")
	}

	if amount <= 0 {
		return 0, errors.New("amount must be greater than 0")
	}

	expense := domain.Expense{
		Date:        time.Now(),
		Description: desc,
		Amount:      amount,
	}

	return s.repo.Save(expense)
}

func (s *ExpenseService) ListExpenses() ([]domain.Expense, error) {
	return s.repo.FindAll()
}

func (s *ExpenseService) GetTotalAmount() (int, error) {
	expenses, err := s.repo.FindAll()
	if err != nil {
		return 0, err
	}

	total := 0

	for _, e := range expenses {
		total += e.Amount
	}

	return total, nil
}

func (s *ExpenseService) GetTotalAmountByMonth(month int) (int, error) {
	if month < 1 || month > 12 {
		return 0, errors.New("Invalid month")
	}

	expenses, err := s.repo.FindAll()
	if err != nil {
		return 0, err
	}

	total := 0
	for _, e := range expenses {
		if int(e.Date.Month()) == month {
			total += e.Amount
		}
	}

	return total, nil
}

func (s *ExpenseService) DeleteExpense(id int) error {
	if id <= 0 {
		return errors.New("invalid expense id")
	}

	return s.repo.DeleteByID(id)
}

func (s *ExpenseService) UpdateExpense(id int, newDate *time.Time, newDesc *string, newAmount *int) (int, error) {
	if id <= 0 {
		return 0, errors.New("invalid expense id")
	}

	if newDate == nil && newDesc == nil && newAmount == nil {
		return 0, errors.New("nothing to update")
	}

	expenses, err := s.repo.FindAll()
	if err != nil {
		return 0, err
	}

	var target *domain.Expense
	for _, e := range expenses {
		if e.ID == id {
			target = &e
			break
		}
	}

	if target == nil {
		return 0, errors.New("expense not found")
	}

	if newDate != nil {
		target.Date = *newDate
	}

	if newDesc != nil {
		target.Description = *newDesc
	}

	if newAmount != nil {
		target.Amount = *newAmount
	}

	if err := s.repo.UpdateByID(*target); err != nil {
		return 0, err
	}

	return target.ID, nil
}
