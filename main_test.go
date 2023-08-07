package main

import (
	"testing"
	"time"

	"github.com/pedrobertao/go-crud/server"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	require := require.New(t)

	var err error
	go func() {
		err = server.Start(":3030")
	}()

	time.Sleep(2 * time.Second)
	require.Nil(err)

	err = server.Close()
	require.Nil(err)

}
