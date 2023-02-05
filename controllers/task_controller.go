package controllers

import (
	"net/http"

	"github.com/regmarmcem/apprunner-example/services"
)

type TaskController struct {
	service services.TaskServicer
}

func NewTaskController(s services.TaskServicer) *TaskController {
	return &TaskController{service: s}
}

func (c *TaskController) GetTaskHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome"))
}
