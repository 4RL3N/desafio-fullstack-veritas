package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := loadFromFile(); err != nil {
		log.Println("aviso: não foi possível carregar tasks do arquivo:", err)
	}

	http.HandleFunc("/tasks", handleTasks)   // GET /tasks, POST /tasks
	http.HandleFunc("/tasks/", handleTask)   // PUT /tasks/{id}, DELETE /tasks/{id}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Mini Kanban API (Go). Endpoints: GET/POST /tasks | PUT/DELETE /tasks/{id}")
	})

	addr := ":8080"
	log.Println("Servidor iniciado em", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
