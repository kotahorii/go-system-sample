package q1

import (
	"crypto/rand"
	"io"
	"os"
)

func WriteRandomSize() error {
	f, err := os.Create("new.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = io.CopyN(f, rand.Reader, 1024)
	if err != nil {
		return err
	}

	return nil
}
