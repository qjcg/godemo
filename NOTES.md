# Go Towards Greatness!

Instructor: John Gosset
Apr 15-18 2019

----------

# Docker vim-go Environment

 - Download Docker
    - Windows: https://docs.docker.com/docker-for-windows/install/
    - OSX: https://docs.docker.com/docker-for-mac/install/
- Run: docker pull gossetx/go-vim-go


# Recommended Tools

- linter: [golangci-lint](https://github.com/golangci/golangci-lint)

----------

# REQUESTED TOPICS

- Memory leak detection / avoidance
- Writing memory efficient code
- Writing REST APIs

----------

# INTERESTING LINKS

## Go Community Practices

- Go User Survey 2018 Results: https://blog.golang.org/survey2018-results
    - See in particular: https://blog.golang.org/survey2018-results#TOC_6.

## Go Language Info

    Go Module wiki page: https://github.com/golang/go/wiki/Modules

    Go Code Review Comments (Best Practices): https://github.com/golang/go/wiki/CodeReviewComments

    Blog illustrating how slices work: https://blog.golang.org/go-slices-usage-and-internals

    Generics: https://research.swtch.com/generic


## Go Future Directions ("Go 2")

    Go 2 Draft Designs: https://go.googlesource.com/proposal/+/master/design/go2draft.md

    Towards Go 2: https://blog.golang.org/toward-go2

    Go 2 here we come: https://blog.golang.org/go2-here-we-come


## Go Design Philosophy

    Go Proverbs: https://go-proverbs.github.io/

    Talk on Golang's Garbage Collector: https://blog.golang.org/ismmkeynote

    Go is a shop-built jig: http://robnapier.net/go-is-a-shop-built-jig


## Writings by Russ Cox (Go's project lead & "Philosopher King")

    Go & Versioning (6 part series): https://research.swtch.com/vgo

    Go modules use Semantic Versioning: https://semver.org/

    Dependency management: https://research.swtch.com/deps


## Optimizing Memory Usage

    The Lost Art of Structure Packing: http://www.catb.org/esr/structure-packing/

    Note the section on Go & Rust: http://www.catb.org/esr/structure-packing/#_go_and_rust


## Shrinking Go Binaries

    Shrink Your Go Binaries With This One Weird Trick: https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

## Misc. Topics

- [Python Unicode Howto](https://docs.python.org/3/howto/unicode.html) --- great
  as a general guide to 


-----------

# ENVIRONMENT VARIABLES

    Cross Compiling

    GOOS

    GOARCH

    For example: GOOS=linux GOARCH=arm go build -o hello.arm main.go

    List of valid GOOS & GOARCH combinations: https://golang.org/doc/install/source#environment

    GOPATH

    PATH

    GO111MODULE


----------

# HELLO WORLD

Write a simple hello world program. Follow along with the following steps:
    

    Set your go environment to always use the newer module-based approach: "export GO111MODULE=on"

    Create a new directory for your code (drwhello)

    Initialize a new module ("go mod init drw/hello"). This will create your go.mod file.

    Write something similar to the following in your main.go:

      
      package main
      
      import "github.com/fatih/color"
      
      func main() {
          color.Cyan("Hello, DRW!")
      }
      

    Compile and install your code (this will also install any 3rd party deps): "go install"

    Run your code (type "drwhello" at your command prompt)


----------

# NOTES & EXERCISES

LANGUAGE FUNDAMENTALS

https://github.com/qjcg/gtg/blob/master/topics/courses/go/language/README.md

Variables
Notes:
Doc for "fmt": https://golang.org/pkg/fmt/

Struct Types
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/struct_types/README.md
Exercise: https://play.golang.org/p/h-7BEn2U3Rz

Pointers
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/pointers/README.md
Exercise 1: https://play.golang.org/p/6QYTKWzF8s8
Exercise 2: https://play.golang.org/p/nolKjrgBX-X

Constants
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/constants/README.md
Exercise 1: https://play.golang.org/p/4Gs3Ls_5_pi

Functions
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/functions/README.md
Exercise 1: https://play.golang.org/p/5vEQxEzq3i_D

Slices
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/slices/README.md

Interfaces
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/interfaces/README.md
Exercise: https://play.golang.org/p/hfC2-yPI9y6

Embedding
Notes: https://github.com/qjcg/gtg/blob/master/topics/go/language/embedding/README.md
Exercise: https://play.golang.org/p/kdHgALCIPIs

----------

# CONCURRENCY

- [Notes](https://github.com/qjcg/gtg/blob/master/topics/go/concurrency/goroutines/README.md)
- [Exercise template](https://play.golang.org/p/O0FB2gd6-7d)

## Channels
- [Exercise 1](https://play.golang.org/p/gv9lxA3qhH-)
- [Exercise 2](https://play.golang.org/p/9_b6YcBuSOR)
