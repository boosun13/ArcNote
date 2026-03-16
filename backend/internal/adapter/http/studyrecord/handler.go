package studyrecord

import (
	"encoding/json"
	"fmt"
	"net/http"

	application "github.com/boosun13/ArcNote/backend/internal/application/studyrecord"
)

type createRequest struct {
	DurationMinutes int    `json:"durationMinutes"`
	Content         string `json:"content"`
	StudiedOn       string `json:"studiedOn"`
}

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req createRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		useCase := application.NewCreateUseCase()
		record, err := useCase.Execute(application.CreateInput{
			DurationMinutes: req.DurationMinutes,
			Content:         req.Content,
			StudiedOn:       req.StudiedOn,
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Location", "/study-records/"+record.ID)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(fmt.Sprintf("study record %s created\n", record.ID)))
	})
}
