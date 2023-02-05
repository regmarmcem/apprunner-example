package repositories

import (
	"database/sql"
	"log"

	"github.com/regmarmcem/apprunner-example/models"
)

func SelectTaskDetail(db *sql.DB, taskID int) (models.Task, error) {
	const sqlStr = `
		select task_id, title, detail, created_at
		from tasks
		where task_id = ?;
	`
	row := db.QueryRow(sqlStr, taskID)
	if err := row.Err(); err != nil {
		log.Println("SelectTaskDetail failed")
		return models.Task{}, nil
	}

	var task models.Task
	var createdTime sql.NullTime

	err := row.Scan(&task.ID, &task.Title, &task.Detail, &createdTime)
	if err != nil {
		return models.Task{}, err
	}

	if createdTime.Valid {
		task.CreatedAt = createdTime.Time
	}

	return task, nil
}
