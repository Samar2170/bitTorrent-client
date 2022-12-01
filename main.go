package main

import (
	"log"

	"github.com/Samar2170/bitTorrent-client/torrentfile"
)

func main() {
	// inPath := os.Args[1]
	// outPath := os.Args[2]
	inPath := "debian-11.5.0-amd64-netinst.iso.torrent"
	outPath := "debian-11.5.0-amd64-netinst.iso"

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
