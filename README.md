# Expense Tracker
Expense Tracker is a simple command-line application written in Go for managing personal expenses.
This project is built as part of the roadmap.sh backend learning path to practice building a real-world CLI application with clean structure, clear separation of concerns, and file-based persistence using CSV.

## Roadmap.sh Project Reference
This project is based on the following roadmap.sh project:
[roadmap.sh/project/expense-tracker](https://roadmap.sh/projects/expense-tracker)
 
## Features
- Add an expense with a description and amount.
- Update an expense.
- Delete an expense.
- View all expense.
- View a summary of all expenses.
- View a summary of expenses for a spesific month (of current year).

## Commands and output
```bash
$ expense-tracker add --description "Lunch" --amount 20
Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
Expense added successfully (ID: 2)

$ expense-tracker list
ID  Date        Description Amount
1   23/01/2026  Lunch       $20
2   23/01/2026  Dinner      $10

$ expense-tracker summary
Total expenses: $30

$ expense-tracker summary --month 1
Total expenses for January: $10

$ expense-tracker update --id 1 --description "Breakfast"
Expense updated successfully (ID: 1)

$ expense-tracker update --id 1 --amount 10
Expense updated successfully (ID: 1)

$ expense-tracker update --id 1 --date 22/01/2026
Expense updated successfully (ID: 1)

$ expense-tracker delete --id 2
Expense deleted successfully (ID: 2)
```

## Notes & Constraints
- Dates use format DD/MM/YYYY
- All amounts are positive integers.
- Currency is fixed and not configurable.
- When updating an expense:
    - Only provided fields will be updated.
    - Fields with empty or zero values are ignored.
    - At least one field must be provided.
- Monthly summary calculates expenses for the given month in the current year (based on system time).
