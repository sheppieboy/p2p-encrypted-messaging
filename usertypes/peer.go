package usertypes

import "sync"

type Peer struct {
	Name string
	Port string
	PeerIPAddr string 
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
	//handle if peer is already in list
	pl.Peers[p.Name] = p
}