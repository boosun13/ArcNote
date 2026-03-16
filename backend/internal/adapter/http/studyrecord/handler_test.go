package studyrecord

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	application "github.com/boosun13/ArcNote/backend/internal/application/studyrecord"
	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
)

type fakeStudyRecorder struct {
	called bool
	input  application.RecordInput
}

func (f *fakeStudyRecorder) Execute(input application.RecordInput) (domain.StudyRecord, error) {
	f.called = true
	f.input = input
	return domain.New("temporary-id", input.DurationMinutes, input.Content, input.StudiedOn)
}

func TestHandlerCreatesStudyRecord(t *testing.T) {
	body := []byte(`{"durationMinutes":60,"content":"Go HTTP basics","studiedOn":"2026-03-16"}`)
	req := httptest.NewRequest(http.MethodPost, "/study-records", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	recorder := &fakeStudyRecorder{}
	NewHandler(recorder).ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}

	if !recorder.called {
		t.Fatal("recorder Execute() was not called")
	}

	if recorder.input != (application.RecordInput{
		DurationMinutes: 60,
		Content:         "Go HTTP basics",
		StudiedOn:       "2026-03-16",
	}) {
		t.Fatalf("input = %+v", recorder.input)
	}

	if got := rec.Header().Get("Location"); got != "/study-records/temporary-id" {
		t.Fatalf("location = %q, want %q", got, "/study-records/temporary-id")
	}

	if got := rec.Body.String(); got != "study record temporary-id created\n" {
		t.Fatalf("body = %q, want %q", got, "study record temporary-id created\n")
	}
}
