package server

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func Example() {
	listner, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}

		go prcessSession(conn)
	}
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func prcessSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("timeout")
				break
			}

			if errors.Is(err, io.EOF) {
				break
			}

			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		response := http.Response{
			StatusCode: http.StatusOK,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		if isGZipAcceptable(request) {
			content := "Hello World\n"

			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello World\n"
			response.Body = io.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}

		response.Write(conn)
	}
}
