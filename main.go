package main

import "github.com/go-system-test/q1"

func main() {
	err := q1.ZipWrite()
	if err != nil {
		panic(err)
	}
}
