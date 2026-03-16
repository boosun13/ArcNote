package studyrecord

import "errors"

type StudyRecord struct {
	ID              string
	DurationMinutes int
	Content         string
	StudiedOn       string
}

func New(id string, durationMinutes int, content string, studiedOn string) (StudyRecord, error) {
	if id == "" {
		return StudyRecord{}, errors.New("study record id is required")
	}
	if durationMinutes <= 0 {
		return StudyRecord{}, errors.New("study duration must be positive")
	}
	if content == "" {
		return StudyRecord{}, errors.New("study content is required")
	}
	if studiedOn == "" {
		return StudyRecord{}, errors.New("study date is required")
	}

	return StudyRecord{
		ID:              id,
		DurationMinutes: durationMinutes,
		Content:         content,
		StudiedOn:       studiedOn,
	}, nil
}
