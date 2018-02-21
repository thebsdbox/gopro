package examples

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Examples - consists of the name of the example and the demonstration code
type Examples struct {
	Example []Example
}

// Example - Contains the basic structure for an example
type Example struct {
	Name         string
	SourceFiles  []Source
	PackageFiles []Source
}

// Source - An array of source files that are added to an example
type Source struct {
	Filename string
	Code     string
}

var examples Examples

func (examples *Examples) lastExample() int {
	return len(examples.Example)
}

func (example *Example) addSourceFile(filename string, code string) {
	example.SourceFiles = append(example.SourceFiles, Source{filename, code})
}

func (example *Example) addPackageFile(filename string, code string) {
	example.PackageFiles = append(example.PackageFiles, Source{filename, code})
}

func (examples *Examples) addExample(newExample Example) {
	examples.Example = append(examples.Example, newExample)
}

// Init - This will initialise all of the examples and set any per project settings
func Init(project string) {
	examples.addWebServer(project)
	examples.addHello(project)
	examples.addPackage(project)
}

// BuildPackagePath - This will build the correct path to be used for any imports
func (examples *Examples) BuildPackagePath(projectName string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentPath := strings.Replace(dir, os.Getenv("GOPATH")+"/src/", "", 1)
	packagePath := fmt.Sprintf("%s/%s", currentPath, projectName)
	return packagePath
}

// GetAllExamples returns all of the examples
func GetAllExamples() []Example {
	return examples.Example
}

// GetExample - Returns a specific example if found
func GetExample(exampleName string) *Example {
	for _, e := range examples.Example {
		if e.Name == exampleName {
			return &e
		}
	}
	return nil
}
