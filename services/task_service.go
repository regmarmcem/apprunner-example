package services

import (
	"database/sql"
	"errors"
	"log"

	"github.com/regmarmcem/apprunner-example/models"
	"github.com/regmarmcem/apprunner-example/repositories"
)

func (s *TaskAppService) GetTaskService(taskID int) (models.Task, error) {
	var task models.Task
	task, err := repositories.SelectTaskDetail(s.db, taskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Record not found")
			return models.Task{}, err
		}
		return models.Task{}, err
	}
	return task, nil
}

func (s *TaskAppService) PostTaskService(task models.Task) (models.Task, error) {
	task, err := repositories.InsertTask(s.db, task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
