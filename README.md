Jokes API example in Go lang
============================

This is a basic example of a RESTful API which serves jokes from the internet and a client to make requesting jokes easier.

This is my first ever Go program so its probably not going to be idiomatic Go lang and not a good example of a "real life" API.

The main purpose is for me to try out some of the Go concepts I have learnt about in a [Tour of Go](http://tour.golang.org/) and a starting point for a Go lang Dojo at work.

The code will only use packages from the [standard library](http://golang.org/pkg/).

TODO:
-----
* Figure out how to write tests (yes, not doing TDD as I don't even know the syntax yet!)
    * Go testing package
    * Gingko - BDD
    * GoConvery - BDD

Setup
-----

Install Go compiler and tools: https://golang.org/doc/install

On Mac OS X, you can use [Homebrew](http://brew.sh/):

    brew install go

Prepare your Go workspace and set your [GOPATH](https://golang.org/doc/code.html#GOPATH):

    $ mkdir -p $HOME/go/{src,pkg,bin}
    $ export GOPATH=$HOME/go
    $ echo "export GOPATH=$HOME/go" >> ~/.bashrc

Read more about Go lang [workspaces](https://golang.org/doc/code.html#Workspaces).

Fetch this project into your Go workspace:

    $ cd $GOPATH
    $ go get github.com/clarenceb/jokes_api_golang

Using the Server
----------------

    # Build the server
    cd $GOPATH/src/github.com/clarenceb/jokes_api/server
    go build

    # Run the server (on port localhost:8080)
    ./server

    # Test server with curl:
    curl -H "Accept: application/json" http://localhost:8080/joke
    # => {"joke":"Chuck Norris finished World of Warcraft."}

Using the Client
----------------

    # Build the client
    cd $GOPATH/src/github.com/clarenceb/jokes_api/client
    go build

    # View usage
    ./client --help

    # Fetch one joke
    ./client
    # => 1: Chuck Norris finished World of Warcraft.

    # Fetch multiple jokes
    ./client -n 2
    # => 1: No one has ever spoken during review of Chuck Norris' code and lived to tell about it.
    # => 2: Chuck Norris once ate four 30lb bowling balls without chewing.

    # Fetch multiple jokes concurrently
    ./client -n 2
    # => 2: Chuck Norris once ate four 30lb bowling balls without chewing.
    # => 1: No one has ever spoken during review of Chuck Norris' code and lived to tell about it.

About `jokes_api/server`
-----------------------

Demonstrates:

* HTTP RESTful API
* JSON encoding / decoding of messages
* Basic logging for debugging / troubleshooting
* Error handling
* Testing

About `jokes_api/client`
-----------------------

Demonstrates:

* Command-line tool
* Parsing command-line arguments
* Making concurrent HTTP client requests (via goroutines and channels)
* Error handling
* Packages and splitting code over multiple files
* Testing

Potential Improvements
----------------------

Further Learning
----------------

* Start here: [Tour of Go](http://tour.golang.org/)
* Understand why Go was made and its design goals:
    * [Go at Google: Language Design in the Service of Software Engineering](http://talks.golang.org/2012/splash.article)
* Quick reference to learn syntax and common usage:
    * https://gobyexample.com/
* Learn more about commandline flags:
    * http://golang.org/pkg/flag/
    * https://gobyexample.com/command-line-flags
* Learn more about JSON encoding and decoding:
    * http://blog.golang.org/json-and-go
* Per-project GOPATH setup with [direnv](http://tammersaleh.com/posts/manage-your-gopath-with-direnv/)

Credits
-------

* [The Internet Chuck Norris Database](http://www.icndb.com/api/) - for providing a free public API
* [How I Start - Go](http://howistart.org/posts/go/1) - for a good example of setting up Go and a providing a RESTful API demo

License
-------

The MIT License (MIT)

Copyright (c) 2014 Clarence Bakirtzidis

(see LICENSE file for details)
