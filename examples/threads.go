package examples

// This will create a simple threads example
var exampleThreads = `package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		go thread(i)
	}
	time.Sleep(1 * time.Second)
}

func thread(threadNumber int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Microsecond)

	fmt.Printf("Thread %d waited %d ms\n", threadNumber, r)
}
`

func (examples *Examples) addThreads(project string) {
	var hello Example
	hello.Name = "threads"
	hello.addSourceFile("threads.go", exampleThreads)
	examples.addExample(hello)
}
