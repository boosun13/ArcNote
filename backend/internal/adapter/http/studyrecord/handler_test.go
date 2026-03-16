package studyrecord

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerCreatesStudyRecord(t *testing.T) {
	body := []byte(`{"durationMinutes":60,"content":"Go HTTP basics","studiedOn":"2026-03-16"}`)
	req := httptest.NewRequest(http.MethodPost, "/study-records", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	NewHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}

	if got := rec.Header().Get("Location"); got != "/study-records/temporary-id" {
		t.Fatalf("location = %q, want %q", got, "/study-records/temporary-id")
	}

	if got := rec.Body.String(); got != "study record temporary-id created\n" {
		t.Fatalf("body = %q, want %q", got, "study record temporary-id created\n")
	}
}
