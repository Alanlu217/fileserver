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

	mux.HandleFunc("GET /f/{filepath...}", func(w http.ResponseWriter, r *http.Request) {
		filepath := r.PathValue("filepath")
		http.ServeFile(w, r, fs.GetCurrFilePath(filepath))
	})

	mux.HandleFunc("POST /f/{filepath...}", func(w http.ResponseWriter, r *http.Request) {
		filepath := r.PathValue("filepath")

		r.ParseMultipartForm(1 << 20)
		file, handler, err := r.FormFile("f")
		if err != nil {
			log.Error(err)
			return
		}
		defer file.Close()

		err = fs.Upload(file, filepath)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		log.Info("Upload Success", "file", handler.Filename, "path", filepath)
	})

	server := http.Server{
		Addr:    Flags.address,
		Handler: mux,
	}

	log.Logf(log.InfoLevel, "starting on %v", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Logf(log.FatalLevel, "%v", err)
	}
}
