# GOPROject

A utility for creating the bare-bones needed to get going on a GO project:

## Usage

**Install**

`go get github.com/thebsdbox/gopro/`

### Create a new project

This will create a new folder with a simple source code file:

```
gopro example1
~~ GOPRO-ject ~~
To begin move to the new project directory with the command $ cd example1
```

This will create a new folder complete with `Makefile` and source code:

```
gopro --makefile example2
~~ GOPRO-ject ~~
Creating Makefile - Ensure the following before running "make run"
 git init; \
 git add *; \
 git commit -m "My first commit" 

To begin move to the new project directory with the command $ cd example2
```

This will create a new folder complete with `README.md`, `Makefile` and `cmd/pkg` directories:

```
gopro --makefile --readme --cmd --pkg example3
~~ GOPRO-ject ~~
Creating README.md
Creating Makefile - Ensure the following before running "make run"
 git init; \
 git add *; \
 git commit -m "My first commit" 

To begin move to the new project directory with the command $ cd example3
```

This will create a new folder complete with `Makefile` and a `./dockerfile` directory

```
gopro --makefile --dockerfile example4
~~ GOPro-ject ~~
Creating Makefile - Ensure the following before running "make run"

 git init; \
 git add *; \
 git commit -m "My first commit" 

Creating dockerfile/dockerfile
To begin move to the new project directory with the command $ cd example4
```

### Examples
There are now a few embedded examples that will create a skeleton project with some basic functions for a few use cases. To view the embedded examples use the command `gopro -listexamples` and when building a new project just specify the example that you would like to create with the `-example <example_name>` flag. The Examples can be found in the `/example` folder as part of the source tree.


### Using your new project

Once you've created a new project simple move to your new directory `cd <project_name>` and if you've passed the `--makefile` flag then you'll need to initialise youre project with `git`:

- Initialise new repo -> `git init`
- Add initial files to repository -> `git add <file names or *>`
- Intitial commit -> `git commit -m "commit message"`

**Makefile usage**

By default running `make` (defaults to `make install`) will compile and install your the binary from your new project in `$GOPATH/bin`, alternatively `make build` and `make run` will do the obvious. 

**Dockerfile usage**

If you have specified the `--dockerfile` flag then the project can easily be automated to create a docker image from your new project. Use the `make docker` command and the project will be built and automated into a `scratch` container.

### NOTES

A `make build` will create a binary that is compatible with your local machine (i.e. the place where the code is being compiled). However a `make docker` will set the binary to be build as a Linux (ELF) binary, this is to ease development in places like Docker4Mac etc.. 

To change this the Makefile can be edited by changing TARGETOS to either `darwin` / `linux` / `win`
