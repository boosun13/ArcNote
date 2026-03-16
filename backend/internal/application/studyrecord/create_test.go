package studyrecord

import "testing"

func TestCreateStudyRecordUseCaseExecute(t *testing.T) {
	useCase := NewRecordStudyUseCase()

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
}
