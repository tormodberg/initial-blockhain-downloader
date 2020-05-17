package main

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockDownloaderPeer struct {

}

func (m *mockDownloaderPeer) Close() error {
	return nil
}

func (m *mockDownloaderPeer) GetHeaders(startingPoint *chainhash.Hash) (error, []wire.BlockHeader) {
	return nil, nil
}

var _ DownloaderPeer = (*mockDownloaderPeer)(nil)

type mockPeerService struct {
	highestChainChannel chan<- int32
	downloaderPeerCallCount int
	highestChainCount int32
}
func (m *mockPeerService) Start() error {
	m.highestChainChannel <- m.highestChainCount
	return nil
}
func (m *mockPeerService) Stop() error { return nil }
func (m *mockPeerService) DownloaderPeer() DownloaderPeer {
	m.downloaderPeerCallCount += 1
	return &mockDownloaderPeer{}
}
func (m *mockPeerService) RegisterHighestChainChannel(highestChainChanel chan<- int32) {
	m.highestChainChannel = highestChainChanel
}
var _ PeerService = (*mockPeerService)(nil)

type mockDownloadState struct {}
func (*mockDownloadState) LatestDownloadedHash() (*chainhash.Hash, error) {return nil, nil}
func (*mockDownloadState) CurrentHeight() int32 {return 42}
var _ DownloadState = (*mockDownloadState)(nil)

func TestNewDownloader_AlreadyDone(t *testing.T) {
	// Given
	mockPeerService := &mockPeerService{highestChainCount: 42}
	downloader := NewDownloader(mockPeerService, &mockDownloadState{})

	// When
	err := downloader.Run()

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 0, mockPeerService.downloaderPeerCallCount)
}

func TestNewDownloader_NotDone(t *testing.T) {
	// Given
	mockPeerService := &mockPeerService{highestChainCount: 47}
	downloader := NewDownloader(mockPeerService, &mockDownloadState{})

	// When
	err := downloader.Run()

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 1, mockPeerService.downloaderPeerCallCount)
}
