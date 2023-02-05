package services

import (
	"github.com/regmarmcem/apprunner-example/models"
)

func (s TaskAppService) GetTaskService(taskID int) (models.Task, error) {
	return models.Task{}, nil
}
