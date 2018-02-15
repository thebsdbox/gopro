package examples

// The default Hello World example
var exampleHelloWorld = `package main

import "fmt"

func main() {

	fmt.Printf("Hello World for your new project created\n")

}
`

func (examples *Examples) addHello() {
	var hello Example
	hello.Name = "hello"
	hello.addFile("hello.go", exampleHelloWorld)
	examples.addExample(hello)
}
