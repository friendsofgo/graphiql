package graphiql

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

const graphqlEndpoint = "/graphql"

func TestGraphiqlHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/graphiql", nil)
	if err != nil {
		t.Fatal(err)
	}

	graphiqlHandler, err := NewGraphiqlHandler(graphqlEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(graphiqlHandler.ServeHTTP)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response if the correct html
	tmpl, _ := preparingTemplate()
	var expected bytes.Buffer
	tmpl.Execute(&expected, data(graphqlEndpoint))

	if rr.Body.String() != expected.String() {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected.String())
	}

}
