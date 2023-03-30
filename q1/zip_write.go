package q1

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func ZipWrite() error {
	f, err := os.Create("test.zip")
	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(f)
	defer zipWriter.Close()

	w, err := zipWriter.Create("test.txt")
	if err != nil {
		return err
	}

	r := strings.NewReader("testestest")

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return nil
}
