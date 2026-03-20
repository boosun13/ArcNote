package studyrecord

import (
	"context"
	"sync"

	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
	"github.com/boosun13/ArcNote/backend/internal/port"
)

var _ port.StudyRecordRepository = (*Repository)(nil)

type Repository struct {
	mu      sync.RWMutex
	records map[string]domain.StudyRecord
}

func NewRepository() *Repository {
	return &Repository{
		records: make(map[string]domain.StudyRecord),
	}
}

func (r *Repository) Save(_ context.Context, record domain.StudyRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.records[record.ID] = record
	return nil
}
