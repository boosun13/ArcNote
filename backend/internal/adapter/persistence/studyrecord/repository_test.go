package studyrecord

import (
	"context"
	"testing"

	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
)

func TestRepositorySaveStoresStudyRecord(t *testing.T) {
	repository := NewRepository()

	record, err := domain.New("study-record-id", 60, "Go HTTP basics", "2026-03-16")
	if err != nil {
		t.Fatalf("domain.New() returned error: %v", err)
	}

	if err := repository.Save(context.Background(), record); err != nil {
		t.Fatalf("Save() returned error: %v", err)
	}

	if len(repository.records) != 1 {
		t.Fatalf("stored records = %d, want %d", len(repository.records), 1)
	}

	if got := repository.records[record.ID]; got != record {
		t.Fatalf("stored record = %+v, want %+v", got, record)
	}
}
