package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "taskmanager/models"
)

var Tasks []models.Task
var NextID = 1

func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)
    task.ID = NextID
    NextID++
    Tasks = append(Tasks, task)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func ToggleTask(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)
    for i, t := range Tasks {
        if t.ID == id {
            Tasks[i].Done = !Tasks[i].Done
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(Tasks[i])
            return
        }
    }
    http.NotFound(w, r)
}
