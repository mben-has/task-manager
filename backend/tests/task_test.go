package tests

import (
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "taskmanager/handlers"
    "taskmanager/models"
    "testing"
)

func TestCreateAndGetTasks(t *testing.T) {
    task := models.Task{Title: "Test task"}
    body, _ := json.Marshal(task)

    req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    handlers.CreateTask(w, req)

    if w.Code != 200 {
        t.Errorf("Expected status 200, got %d", w.Code)
    }

    reqGet := httptest.NewRequest("GET", "/tasks", nil)
    wGet := httptest.NewRecorder()
    handlers.GetTasks(wGet, reqGet)

    if wGet.Code != 200 {
        t.Errorf("Expected status 200, got %d", wGet.Code)
    }
}
