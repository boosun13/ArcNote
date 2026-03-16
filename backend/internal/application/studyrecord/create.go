package studyrecord

import domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"

type CreateInput struct {
	DurationMinutes int
	Content         string
	StudiedOn       string
}

type CreateUseCase struct{}

func NewCreateUseCase() CreateUseCase {
	return CreateUseCase{}
}

func (u CreateUseCase) Execute(input CreateInput) (domain.StudyRecord, error) {
	return domain.New("temporary-id", input.DurationMinutes, input.Content, input.StudiedOn)
}
