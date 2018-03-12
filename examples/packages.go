package examples

import (
	"fmt"
)

// The default Hello World example
var exampleHelloPackage = `package main

import "%s/pkg"

func main() {
	hello.Message()
}
`
var examplePackage = `package hello

import "fmt"

// Message (exported function)
func Message() {

	fmt.Printf("Hello from your package\n")

}
`

func (examples *Examples) addPackage(project string) {
	var hello Example
	hello.Name = "packages"
	hello.addPackageFile("package.go", examplePackage)

	//Create an updated file with correct path to package
	updatedCode := fmt.Sprintf(exampleHelloPackage, examples.BuildPackagePath(project))
	hello.addSourceFile("main.go", updatedCode)

	examples.addExample(hello)
}
