package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocument_ok(t *testing.T) {
	config := NewDefaultConfig()
	c, err := newDocument(config)
	assert.Nil(t, err)
	assert.Equal(t, config, c.config)
	// TODO assert.NotEmpty on other elements of server struct
	//assert.NotEmpty(t, c.storer)
}

func TestNewDocument_err(t *testing.T) {
	badConfigs := map[string]*Config{
	// TODO add bad config instances
	}
	for desc, badConfig := range badConfigs {
		c, err := newDocument(badConfig)
		assert.NotNil(t, err, desc)
		assert.Nil(t, c)
	}
}

// TODO add TestDocument_ENDPOINT_(ok|err) for each ENDPOINT
