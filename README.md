Jokes API example in Go lang
============================

This is a basic example of a RESTful API which serves jokes from the internet and a commandline client to make requesting jokes easier.

This is my first ever Go program so its probably not going to be idiomatic Go lang and not a good example of a "real life" API.

The main purpose is for me to try out some of the Go concepts I have learnt about in a [Tour of Go](http://tour.golang.org/) and a starting point for a Go lang Dojo at work.

The code will only use packages from the [standard library](http://golang.org/pkg/).

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
    ./client -n 2 -c
    # => 2: Chuck Norris once ate four 30lb bowling balls without chewing.
    # => 1: No one has ever spoken during review of Chuck Norris' code and lived to tell about it.

    # You can time synchronous vs. concurrent requests like so:
    time ./client -n 10
    # => real   0m3.681s
    time ./client -n 10 -c
    # => real   0m0.846s

About `jokes_api/server`
-----------------------

Demonstrates:

* HTTP RESTful API
* JSON encoding / decoding of messages
* Basic logging for debugging / troubleshooting
* Error handling
* TODO: Testing

About `jokes_api/client`
-----------------------

Demonstrates:

* Command-line tool
* Parsing command-line arguments
* Making concurrent HTTP client requests (via goroutines and channels)
* Error handling
* Packages and splitting code over multiple files
* TODO: Testing

Potential Improvements
----------------------

* Add some unit testing with the `testing` package and HTTP testing with `net/http/httptest` package
* Try a BDD-style framework to describe the API, e.g. [Gingko](http://onsi.github.io/ginkgo/) or [GoConvey](http://goconvey.co/)
* Add some additional API endpoints (e.g. request specific joke, save/list/retrieve/delete favourite jokes for a user)
* Add other joke APIs - not just the [Internet Chuck Norris Database](icndb.com/api/)

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
* Go "testing" package:
    * http://golang.org/pkg/testing/
    * https://golang.org/doc/code.html#Testing

Credits
-------

* [The Internet Chuck Norris Database](http://www.icndb.com/api/) - for providing a free public API
* [How I Start - Go](http://howistart.org/posts/go/1) - for a good example of setting up Go and a providing a RESTful API demo

License
-------

The MIT License (MIT)

Copyright (c) 2014 Clarence Bakirtzidis

(see LICENSE file for details)
