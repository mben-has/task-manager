package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"taskmanager/handlers"
	"taskmanager/models"
)

func resetTasks() {
	handlers.Tasks = []models.Task{}
	handlers.NextID = 1
}

func TestCreateTask(t *testing.T) {
	resetTasks()

	task := models.Task{Title: "Test Create"}
	body, _ := json.Marshal(task)

	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateTask(w, req)

	if w.Code != 200 {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var created models.Task
	json.NewDecoder(w.Body).Decode(&created)

	if created.ID != 1 || created.Title != "Test Create" || created.Done != false {
		t.Errorf("Created task does not match: %+v", created)
	}
}

func TestGetTasks(t *testing.T) {
	resetTasks()

	handlers.Tasks = append(handlers.Tasks, models.Task{ID: 1, Title: "Task1", Done: false})
	handlers.Tasks = append(handlers.Tasks, models.Task{ID: 2, Title: "Task2", Done: true})

	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	handlers.GetTasks(w, req)

	if w.Code != 200 {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var got []models.Task
	json.NewDecoder(w.Body).Decode(&got)

	if len(got) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(got))
	}
}

func TestToggleTask(t *testing.T) {
	resetTasks()

	handlers.Tasks = append(handlers.Tasks, models.Task{ID: 1, Title: "Test Task to toggle", Done: false})

	req := httptest.NewRequest("PATCH", "/tasks?id=1", nil)
	w := httptest.NewRecorder()
	handlers.ToggleTask(w, req)

	if w.Code != 200 {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var toggled models.Task
	json.NewDecoder(w.Body).Decode(&toggled)

	if toggled.Done != true {
		t.Errorf("Expected Done=true, got Done=%v", toggled.Done)
	}
}

func TestToggleTaskNotFound(t *testing.T) {
	resetTasks()

	req := httptest.NewRequest("PATCH", "/tasks?id=-1", nil)
	w := httptest.NewRecorder()
	handlers.ToggleTask(w, req)

	if w.Code != 404 {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestCreateAndToggleTasks(t *testing.T) {
	resetTasks()

	task1 := models.Task{Title: "Task 1"}
	body1, _ := json.Marshal(task1)
	req1 := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body1))
	w1 := httptest.NewRecorder()
	handlers.CreateTask(w1, req1)

	task2 := models.Task{Title: "Task 2"}
	body2, _ := json.Marshal(task2)
	req2 := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body2))
	w2 := httptest.NewRecorder()
	handlers.CreateTask(w2, req2)

	reqGet := httptest.NewRequest("GET", "/tasks", nil)
	wGet := httptest.NewRecorder()
	handlers.GetTasks(wGet, reqGet)

	var tasks []models.Task
	json.NewDecoder(wGet.Body).Decode(&tasks)

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	reqToggle := httptest.NewRequest("PATCH", "/tasks?id=1", nil)
	wToggle := httptest.NewRecorder()
	handlers.ToggleTask(wToggle, reqToggle)

	var toggled models.Task
	json.NewDecoder(wToggle.Body).Decode(&toggled)

	if toggled.Done != true {
		t.Errorf("Expected Done=true, got %v", toggled.Done)
	}

	reqToggle2 := httptest.NewRequest("PATCH", "/tasks?id=2", nil)
	wToggle2 := httptest.NewRecorder()
	handlers.ToggleTask(wToggle2, reqToggle2)

	var toggled2 models.Task
	json.NewDecoder(wToggle2.Body).Decode(&toggled2)

	if toggled2.Done != true {
		t.Errorf("Expected Done=true, got %v", toggled2.Done)
	}

	reqGetFinal := httptest.NewRequest("GET", "/tasks", nil)
	wGetFinal := httptest.NewRecorder()
	handlers.GetTasks(wGetFinal, reqGetFinal)

	var finalTasks []models.Task
	json.NewDecoder(wGetFinal.Body).Decode(&finalTasks)

	if finalTasks[0].Done != true || finalTasks[1].Done != true {
		t.Errorf("Expected both tasks Done=true, got %+v", finalTasks)
	}
}