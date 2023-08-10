package main

import (
	"os"
	"testing"
	"time"

	"github.com/pedrobertao/go-crud/logging"
	"github.com/pedrobertao/go-crud/server"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	logging.Start()
	code := m.Run()
	os.Exit(code)
}

func TestMainServer(t *testing.T) {
	logging.Start()
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
