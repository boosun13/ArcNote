package port

import (
	"context"
	"testing"

	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
)

type studyRecordRepositoryStub struct{}

func (studyRecordRepositoryStub) Save(_ context.Context, _ domain.StudyRecord) error {
	return nil
}

func TestStudyRecordRepositoryDefinesSaveContract(t *testing.T) {
	var repository StudyRecordRepository = studyRecordRepositoryStub{}

	record, err := domain.New("study-record-id", 60, "Go HTTP basics", "2026-03-16")
	if err != nil {
		t.Fatalf("domain.New() returned error: %v", err)
	}

	if err := repository.Save(context.Background(), record); err != nil {
		t.Fatalf("Save() returned error: %v", err)
	}
}
