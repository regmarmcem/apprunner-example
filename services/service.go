package services

import (
	"database/sql"

	"github.com/regmarmcem/apprunner-example/models"
)

type TaskServicer interface {
	GetTaskService(taskID int) (models.Task, error)
}

type TaskAppService struct {
	db *sql.DB
}

func NewTaskAppService(db *sql.DB) *TaskAppService {
	return &TaskAppService{db: db}
}
