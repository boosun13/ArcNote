package httpadapter

import (
	"net/http"

	studyrecordadapter "github.com/boosun13/ArcNote/backend/internal/adapter/http/studyrecord"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/study-records", studyrecordadapter.NewHandler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ArcNote API is running\n"))
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	return mux
}
