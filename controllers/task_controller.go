package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/regmarmcem/apprunner-example/models"
	"github.com/regmarmcem/apprunner-example/services"
)

type TaskController struct {
	service services.TaskServicer
}

func NewTaskController(s services.TaskServicer) *TaskController {
	return &TaskController{service: s}
}

func (c *TaskController) GetTaskHandler(w http.ResponseWriter, req *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Path parameter must be a number", http.StatusBadRequest)
		return
	}
	task, err := c.service.GetTaskService(taskID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (c *TaskController) PostTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task models.Task
	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		http.Error(w, "Cannot decode request body", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(task)
}
