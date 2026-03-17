package studyrecord

import (
	"context"

	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
	"github.com/boosun13/ArcNote/backend/internal/port"
)

type RecordInput struct {
	DurationMinutes int
	Content         string
	StudiedOn       string
}

type Recorder interface {
	Execute(input RecordInput) (domain.StudyRecord, error)
}

type RecordStudyUseCase struct {
	repository port.StudyRecordRepository
}

func NewRecordStudyUseCase(repository port.StudyRecordRepository) RecordStudyUseCase {
	return RecordStudyUseCase{repository: repository}
}

func (u RecordStudyUseCase) Execute(input RecordInput) (domain.StudyRecord, error) {
	record, err := domain.New("temporary-id", input.DurationMinutes, input.Content, input.StudiedOn)
	if err != nil {
		return domain.StudyRecord{}, err
	}

	if err := u.repository.Save(context.Background(), record); err != nil {
		return domain.StudyRecord{}, err
	}

	return record, nil
}
