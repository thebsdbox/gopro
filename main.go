package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/thebsdbox/gopro/examples"
)

var makefile = `
SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := %s
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 1.0.0
BUILD := ` + "`git rev-parse HEAD`" + `

# Operating System Default (LINUX)
TARGETOS=linux

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -s"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

DOCKERTAG=latest

.PHONY: all build clean install uninstall fmt simplify check run

all: check install

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@echo Building and Installing project
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

docker:
	@GOOS=$(TARGETOS) make build
	@mv $(TARGET) ./dockerfile
	@docker build -t $(TARGET):$(DOCKERTAG) ./dockerfile/
	@rm ./dockerfile/$(TARGET)
	@echo New Docker image created

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

run: install
	@$(TARGET)
`

var readme = `
# %s

## Please complete heading

Additional text

Created on %s
`

var dockerfile = `#Start from Scratch
FROM scratch
#Copy binary
COPY %s %s
#Run binary
CMD ["./%s"]
`

type project struct {
	readme        *bool
	pkg           *bool
	makefile      *bool
	cmd           *bool
	dockerfile    *bool
	name          string
	exampleSource *examples.Example
}

func main() {
	fmt.Println("~~ \033[32mGOPro-ject\033[m ~~")
	var p project

	p.readme = flag.Bool("readme", false, "Create a README.md")
	p.pkg = flag.Bool("pkg", false, "Create a package directory")
	p.makefile = flag.Bool("makefile", false, "Create a 'Makefile'")
	p.cmd = flag.Bool("cmd", false, "Create a cmd directory")
	p.dockerfile = flag.Bool("dockerfile", false, "Create a dockerfile directory and dockerfile to create a project container")

	// Example code
	listExamples := flag.Bool("listexamples", false, "List the names of the embedded examples")
	exampleName := flag.String("example", "", "Build a new project from one of the embedded examples")

	flag.Parse()

	// Read all embedded example names and print them out
	if *listExamples == true {
		examples.Init("")
		a := examples.GetAllExamples()
		fmt.Printf("\n -- Embedded Examples --\n")
		for _, example := range a {
			fmt.Printf("\t%s\n", example.Name)
		}
		os.Exit(0)
	}

	// Check that the project name is the remaining argument, if not print out the errors
	remArgs := flag.Args()
	if len(remArgs) == 0 {
		fmt.Printf("USAGE: %s [options] <project name> \n\n", filepath.Base(os.Args[0]))
		fmt.Printf("'project name' specifies the name of the project/dir that will be created\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Ensure that the project name is added to the struct
	p.name = remArgs[0]

	// Initialise the examples with a project name
	examples.Init(p.name)

	// Build a project from one of the examples in the ~/example folder
	if *exampleName != "" {
		p.exampleSource = examples.GetExample(*exampleName)
		if p.exampleSource == nil {
			fmt.Printf("[Error] Example not found\n")
			os.Exit(1)
		}
	} else {
		// DEFAULT - Use the hello world example
		p.exampleSource = examples.GetExample("hello")
	}

	err := createProject(p)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("To begin move to the new project directory with the command $ cd %s\n", p.name)
}

func createProject(p project) error {

	if p.name == "" {
		return fmt.Errorf("[ERR] No Project name specified")
	}

	var err error
	err = os.Mkdir(p.name, 0766)
	if err != nil {
		return err
	}

	if *p.cmd == true {
		err = os.Mkdir(p.name+"/cmd", 0766)
		if err != nil {
			return err
		}
		fmt.Println("Creating \033[32mcmd\033[m directory")

	}

	if *p.pkg == true {
		err = os.Mkdir(p.name+"/pkg", 0766)
		if err != nil {
			return err
		}
		fmt.Println("Creating \033[32mpkg\033[m directory")
	}

	// Iterate through source files and write them to disk
	err = writeSourceCodeToDisk(p)
	if err != nil {
		return err
	}

	// Iterate through package source files and write them to disk
	err = writePackageSourceCodeToDisk(p)
	if err != nil {
		return err
	}

	if *p.readme == true {
		fmt.Println("Creating \033[32mREADME.md\033[m")
		readmeData := []byte(fmt.Sprintf(readme, p.name, time.Now().Format("2006-01-02 15:04:05")))
		err = ioutil.WriteFile(p.name+"/Readme.md", readmeData, 0644)
		if err != nil {
			return err
		}
	}

	if *p.makefile == true {
		makeData := []byte(fmt.Sprintf(makefile, p.name))
		err = ioutil.WriteFile(p.name+"/Makefile", makeData, 0644)
		if err != nil {
			return err
		}
		fmt.Printf("Creating \033[32mMakefile\033[m\nEnsure the following before running \"make run\"\n\n \033[32mgit init; \\\n git add *; \\\n git commit -m \"My first commit\" \033[m\n\n")
	}

	if *p.dockerfile == true {
		err = os.Mkdir(p.name+"/dockerfile", 0766)
		if err != nil {
			return err
		}
		dockerData := []byte(fmt.Sprintf(dockerfile, p.name, p.name, p.name))
		err = ioutil.WriteFile(p.name+"/dockerfile/dockerfile", dockerData, 0644)
		if err != nil {
			return err
		}
		fmt.Println("Creating \033[32mdockerfile/dockerfile\033[m")
	}
	return nil
}

func writeSourceCodeToDisk(p project) error {
	for _, s := range p.exampleSource.SourceFiles {
		goData := []byte(s.Code)
		err := ioutil.WriteFile(p.name+"/"+s.Filename, goData, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func writePackageSourceCodeToDisk(p project) error {
	if *p.pkg == false {
		fmt.Printf("[ERROR] the -pkg flag wasn't passed\n")
		os.Exit(1)
	}
	for _, s := range p.exampleSource.PackageFiles {
		goData := []byte(s.Code)
		err := ioutil.WriteFile(p.name+"/pkg/"+s.Filename, goData, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
