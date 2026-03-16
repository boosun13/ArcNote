package httpadapter

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServerReturnsHealthz(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	NewServer(Routes{}).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if body := rec.Body.String(); body != "ok\n" {
		t.Fatalf("body = %q, want %q", body, "ok\n")
	}
}

func TestNewServerReturnsRootMessage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	NewServer(Routes{}).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if body := rec.Body.String(); body != "ArcNote API is running\n" {
		t.Fatalf("body = %q, want %q", body, "ArcNote API is running\n")
	}
}

func TestNewServerDelegatesStudyRecordRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/study-records", bytes.NewReader([]byte(`{}`)))
	rec := httptest.NewRecorder()

	handlerCalled := false
	studyRecordHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusCreated)
	})

	NewServer(Routes{StudyRecord: studyRecordHandler}).ServeHTTP(rec, req)

	if !handlerCalled {
		t.Fatal("study record handler was not called")
	}

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}
}
