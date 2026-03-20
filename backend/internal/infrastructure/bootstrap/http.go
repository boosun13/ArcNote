package bootstrap

import (
	"net/http"

	httpadapter "github.com/boosun13/ArcNote/backend/internal/adapter/http"
	studyrecordadapter "github.com/boosun13/ArcNote/backend/internal/adapter/http/studyrecord"
	studyrecordpersistence "github.com/boosun13/ArcNote/backend/internal/adapter/persistence/studyrecord"
	application "github.com/boosun13/ArcNote/backend/internal/application/studyrecord"
)

func NewHTTPServer() http.Handler {
	studyRecorder := application.NewRecordStudyUseCase(studyrecordpersistence.NewRepository())

	return httpadapter.NewServer(httpadapter.Routes{
		StudyRecord: studyrecordadapter.NewHandler(studyRecorder),
	})
}
