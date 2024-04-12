package usertypes

import "sync"

type Peer struct {
	Name string
	Port string
	Addr string 
}

type PeerList struct {
	RWMutex *sync.RWMutex //need to manage access to the users peerlist
	Peers map[string]*Peer
}

func newPeerList()*PeerList{
	return &PeerList{
		RWMutex: &sync.RWMutex{},
		Peers: make(map[string]*Peer),
	}
}

func (pl *PeerList) addPeer(p *Peer){
	pl.RWMutex.Lock()
	defer pl.RWMutex.Unlock()
	pl.Peers[p.Name] = p
}