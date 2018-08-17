package server

import (
	api "github.com/elixirhealth/document/pkg/documentapi"
	"google.golang.org/grpc"
)

// Start starts the server and eviction routines.
func Start(config *Config, up chan *Document) error {
	c, err := newDocument(config)
	if err != nil {
		return err
	}

	// start Document aux routines
	// TODO add go x.auxRoutine() or delete comment

	registerServer := func(s *grpc.Server) { api.RegisterDocumentServer(s, c) }
	return c.Serve(registerServer, func() { up <- c })
}
