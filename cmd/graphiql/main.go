package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/friendsofgo/graphiql"
)

var (
	name    = "graphiql"
	version = "undefined"
	graphiqlEndpoint = "/graphiql"
)

func main() {
	flag.Usage = usage
	endpoint := flag.String("endpoint", "http://localhost:8080/graphql", "Endpoint where are hosted your API, eg: http://localhost:8080/graphql")
	port := flag.String("port", "4040", "Port where will be launched the GraphiQL client")
	showVersion := flag.Bool("version", false, "Show the graphiql version information")
	showClientGraphiQLVersion := flag.Bool("graphiql", false, "Show the js client of GraphiQL version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s %s\n", name, version)
		return
	}

	if *showClientGraphiQLVersion {
		fmt.Println(graphiql.GraphiQLVersion)
		return
	}

	run(*endpoint, *port)
}

func run(endpoint, port string) {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler(endpoint)
	if err != nil {
		panic(err)
	}

	http.Handle(graphiqlEndpoint, graphiqlHandler)
	httpAddr := ":" + port

	log.Printf("GraphiQL run on http://localhost%s%s...\n", httpAddr, graphiqlEndpoint)
	log.Printf("GraphQL endpoint %s\n", endpoint)
	log.Printf("Use (ctrl+c) to terminate the execution")

	if err := http.ListenAndServe(httpAddr, nil); err != nil {
		panic(err)
	}
}

func usage() {
	msg := fmt.Sprintf(`usage: %s [-version] [OPTIONS]
	%s is an standalone server to use Graphiql, based on https://github.com/graphql/graphiql
	`, name, name)
	fmt.Println(msg)
	flag.PrintDefaults()
}
