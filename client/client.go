package client

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func Example() {
	const addr = "localhost:8888"

	sendMessages := []string{
		"ASCII",
		"PROGRAMING",
		"PLUS",
	}

	current := 0

	var conn net.Conn = nil

	for {
		var err error

		if conn == nil {
			conn, err = net.Dial("tcp", addr)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Access: %d\n", current)
		}

		request, err := http.NewRequest("POST", addr, strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}

		request.Header.Set("Accept-Encoding", "gzip")

		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))
		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}

			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		current++
		if current == len(sendMessages) {
			break
		}
	}
}
