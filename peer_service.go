package main

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// PeerService is a high level abstraction on top of the btcd primitives. It maintains knowledge of
// p2p peers and once it has started it can provide DownloaderPeer instances that can be used to send
// and receive messages to.
type PeerService interface {
	Start() error
	Stop() error

	// Returns a DownloaderPeer that can receive and send messages
	DownloaderPeer() DownloaderPeer
	// The user of this interface can register a channel that will notify of the highest chain height
	// that has been received from peers.
	RegisterHighestChainChannel(chan<- int32)
}

// A higher level abstraction around Peer
type DownloaderPeer interface {
	Close() error
	GetHeaders(startingPoint *chainhash.Hash) (error, []wire.BlockHeader)
}
