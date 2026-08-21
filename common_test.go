package goztl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpointOptionsSupportsField(t *testing.T) {
	eo := EndpointOptions{
		Actions: map[string]map[string]EndpointField{
			"POST": {
				"source": {Type: "field"},
			},
		},
	}

	supported, known := eo.SupportsField("POST", "source")
	assert.True(t, known)
	assert.True(t, supported)

	supported, known = eo.SupportsField("POST", "yolo")
	assert.True(t, known)
	assert.False(t, supported)

	// without the permission for a method, Zentral leaves it out, and the fields it accepts are
	// unknown — not absent
	supported, known = eo.SupportsField("PUT", "source")
	assert.False(t, known)
	assert.False(t, supported)
}
