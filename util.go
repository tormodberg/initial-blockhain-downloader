package main

import (
	"github.com/btcsuite/btcd/peer"
	"github.com/btcsuite/btclog"
	"os"
)

// The minimal equivalent of the logging setup in btcd/log.go
func initLogging() {
	backend := btclog.NewBackend(os.Stdout)
	peer.UseLogger(backend.Logger("PEER"))
}
