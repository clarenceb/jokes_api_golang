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

Running the Client
------------------


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

Credits
-------

License
-------
