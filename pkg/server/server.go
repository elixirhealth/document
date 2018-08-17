package server

import (
	"github.com/elixirhealth/document/pkg/server/storage"
	"github.com/elixirhealth/service-base/pkg/server"
)

// Document implements the DocumentServer interface.
type Document struct {
	*server.BaseServer
	config *Config

	storer storage.Storer
	// TODO maybe add other things here
}

// newDocument creates a new DocumentServer from the given config.
func newDocument(config *Config) (*Document, error) {
	baseServer := server.NewBaseServer(config.BaseConfig)
	storer, err := getStorer(config, baseServer.Logger)
	if err != nil {
		return nil, err
	}
	// TODO maybe add other init

	return &Document{
		BaseServer: baseServer,
		config:     config,
		storer:     storer,
		// TODO maybe add other things
	}, nil
}

// TODO implement documentapi.Document endpoints
