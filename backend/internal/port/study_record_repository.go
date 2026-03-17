package port

import (
	"context"

	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
)

type StudyRecordRepository interface {
	Save(ctx context.Context, record domain.StudyRecord) error
}
