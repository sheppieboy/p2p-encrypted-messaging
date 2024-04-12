/*
This file will be the peer finder, it uses udp broadcasting to find peers and broastcast itself so peers can find the user.

On the networking level the broadcaster will broadcast a udp message that contains a unique string to help peers identify themselves.  It will broadcast the message on a frequent basis so
that peers will be quickly discoverable.  To do this the broacaster will read from a udp connection in a udp listener, parse the peer data and add it to the users peerlist along with the relevant
peer information.  The broadcaster will also broadcast the users relevant information so the peers on the network can discover the user and the user to their peer list.

The broadcaster will be a struct that is used by the user struct.  It will use composition to add the functions to broadcast itself, listen for peers, and start itself.
*/
package p2pbroadcasting

import (
	"net"
	"strings"
)


type UniqueBroadcastMessage struct{
	UniqueIdentfier string
	Name string
	Port string
}

func readBroadcastPacketFromUDPConnection(udpConn *net.UDPConn)(*UniqueBroadcastMessage, *net.UDPAddr, error){
	broadcastMessageBuffer := make([]byte, 1024)

	n, addr, err := udpConn.ReadFromUDP(broadcastMessageBuffer)

	if err != nil {
		return nil, nil, err;
	}

	trimmedBytesArr:= broadcastMessageBuffer[:n]; //trim null bytes

	userInfo := strings.Split(string(trimmedBytesArr), ":")

	return &UniqueBroadcastMessage{
		UniqueIdentfier: userInfo[0],
		Name: userInfo[1],
		Port: userInfo[2],
	},addr,nil
}

type P2PFinder struct{
	User string
}

func newP2PFinder(user string)*P2PFinder{
	return &P2PFinder{
		User: user,
	}
}

func (pf *P2PFinder) broadCastToPeers(){}

func (pf *P2PFinder) listenForPeers(){}

func (pf *P2PFinder) startP2PDiscovery(){}






