package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	haves := []string{
		"foo",
		"bar",
	}

	assert.True(t, contains(haves, "foo"))
}
