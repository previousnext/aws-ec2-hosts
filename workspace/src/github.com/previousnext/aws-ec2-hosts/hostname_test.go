package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostname(t *testing.T) {
	name, err := hostname("{{ .Name }}-bar", "foo")
	assert.Nil(t, err)
	assert.Equal(t, "foo-bar", name)

	name, err = hostname("{{ .Name }}", "foo")
	assert.Nil(t, err)
	assert.Equal(t, "foo", name)
}
