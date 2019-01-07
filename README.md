# Golang GraphiQL
[![GitHub version](https://badge.fury.io/gh/friendsofgo%2Fgraphiql.svg)](https://badge.fury.io/gh/friendsofgo%2Fgraphiql)
[![Build Status](https://travis-ci.org/friendsofgo/graphiql.svg?branch=master)](https://travis-ci.org/friendsofgo/graphiql)
[![codecov](https://codecov.io/gh/friendsofgo/graphiql/branch/master/graph/badge.svg)](https://codecov.io/gh/friendsofgo/graphiql)
[![Go Report Card](https://goreportcard.com/badge/github.com/friendsofgo/graphiql)](https://goreportcard.com/report/github.com/friendsofgo/graphiql)
[![CDNJS](https://img.shields.io/cdnjs/v/graphiql.svg)](https://cdnjs.com/libraries/graphiql)
[![GoDoc](https://godoc.org/graphql.co/graphql?status.svg)](https://godoc.org/github.com/friendsofgo/graphiql)

<p align="center">
  <img src="http://bit.ly/2zjRTXF" alt="Golang GraphiQL"/>
</p>

An utility wrote on Go for using GraphiQL without need to use NPM.

[Try the live demo.](http://graphql.org/swapi-graphql)

## Getting started
To install the library, run

```bash
go get -u github.com/friendsofgo/graphiql
``` 

To build graphi's CLI you must run
```bash
make build-cli
```
This will generate a binary in the project's bin directory called graphiql. You can then move this binary to anywhere in your `PATH`.

### Use GraphiQL on your own server
If you've a GraphQL server on GO and you want to include the GraphiQL
on it, it's as simple as this:

```go
package main

import(
	"net/http"
	"github.com/friendsofgo/graphiql"
)

func main() {
	
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	
	http.HandleFunc("/graphql", myGraphQLHandler)
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":8080", nil)
}
```

### Use GraphiQL standalone
If you want launch GraphiQL for use on your GraphQL API but you don't want
include it in your own project you can launch your own GraphiQL Server:

```go
package main

import(
	"net/http"
	"github.com/friendsofgo/graphiql"
)

func main() {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("http://localhost:8080/graphql")
	if err != nil {
    		panic(err)
	}
    	
	
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":4040", nil)
}
```

## CLI
Yo can use `graphiql` as command
```bash
$ graphiql --help
usage: graphiql [-version] [OPTIONS]
        graphiql is an standalone server to use Graphiql, based on https://github.com/graphql/graphiql

  -endpoint string
        Endpoint where are hosted your API, eg: http://localhost:8080/graphql (default "http://localhost:8080/graphql")
  -graphiql
        Show the js client of GraphiQL version
  -port string
        Port where will be launched the GraphiQL client (default "4040")
  -version
        Show the graphiql version information
```

And it'll launch a `GraphiQL` server
```bash
$ graphiql

2018/11/24 17:10:36 GraphiQL run on http://localhost:4040/graphiql...
2018/11/24 17:10:36 GraphQL endpoint http://localhost:8080/graphql
2018/11/24 17:10:36 Use (ctrl+c) to terminate the execution
```

## Docker
The application is also available through [Docker](https://hub.docker.com/r/friendsofgo/graphiql), just run:

```bash
docker run -it --rm -p 4040:4040 friendsofgo/graphiql
```

Remember to use the [-p](https://docs.docker.com/engine/reference/run/) flag to expose the container port where the application is listening (4040 by default).

## Contribute
[Contributions](https://github.com/friendsofgo/graphiql/issues?q=is%3Aissue+is%3Aopen) are more than welcome, if you are interested please fork this repo and send your Pull Request.

## License
MIT License, see [LICENSE](https://github.com/friendsofgo/graphiql/blob/master/LICENSE)
