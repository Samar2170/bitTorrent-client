package client

import (
	"net"

	"github.com/Samar2170/bitTorrent-client.git/bitfield"
	"github.com/Samar2170/bitTorrent-client/peers"
)

type Client struct {
	Conn     net.Conn
	Choked   bool
	BitField bitfield.BitField
	peer     peers.Peer
	infoHash [20]byte
	peerID   [20]byte
}
