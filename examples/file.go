package examples

// The default Hello World example
var exampleFile = `package main

import "io/ioutil"
	

func main() {
	string := "Hello World for your new project created\n" 

	fileData :=[]byte(string)
	err := ioutil.WriteFile("hello", fileData, 0644)
	if err != nil {
		panic(err)
	}
}
`

func (examples *Examples) addFile(project string) {
	var file Example
	file.Name = "file"
	file.addSourceFile("file.go", exampleFile)
	examples.addExample(file)
}
