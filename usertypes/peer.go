package usertypes

import (
	"net"
	"sync"
)

type Peer struct {
	Name string
	PeerUDPAddr *net.UDPAddr
	
}

type PeerList struct {
	RWMutex *sync.RWMutex //need to manage access to the users peerlist
	Peers map[string]*Peer
}

func NewPeerList()*PeerList{
	return &PeerList{
		RWMutex: &sync.RWMutex{},
		Peers: make(map[string]*Peer),
	}
}

func (pl *PeerList) AddPeer(p *Peer){
	pl.RWMutex.Lock()
	defer pl.RWMutex.Unlock()
	//need to handle if Peer is in peerlist
	pl.Peers[p.Name] = p
}