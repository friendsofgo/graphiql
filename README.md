#Golang GraphiQL

An utility wrote on Go for using GraphiQL without need to use NPM.

[Try the live demo.](http://graphql.org/swapi-graphql)

## Getting started
To install the library, run:

```bash
go get -u github.com/ubeep/graphiql
``` 

### Use GraphiQL on your own server
If you've a GraphQL server on GO and you want to include the GraphiQL
on it, it's as simple as this:

```go
package main

import(
	"net/http"
	"github.com/ubeep/graphiql"
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
	"github.com/ubeep/graphiql"
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
