package main

// A Downloader instance will download headers and blocks
// from DownloaderPeer instances until

type Downloader struct {
	peerService PeerService
	downloadState DownloadState
}


func NewDownloader(peerService PeerService, downloadState DownloadState) *Downloader {
	return &Downloader{peerService, downloadState}
}

func (d *Downloader) Run() {
	d.peerService.Start()
	downloaderPeer := d.peerService.DownloaderPeer()


	downloaderPeer.GetHeaders()
}
