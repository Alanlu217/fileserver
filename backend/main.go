package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v/latest/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/v/latest")
		log.Logf(log.InfoLevel, "GET latest: %s", path)
	})

	mux.HandleFunc("POST /v/latest/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("f")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		path := "/root" + strings.TrimPrefix(r.URL.Path, "/v/latest")
		log.Logf(log.InfoLevel, "path: %s", path)

		err = os.WriteFile(path, fileBytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	})

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	log.Logf(log.InfoLevel, "starting on %v", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Logf(log.FatalLevel, "%v", err)
	}
}
