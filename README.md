Cocktail
========

Inspired by [martini](https://github.com/go-martini/martini), cocktail is a different approach for middleware from martini perspective. Cocktail has all martini's features because it was re-written base on martini. However, compare to martini, cocktail more focus on RESTful and flexibility to extend the framework.

- At the core level, cocktail only allow GET, POST, PATCH, and DELETE. This means that one url path can only have maximum 4 handlers.
- Also, from martini perspective, the whole middleware had been built on top of dependency injection. Cocktail had difference approach, everything should be built from interface, not from concrete. This approach will allow developers to easily replace the cocktail built-in components with theirs custom components.
- Last and not least, Cocktail is 100% compatible with Openshift

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file. We'll call it `server.go`.

~~~ go
package main

import "github.com/phuc0302/go-cocktail"

func main() {
  c := cocktail.Classic()
  c.Get("/", true, func() string {
    return "Hello world!"
  })
  c.Run()
}
~~~

Then install Cocktail package (**go 1.1** and greater is required):
~~~
go get -u github.com/phuc0302/go-cocktail
~~~

Then run your server:
~~~
go run server.go
~~~

You will now have a Cocktail webserver running on `localhost:8080`.
