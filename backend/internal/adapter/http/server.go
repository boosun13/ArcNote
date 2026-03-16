package httpadapter

import (
	"net/http"
)

type Routes struct {
	StudyRecord http.Handler
}

func NewServer(routes Routes) http.Handler {
	mux := http.NewServeMux()
	if routes.StudyRecord != nil {
		mux.Handle("/study-records", routes.StudyRecord)
	}
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
