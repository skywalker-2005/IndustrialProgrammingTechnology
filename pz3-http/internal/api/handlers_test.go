package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/skywalker-2005/IndustrialProgrammingTechnology/pz3-http/internal/storage"
)

func newTestHandlers() *Handlers {
	return NewHandlers(storage.NewMemoryStore())
}

func TestCreateTask_Success(t *testing.T) {
	h := newTestHandlers()

	body := bytes.NewBufferString(`{"title":"Купить молоко"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}

	var task storage.Task
	if err := json.NewDecoder(rec.Body).Decode(&task); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if task.Title != "Купить молоко" {
		t.Errorf("unexpected title: %q", task.Title)
	}
	if task.ID == 0 {
		t.Error("expected non-zero id")
	}
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	h := newTestHandlers()

	body := bytes.NewBufferString(`{"title":"   "}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestCreateTask_InvalidJSON(t *testing.T) {
	h := newTestHandlers()

	body := bytes.NewBufferString(`{title:}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestGetTask_NotFound(t *testing.T) {
	h := newTestHandlers()

	req := httptest.NewRequest(http.MethodGet, "/tasks/999", nil)
	rec := httptest.NewRecorder()

	h.GetTask(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}

func TestListTasks_Empty(t *testing.T) {
	h := newTestHandlers()

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()

	h.ListTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var tasks []storage.Task
	if err := json.NewDecoder(rec.Body).Decode(&tasks); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected empty list, got %d", len(tasks))
	}
}

func TestListTasks_FilterByQuery(t *testing.T) {
	h := newTestHandlers()
	h.Store.Create("Купить молоко")
	h.Store.Create("Помыть посуду")

	req := httptest.NewRequest(http.MethodGet, "/tasks?q=молоко", nil)
	rec := httptest.NewRecorder()

	h.ListTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var tasks []storage.Task
	if err := json.NewDecoder(rec.Body).Decode(&tasks); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Title != "Купить молоко" {
		t.Errorf("unexpected title: %q", tasks[0].Title)
	}
}

func TestGetTask_InvalidID(t *testing.T) {
	h := newTestHandlers()

	req := httptest.NewRequest(http.MethodGet, "/tasks/abc", nil)
	rec := httptest.NewRecorder()

	h.GetTask(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

