package examples

var exampleHelloWorld = `package main

import "fmt"

func main() {

	fmt.Printf("Hello World for Project %s created\n")

}
`

func (examples *Examples) addHello() {
	var hello Example
	hello.Name = "hello"
	hello.addFile("hello.go", exampleHelloWorld)
	examples.addExample(hello)
}
