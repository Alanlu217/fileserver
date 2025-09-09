package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	ParseFlags()
	fs, err := NewFs(Flags.root)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("root=%v", Flags.root)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /f/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")
		http.ServeFile(w, r, fs.GetCurrPath(path))
	})

	mux.HandleFunc("POST /f/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")

		r.ParseMultipartForm(1 << 20)
		file, handler, err := r.FormFile("f")
		if err != nil {
			log.Error(err)
			return
		}
		defer file.Close()

		err = fs.Upload(file, path)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		log.Info("Upload Success", "file", handler.Filename, "path", path)
	})

	mux.HandleFunc("DELETE /f/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")
		if err := fs.Delete(path); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	server := http.Server{
		Addr:    Flags.address,
		Handler: mux,
	}

	log.Logf(log.InfoLevel, "starting on %v", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Errorf("%v", err)
	}
}
