package q1

import (
	"flag"
	"io"
	"os"
)

func CopyFile() error {
	strPtr := flag.String("f", "", "test.txt")
	flag.Parse()

	oldFile, err := os.Open(*strPtr)
	if err != nil {
		return err
	}

	defer oldFile.Close()

	newFile, err := os.Create("new.txt")
	if err != nil {
		return err
	}

	defer newFile.Close()

	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		return err
	}

	return nil
}
