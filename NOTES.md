# Go Towards Greatness!

John Gosset
Apr 15-18 2019

----------

# Docker vim-go Environment

- Download Docker
    - Windows: https://docs.docker.com/docker-for-windows/install/
    - OSX: https://docs.docker.com/docker-for-mac/install/
- Run: docker pull gossetx/go-vim-go

----------

# Recommended Tools

- linter: [golangci-lint](https://github.com/golangci/golangci-lint)
- mkcert: https://github.com/FiloSottile/mkcert

----------

# REQUESTED TOPICS

- Memory leak detection / avoidance
- Writing memory efficient code

----------

# INTERESTING LINKS

- https://gopherize.me
- https://blog.golang.org/profiling-go-programs

## Go Community Practices

- [Go User Survey 2018 Results](https://blog.golang.org/survey2018-results)
- See in particular: https://blog.golang.org/survey2018-results#TOC_6.

## Go Language Info


- [Go Module wiki page](https://github.com/golang/go/wiki/Modules)

- Go Code Review Comments (Best Practices): https://github.com/golang/go/wiki/CodeReviewComments
- Blog illustrating how slices work: https://blog.golang.org/go-slices-usage-and-internals
- Generics: https://research.swtch.com/generic


## Go Future Directions ("Go 2")

- Go 2 Draft Designs: https://go.googlesource.com/proposal/+/master/design/go2draft.md
- Towards Go 2: https://blog.golang.org/toward-go2
- Go 2 here we come: https://blog.golang.org/go2-here-we-come


## Go Design Philosophy

Go Proverbs: https://go-proverbs.github.io/
Talk on Golang's Garbage Collector: https://blog.golang.org/ismmkeynote
Go is a shop-built jig: http://robnapier.net/go-is-a-shop-built-jig


## Writings by Russ Cox (Go's project lead & "Philosopher King")

- Go & Versioning (6 part series): https://research.swtch.com/vgo
- Go modules use Semantic Versioning: https://semver.org/
- Dependency management: https://research.swtch.com/deps


## Optimizing Memory Usage

- The Lost Art of Structure Packing: http://www.catb.org/esr/structure-packing/
- Note the section on Go & Rust: http://www.catb.org/esr/structure-packing/#_go_and_rust


## Shrinking Go Binaries

- Shrink Your Go Binaries With This One Weird Trick: https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

## Misc. Topics

- [Python Unicode Howto](https://docs.python.org/3/howto/unicode.html) --- great as a general guide

# ENVIRONMENT VARIABLES

Cross Compiling

GOOS
GOARCH
    - For example: GOOS=linux GOARCH=arm go build -o hello.arm main.go
    - List of valid GOOS & GOARCH combinations: https://golang.org/doc/install/source#environment
GOPATH
PATH
GO111MODULE


----------

# HELLO WORLD

Write a simple hello world program. Follow along with the following steps:

- Set your go environment to always use the newer module-based approach: "export GO111MODULE=on"
- Create a new directory for your code (drwhello)
- Initialize a new module ("go mod init drw/hello"). This will create your go.mod file.
- Write something similar to the following in your main.go:


    ```go
    package main

    import "github.com/fatih/color"

    func main() {
        color.Cyan("Hello, DRW!")
    }
    ```

- Compile and install your code (this will also install any 3rd party deps): "go install"
- Run your code (type "drwhello" at your command prompt)


----------

# NOTES & EXERCISES

## LANGUAGE FUNDAMENTALS

https://github.com/qjcg/gtg/blob/master/topics/courses/go/language/README.md

# Variables
Notes:
Doc for "fmt": https://golang.org/pkg/fmt/

# Struct Types
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/struct_types/README.md
Exercise: https://play.golang.org/p/h-7BEn2U3Rz

# Pointers
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/pointers/README.md
Exercise 1: https://play.golang.org/p/6QYTKWzF8s8
Exercise 2: https://play.golang.org/p/nolKjrgBX-X

# Constants
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/constants/README.md
Exercise 1: https://play.golang.org/p/4Gs3Ls_5_pi

# Functions
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/functions/README.md
Exercise 1: https://play.golang.org/p/5vEQxEzq3i_D

# Slices
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/slices/README.md

# Interfaces
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/interfaces/README.md
Exercise: https://play.golang.org/p/hfC2-yPI9y6

# Embedding
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/embedding/README.md
Exercise: https://play.golang.org/p/kdHgALCIPIs

----------

# CONCURRENCY

    - [Notes](https://github.com/qjcg/gtg/blob/master/topics/go/concurrency/goroutines/README.md)
    - [Exercise template](https://play.golang.org/p/O0FB2gd6-7d)

## Channels
    - [Exercise 1](https://play.golang.org/p/gv9lxA3qhH-)
    - [Exercise 2](https://play.golang.org/p/9_b6YcBuSOR)

# EXERCISE: Unit Testing

    Write and run ONE OF EACH of the following for either:

    - code you've already written in class
    - 3rd party code you find on the net
    - something similar to what I have in hello.go

    1. Unit test
    2. Benchmark
    3. Example

    To run your tests, run:

    - `go test` (unit tests & examples, add `-v` for details)
- `go test -bench .` (benchmarks)


-----

# EXERCISE: Go Debugger

1. Install Delve, the go debugger: `go get -u github.com/go-delve/delve/cmd/dlv`
2. Choose a Go executable, and inspect it via the debugger:
    `dlv exec /path/to/your/executable` (absolute path or binary name from your PATH)

    # Set a breakpoint at your main() function, and move the cursor there:
    b main.main
    continue

    # Type "help" and review commands"

    # Set breakpoints with "breakpoint (b)"

    # Print values of expressions with "p"

    # Inspect and switch goroutines with "goroutine(s)"

-----

# STDLIB REVIEW

- [io](https://golang.org/pkg/io/)


# AGENDA FOR TODAY

- Creating CLI Tools
- Creating Web servers & clients
- gRPC & protobufs
- scrapers
- Regex
- Networking
- Disadvantages of Go
- Large projects / orgs using Go
- Memory leak detection / avoidance
- Writing memory efficient code

---------

# CLI Tool Exercise

- Take an application that you've already written in class (or adapt the example
  code), and update it to accept at least 3 configuration parameters.

- Use either the simple standard-lib based approach (using "flag"), or consider
  viper for a more complex project (ex: with config file(s) & env vars)

- Examples you can use for inspiration are:

    - cmd/gtg-cli-simple
    - cmd/gtg-cli-viper
    - cmd/gtg-cli-cobra

- Run your CLI tool and make sure that your config params are being used as expected


-----------

# HTTP API Exercise

Write an HTTP(S) API that provides at least 3 endpoints.

Some ideas:

- /accounts/desjardins (a GET might return the details of a bank account)
- /users/5 (a GET might return the details for user id #5)
- /func/squared (might return the square of a POSTed number)

Feel free to use the 3 example apps in the Git repo as inspiration:

- cmd/gtg-http-basic
- cmd/gtg-http-chi-router
- cmd/gtg-http-tls

