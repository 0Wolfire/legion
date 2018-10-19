package network

import (
	"github.com/gladiusio/legion/utils"
)

// NewPeer returns a new peer from the given remote. It also
// sets up the reading and writing channels
func NewPeer(remote utils.KCPAddress) *Peer {
	return &Peer{}
}

// Peer is an type that allows easy communication with
// a remote peer
type Peer struct {
	// The remote Address to dial
	remote utils.KCPAddress

	// The internal channel we write to to send a new message
	// to the remote
	sendQueue chan *Message

	// The channel of incoming messages
	recieveChan chan *Message

	// TODO: Need some KCP connection here
}

// QueueMessage queues the specified message to be sent to the remote
func (p *Peer) QueueMessage(m *Message) {
	go func() { p.sendQueue <- m }()
}

// IncomingMessages returns a channel of every message recieved from
// the remote peer
func (p *Peer) IncomingMessages() chan *Message {
	return p.recieveChan
}

// OpenStream dials the remote and opens a stream to the peer
func (p *Peer) OpenStream() error {
	return nil
}

// RecieveStream takes an incoming connection and creates a stream from it
func (p *Peer) RecieveStream() error {
	return nil
}

// Close closes the stream if it exists
func (p *Peer) Close() error {
	return nil
}
