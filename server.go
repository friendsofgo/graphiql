package graphiql

import (
	"bytes"
	"html/template"
	"net/http"
)

// Handler define the struct of graphiql server
type Handler struct {
	template *template.Template
	Endpoint string
}

// NewGraphiqlHandler preparing the graphiql client to execute on your own server

// The endpoint is the url where you have your graphql api hosted
func NewGraphiqlHandler(endpoint string) (*Handler, error) {
	t, err := preparingTemplate()
	if err != nil {
		return nil, err
	}

	return &Handler{Endpoint: endpoint, template: t}, nil
}

// ServeHTTP is a handler function for resolve graphiql client on servers
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	graphiql := new(bytes.Buffer)

	h.template.Execute(w, data(h.Endpoint))

	w.Header().Set("Content-Type", "text/html")
	w.Write(graphiql.Bytes())
}
