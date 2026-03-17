package bootstrap

import (
	"context"
	"net/http"

	httpadapter "github.com/boosun13/ArcNote/backend/internal/adapter/http"
	studyrecordadapter "github.com/boosun13/ArcNote/backend/internal/adapter/http/studyrecord"
	application "github.com/boosun13/ArcNote/backend/internal/application/studyrecord"
	domain "github.com/boosun13/ArcNote/backend/internal/domain/studyrecord"
)

type noopStudyRecordRepository struct{}

func (noopStudyRecordRepository) Save(_ context.Context, _ domain.StudyRecord) error {
	return nil
}

func NewHTTPServer() http.Handler {
	studyRecorder := application.NewRecordStudyUseCase(noopStudyRecordRepository{})

	return httpadapter.NewServer(httpadapter.Routes{
		StudyRecord: studyrecordadapter.NewHandler(studyRecorder),
	})
}
