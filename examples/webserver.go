package examples

// examplewebserver - this is a simple webserving application
var exampleWebserverCode = `package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}`

func (examples *Examples) addWebServer(project string) {
	var webServer Example
	webServer.Name = "webserver"
	webServer.addSourceFile("webserver.go", exampleWebserverCode)
	examples.addExample(webServer)
}
