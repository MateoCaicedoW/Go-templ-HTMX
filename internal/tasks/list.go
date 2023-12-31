package tasks

import (
	"database/sql"
	"templ/models"
)

func List(db *sql.DB) (models.Tasks, error) {
	rows, err := db.Query("SELECT * FROM tasks ORDER BY active DESC, created_at DESC")
	if err != nil {
		return nil, err
	}

	var tasks models.Tasks
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Active, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}
