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

	NewServer().ServeHTTP(rec, req)

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

	NewServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if body := rec.Body.String(); body != "ArcNote API is running\n" {
		t.Fatalf("body = %q, want %q", body, "ArcNote API is running\n")
	}
}

func TestNewServerCreatesStudyRecord(t *testing.T) {
	body := []byte(`{"durationMinutes":60,"content":"Go HTTP basics","studiedOn":"2026-03-16"}`)
	req := httptest.NewRequest(http.MethodPost, "/study-records", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	NewServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}

	if got := rec.Header().Get("Location"); got != "/study-records/temporary-id" {
		t.Fatalf("location = %q, want %q", got, "/study-records/temporary-id")
	}

	if got := rec.Body.String(); got != "study record created\n" {
		t.Fatalf("body = %q, want %q", got, "study record created\n")
	}
}
