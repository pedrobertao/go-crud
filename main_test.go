package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromTestify(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	assert.True(true)
	require.True(true)
}
