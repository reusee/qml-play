package main

/*
void start();

#cgo pkg-config: Qt5Gui Qt5Quick
*/
import "C"
import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var httpPort int

func main() {

	var ln net.Listener
	for {
		httpPort = 30000 + rand.Intn(20000)
		var err error
		ln, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", httpPort))
		if err != nil {
			continue
		}
		break
	}

	http.HandleFunc("/qmldir", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, `
    module Main
    `)
	})
	http.HandleFunc("/main.qml", func(w http.ResponseWriter, req *http.Request) {
		greetings := req.URL.Query().Get("hello")
		if greetings == "" {
			greetings = "hello"
		}
		io.WriteString(w, `
    import QtQuick 2.13
    Text {
      text: `+fmt.Sprintf("%q", greetings)+`
      font.pixelSize: 500
    }
    `)
	})
	go func() {
		server := &http.Server{}
		if err := server.Serve(ln); err != nil {
			panic(err)
		}
	}()
	for {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d", httpPort))
		if err != nil {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		resp.Body.Close()
		break
	}

	C.start()
}
