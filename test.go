package main

import (
	"fmt"

	"github.com/Samar2170/bitTorrent-client/torrentfile"
)

func testTorrentFile() {
	tf, err := torrentfile.Open("debian-11.5.0-amd64-netinst.iso.torrent")
	if err != nil {
		panic(err)
	}
	println(tf.Announce)
	fmt.Printf("%x\n", tf.InfoHash)
	// fmt.Printf("%x", tf.PieceHashes)
	fmt.Printf("%d", tf.PieceLength)
	fmt.Printf("%d\n", tf.Length)
	fmt.Printf("%s\n", tf.Name)

	peers, err := tf.RequestPeers([20]byte{}, torrentfile.Port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", peers)
}
