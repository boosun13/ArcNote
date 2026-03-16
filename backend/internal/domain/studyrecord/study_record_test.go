package studyrecord

import "testing"

func TestNewCreatesStudyRecord(t *testing.T) {
	record, err := New("temporary-id", 60, "Go HTTP basics", "2026-03-16")
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}

	if record.ID != "temporary-id" {
		t.Fatalf("ID = %q, want %q", record.ID, "temporary-id")
	}

	if record.DurationMinutes != 60 {
		t.Fatalf("DurationMinutes = %d, want %d", record.DurationMinutes, 60)
	}

	if record.Content != "Go HTTP basics" {
		t.Fatalf("Content = %q, want %q", record.Content, "Go HTTP basics")
	}

	if record.StudiedOn != "2026-03-16" {
		t.Fatalf("StudiedOn = %q, want %q", record.StudiedOn, "2026-03-16")
	}
}

func TestNewRejectsInvalidStudyRecord(t *testing.T) {
	_, err := New("", 0, "", "")
	if err == nil {
		t.Fatal("New() error = nil, want non-nil")
	}
}
