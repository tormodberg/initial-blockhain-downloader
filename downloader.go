package main

import (
	"log"
)

// A Downloader instance will download headers and blocks
// from DownloaderPeer instances until

type Downloader struct {
	peerService PeerService
	downloadState DownloadState
}

func NewDownloader(peerService PeerService, downloadState DownloadState) *Downloader {
	return &Downloader{peerService, downloadState}
}

func (d *Downloader) Run() error {
	highestPeerChannel := make(chan int32, 1)
	d.peerService.RegisterHighestChainChannel(highestPeerChannel)

	err := d.peerService.Start()
	if err != nil {
		return err
	}

	highestPeer := <-highestPeerChannel
	if highestPeer <= d.downloadState.CurrentHeight() {
		log.Printf(
			"This nothing to do as the height of peers (%d) is lower or equal to our local height (%d)",
			highestPeer,
			d.downloadState.CurrentHeight(),
		)
		return nil
	}

	downloaderPeer := d.peerService.DownloaderPeer()
	hash, err := d.downloadState.LatestDownloadedHash()
	_, _ = downloaderPeer.GetHeaders(hash)
	return nil
}
