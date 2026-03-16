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

func NewHandler(recorder application.Recorder) http.Handler {
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

		record, err := recorder.Execute(application.RecordInput{
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
