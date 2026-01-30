package csvrepo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/yasefha/expense-tracker/internal/domain"
)

type CSVExpenseRepository struct {
	filePath string
}

func NewCSVExpenseRepository(filepath string) *CSVExpenseRepository {
	return &CSVExpenseRepository{
		filePath: filepath,
	}
}

func (r *CSVExpenseRepository) Save(expense domain.Expense) (int, error) {
	expenses, err := r.FindAll()
	if err != nil {
		return 0, err
	}

	nextID := 1
	for _, e := range expenses {
		if nextID <= e.ID {
			nextID = e.ID
			nextID++
		}
	}

	file, err := os.OpenFile(r.filePath, os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}

	w := csv.NewWriter(file)
	err = w.Write([]string{
		strconv.Itoa(nextID),
		expense.Date.Format("02/01/2006"),
		expense.Description,
		strconv.Itoa(expense.Amount),
	})
	if err != nil {
		return 0, err
	}

	defer file.Close()
	w.Flush()

	return nextID, w.Error()
}

func (r *CSVExpenseRepository) FindAll() ([]domain.Expense, error) {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	expenses := []domain.Expense{}

	for i, row := range records {
		if len(row) < 4 {
			fmt.Printf("warning: skipping row %d, not enough columns\n", i)
			continue
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Printf("warning: skipping row %d, invalid id '%s'\n", i, row[0])
			continue
		}

		date, err := time.Parse("02/01/2006", row[1])
		if err != nil {
			fmt.Printf("warning: skipping row %d, invalid date '%s'\n", i, row[1])
			continue
		}

		amount, err := strconv.Atoi(row[3])
		if err != nil {
			fmt.Printf("warning: skipping row %d, invalid amount '%s'\n", i, row[3])
			continue
		}

		expense := domain.Expense{
			ID:          id,
			Date:        date,
			Description: row[2],
			Amount:      amount,
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (r *CSVExpenseRepository) DeleteByID(id int) error {
	file, err := os.Open(r.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("expense file not found")
		}
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) <= 1 {
		return errors.New("no expenses to delete")
	}

	newRecords := [][]string{}

	found := false

	for i, row := range records {
		if i == 0 {
			newRecords = append(newRecords, row)
			continue
		}

		rowID, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Printf("warning: skipping row %d invalid id '%s'\n", i, row[0])
			continue
		}

		if rowID == id {
			found = true
			continue
		}

		newRecords = append(newRecords, row)
	}

	if !found {
		return errors.New("expense not found")
	}

	// rewrite file
	file, err = os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	if err := w.WriteAll(newRecords); err != nil {
		return err
	}
	w.Flush()
	return w.Error()
}

func (r *CSVExpenseRepository) UpdateByID(e domain.Expense) error {
	readFile, err := os.OpenFile(r.filePath, os.O_RDONLY, 0444)
	if err != nil {
		return err
	}
	defer readFile.Close()

	reader := csv.NewReader(readFile)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	target := []string{
		strconv.Itoa(e.ID),
		e.Date.Format("02/01/2006"),
		e.Description,
		strconv.Itoa(e.Amount),
	}

	var newRecords [][]string

	for i, row := range records {
		rowID, err := strconv.Atoi(row[0])
		if err != nil {
			return err
		}

		if rowID == e.ID {
			newRecords = append(newRecords, target)
			continue
		}

		newRecords = append(newRecords, records[i])
	}

	writeFile, err := os.OpenFile(r.filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(writeFile)
	if err := writer.WriteAll(newRecords); err != nil {
		return err
	}

	writer.Flush()
	return writer.Error()

}
