package main

import (
    "log"
    "net/http"
    "taskmanager/handlers"
)

func main() {
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            handlers.GetTasks(w, r)
        case "POST":
            handlers.CreateTask(w, r)
        case "PATCH":
            handlers.ToggleTask(w, r)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
