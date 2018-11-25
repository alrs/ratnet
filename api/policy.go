package api

import "github.com/awgh/bencrypt/bc"

// Policy : defines a "Connection Policy" object
type Policy interface {
	RunPolicy() error
	Stop()
	MarshalJSON() (b []byte, e error)
	GetTransport() Transport
}

// PeerInfo - last contact info for peers
type PeerInfo struct {
	LastPollLocal  int64
	LastPollRemote int64
	TotalBytesTX   int64
	TotalBytesRX   int64
	RoutingPub     bc.PubKey
}
