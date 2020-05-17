package main

import "testing"

type mockPeerService struct {}
func (*mockPeerService) Start() error { return nil }
func (*mockPeerService) Stop() error { return nil }
func (*mockPeerService) DownloaderPeer() DownloaderPeer {return nil}

var _ PeerService = (*mockPeerService)(nil)

func TestNewDownloader(t *testing.T) {
	downloader := NewDownloader(&mockPeerService{})
	downloader.Run()
}
