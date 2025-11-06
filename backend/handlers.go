package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"fmt"
)

var (
	tasksMu sync.Mutex
	tasks   = make(map[string]Task)
	dataFile = "tasks.json" // arquivo opcional de persistência
)

func validaStatus(s string) bool {
	return s == StatusTodo || s == StatusInProgress || s == StatusDone
}

func corsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "id obrigatório no path", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		updateTask(w, r, id)
	case http.MethodDelete:
		deleteTask(w, r, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	list := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		list = append(list, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" {
		http.Error(w, "title é obrigatório", http.StatusBadRequest)
		return
	}
	if input.Status == "" {
		input.Status = StatusTodo
	}
	if !validaStatus(input.Status) {
		http.Error(w, "status inválido", http.StatusBadRequest)
		return
	}

	now := time.Now()
	task := Task{
		ID:          generateID(),
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasksMu.Lock()
	tasks[task.ID] = task
	tasksMu.Unlock()

	_ = saveToFile()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request, id string) {
	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	tasksMu.Lock()
	task, ok := tasks[id]
	if !ok {
		tasksMu.Unlock()
		http.Error(w, "task não encontrada", http.StatusNotFound)
		return
	}

	if input.Title != nil {
		val := strings.TrimSpace(*input.Title)
		if val == "" {
			tasksMu.Unlock()
			http.Error(w, "title é obrigatório", http.StatusBadRequest)
			return
		}
		task.Title = val
	}
	if input.Description != nil {
		task.Description = *input.Description
	}
	if input.Status != nil {
		if !validaStatus(*input.Status) {
			tasksMu.Unlock()
			http.Error(w, "status inválido", http.StatusBadRequest)
			return
		}
		task.Status = *input.Status
	}

	task.UpdatedAt = time.Now()
	tasks[id] = task
	tasksMu.Unlock()

	_ = saveToFile()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request, id string) {
	tasksMu.Lock()
	_, ok := tasks[id]
	if !ok {
		tasksMu.Unlock()
		http.Error(w, "task não encontrada", http.StatusNotFound)
		return
	}
	delete(tasks, id)
	tasksMu.Unlock()

	_ = saveToFile()

	w.WriteHeader(http.StatusNoContent)
}

func saveToFile() error {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	list := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		list = append(list, t)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dataFile, data, 0644)
}

func loadFromFile() error {
	if _, err := os.Stat(dataFile); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}
	var list []Task
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	tasksMu.Lock()
	for _, t := range list {
		tasks[t.ID] = t
	}
	tasksMu.Unlock()
	return nil

}

func generateID() string {
return fmt.Sprintf("%d", time.Now().UnixNano())
}