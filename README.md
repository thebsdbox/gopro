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

### Using your new project

Once you've created a new project simple move to your new directory `cd <project_name>` and if you've passed the `--makefile` flag then you'll need to initialise youre project with `git`:

- Initialise new repo -> `git init`
- Add initial files to repository -> `git add <file names or *>`
- Intitial commit -> `git commit -m "commit message"`

**Makefile usage**

By default running `make` (defaults to `make install`) will compile and install your the binary from your new project in `$GOPATH/bin`, alternatively `make build` and `make run` will do the obvious. 