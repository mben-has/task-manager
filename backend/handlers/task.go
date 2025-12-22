package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "taskmanager/models"
)

var tasks []models.Task
var nextID = 1

func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)
    task.ID = nextID
    nextID++
    tasks = append(tasks, task)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func ToggleTask(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)
    for i, t := range tasks {
        if t.ID == id {
            tasks[i].Done = !tasks[i].Done
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(tasks[i])
            return
        }
    }
    http.NotFound(w, r)
}
