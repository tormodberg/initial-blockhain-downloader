package main

import "github.com/btcsuite/btcd/chaincfg/chainhash"

type DownloadState interface {
	LatestDownloadedHash() (*chainhash.Hash, error)

}
