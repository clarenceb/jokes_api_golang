Jokes API example in Go lang
============================

This is a basic example of a RESTful API which serves jokes from the internet and a client to make requesting jokes easier.

This is my first ever Go program so its probably not going to be idiomatic Go lang and not a good example of a "real life" API.

The main purpose is for me to try out some of the Go concepts I have learnt about in a [Tour of Go](http://tour.golang.org/) and a starting point for a Go lang Dojo at work.

The code will only use packages from the [standard library](http://golang.org/pkg/).

TODO:
-----
* Build server
* Figure out how to write tests (yes, not doing TDD as I don't even know the syntax yet!)
* Build client

Setup
-----

Running the Server
------------------

    cd src/github.com/clarenceb/jokes_api/server
    go build
    ./server

    # Test server with curl:
    curl -H "Accept: application/json" http://localhost:8080/joke
    # => {"joke":"Chuck Norris finished World of Warcraft."}

Running the Client
------------------

    cd src/github.com/clarenceb/jokes_api/client
    go build
    # TODO: pending...

About `jokes_api/server`
-----------------------

Demonstrates:

* HTTP RESTful API
* JSON encoding / decoding of messages
* Basic logging for debugging / troubleshooting
* Testing
* Making concurrent HTTP client requests (via goroutines and channels)
* Packages and splitting code over multiple files
* Error handling

About `jokes_api/client`
-----------------------

Demonstrates:

* Command-line tool
* Parsing command-line arguments
* Basic logging for debugging / troubleshooting
* Testing
* Making concurrent HTTP client requests (via goroutines and channels)

Potential Improvements
----------------------

Further Learning
----------------

* Learn more about commandline flags:
    * http://golang.org/pkg/flag/
    * https://gobyexample.com/command-line-flags
* Learn more about JSON encoding and decoding:
    * http://blog.golang.org/json-and-go

Credits
-------

* [The Internet Chuck Norris Database](http://www.icndb.com/api/) - for providing a free public API
* [How I Start - Go](http://howistart.org/posts/go/1) - for a good example of setting up Go and a providing a RESTful API demo

License
-------

The MIT License (MIT)

Copyright (c) 2014 Clarence Bakirtzidis

(see LICENSE file for details)
