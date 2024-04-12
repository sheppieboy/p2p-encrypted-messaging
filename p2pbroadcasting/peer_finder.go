/*
This file will be the peer finder, it uses udp broadcasting to find peers and broastcast itself so peers can find the user.

On the networking level the broadcaster will broadcast a udp message that contains a unique string to help peers identify themselves.  It will broadcast the message on a frequent basis so
that peers will be quickly discoverable.  To do this the broacaster will read from a udp connection in a udp listener, parse the peer data and add it to the users peerlist along with the relevant
peer information.  The broadcaster will also broadcast the users relevant information so the peers on the network can discover the user and the user to their peer list.

The broadcaster will be a struct that is used by the user struct.  It will use composition to add the functions to broadcast itself, listen for peers, and start itself.
*/
package p2pbroadcasting

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/sheppieboy/p2p-encrypted-messaging/usertypes"
)

const uniqueString = "najdsbajds"

type UniqueBroadcastMessage struct{
	UniqueIdentfier string
	Name string
	Port string
}

func readBroadcastPacketFromUDPConnection(udpConn *net.UDPConn)(*UniqueBroadcastMessage, *net.UDPAddr, error){
	broadcastMessageBuffer := make([]byte, 1024)

	n, senderAddr, err := udpConn.ReadFromUDP(broadcastMessageBuffer)

	if err != nil {
		return nil, nil, err;
	}

	trimmedBytesArr:= broadcastMessageBuffer[:n]; //trim null bytes

	userInfo := strings.Split(string(trimmedBytesArr), ":")

	return &UniqueBroadcastMessage{
		UniqueIdentfier: userInfo[0],
		Name: userInfo[1],
		Port: userInfo[2],
	},senderAddr,nil
}

type P2PFinder struct{
	UserProfile *usertypes.UserProfile
	BroadcastFrequency time.Duration
	BroadcastIP *net.UDPAddr
}

func NewP2PFinder(userProfile *usertypes.UserProfile, frequency time.Duration, broadcastIP *net.UDPAddr) *P2PFinder {
	return &P2PFinder{
		UserProfile: userProfile,
		BroadcastFrequency: frequency,
		BroadcastIP: broadcastIP,
	}
}

func (pf *P2PFinder) broadCastToPeers(){
	udpConn, err := net.DialUDP("udp", nil, pf.BroadcastIP)

	if err != nil{
		log.Fatal(err)
	}
	ticker := time.NewTicker(pf.BroadcastFrequency)

	for {
		<-ticker.C
		broadcastMsg := []byte(fmt.Sprintf("UniqueIdentfier:%s,Name:%s,Port:%s", uniqueString, pf.UserProfile.Name, pf.UserProfile.Port))
		_, err := udpConn.Write(broadcastMsg)
		if err != nil{
			log.Fatal(err)
		}
	}
	
}

func (pf *P2PFinder) listenForPeers(){}

func (pf *P2PFinder) StartP2PDiscovery(){
}






