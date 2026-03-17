package studyrecord

import (
	"context"
	"errors"
	"testing"

	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
)

type studyRecordRepositorySpy struct {
	called bool
	record domain.StudyRecord
	err    error
}

func (s *studyRecordRepositorySpy) Save(_ context.Context, record domain.StudyRecord) error {
	s.called = true
	s.record = record
	return s.err
}

func TestCreateStudyRecordUseCaseExecute(t *testing.T) {
	repository := &studyRecordRepositorySpy{}
	useCase := NewRecordStudyUseCase(repository)

	record, err := useCase.Execute(RecordInput{
		DurationMinutes: 60,
		Content:         "Go HTTP basics",
		StudiedOn:       "2026-03-16",
	})
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}

	if record.ID != "temporary-id" {
		t.Fatalf("ID = %q, want %q", record.ID, "temporary-id")
	}

	if record.Content != "Go HTTP basics" {
		t.Fatalf("Content = %q, want %q", record.Content, "Go HTTP basics")
	}

	if !repository.called {
		t.Fatal("repository Save() was not called")
	}

	if repository.record != record {
		t.Fatalf("saved record = %+v, want %+v", repository.record, record)
	}
}

func TestCreateStudyRecordUseCaseReturnsRepositoryError(t *testing.T) {
	repository := &studyRecordRepositorySpy{err: errors.New("save failed")}
	useCase := NewRecordStudyUseCase(repository)

	_, err := useCase.Execute(RecordInput{
		DurationMinutes: 60,
		Content:         "Go HTTP basics",
		StudiedOn:       "2026-03-16",
	})
	if err == nil {
		t.Fatal("Execute() error = nil, want non-nil")
	}
}
