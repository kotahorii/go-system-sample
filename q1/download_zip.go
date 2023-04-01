package q1

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func DownloadZip() error {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=test.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}

	reader := strings.NewReader("testestetstest")

	_, err = io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
}
