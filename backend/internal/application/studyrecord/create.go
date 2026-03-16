package studyrecord

import domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"

type RecordInput struct {
	DurationMinutes int
	Content         string
	StudiedOn       string
}

type Recorder interface {
	Execute(input RecordInput) (domain.StudyRecord, error)
}

type RecordStudyUseCase struct{}

func NewRecordStudyUseCase() RecordStudyUseCase {
	return RecordStudyUseCase{}
}

func (u RecordStudyUseCase) Execute(input RecordInput) (domain.StudyRecord, error) {
	return domain.New("temporary-id", input.DurationMinutes, input.Content, input.StudiedOn)
}
