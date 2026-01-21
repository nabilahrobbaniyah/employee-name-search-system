package main

import (
	"database/sql"
	"fmt"
)

func getEmployees(db *sql.DB, n int) ([]Employee, error) {
	if n <= 0 {
		return []Employee{}, nil
	}

	query := fmt.Sprintf(
		"SELECT id, name, department FROM employees LIMIT %d",
		n,
	)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee

	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Department); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}
