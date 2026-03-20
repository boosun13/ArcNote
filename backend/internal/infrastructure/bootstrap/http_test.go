package bootstrap

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHTTPServerCreatesStudyRecord(t *testing.T) {
	body := []byte(`{"durationMinutes":60,"content":"Go HTTP basics","studiedOn":"2026-03-16"}`)
	req := httptest.NewRequest(http.MethodPost, "/study-records", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	NewHTTPServer().ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}

	if got := rec.Header().Get("Location"); got != "/study-records/temporary-id" {
		t.Fatalf("location = %q, want %q", got, "/study-records/temporary-id")
	}
}
