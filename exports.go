package main

import "C"
import "fmt"

//export ServerAddr
func ServerAddr() *C.char {
	return C.CString(fmt.Sprintf("http://localhost:%d/main.qml", httpPort))
}
