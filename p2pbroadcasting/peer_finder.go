/*
This file will be the peer finder, it uses udp broadcasting to find peers and broastcast itself so peers can find the user.

On the networking level the broadcaster will broadcast a udp message that contains a unique string to help peers identify themselves.  It will broadcast the message on a frequent basis so
that peers will be quickly discoverable.  To do this the broacaster will read from a udp connection in a udp listener, parse the peer data and add it to the users peerlist along with the relevant
peer information.  The broadcaster will also broadcast the users relevant information so the peers on the network can discover the user and the user to their peer list.

The broadcaster will be a struct that is used by the user struct.  It will use composition to add the functions to broadcast itself, listen for peers, and start itself.
*/
package p2pbroadcasting

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/sheppieboy/p2p-encrypted-messaging/usertypes"
)

const uniqueString = "oogaboogie"

type UniqueBroadcastMessage struct{
	UniqueIdentfier string
	Name string
	PublicKey string
}

func readBroadcastPacketFromUDPConnection(udpConn *net.UDPConn)([]byte, *net.UDPAddr, error){
	broadcastMessageBuffer := make([]byte, 1024)

	_, senderAddr, err := udpConn.ReadFromUDP(broadcastMessageBuffer)

	if err != nil {
		return nil, nil, err;
	}

	return broadcastMessageBuffer, senderAddr, nil
}

func getBroadcastedMessageFromPeers(byteArr []byte)(*UniqueBroadcastMessage, error){
	byteArr = bytes.Trim(byteArr, "\x00") //trim null bytes

	strArr := strings.Split(string(byteArr), ":")

	return &UniqueBroadcastMessage{
		UniqueIdentfier: strArr[0],
		Name: strArr[1],
		PublicKey: strArr[2],
	}, nil
}

type P2PFinder struct{
	UserProfile *usertypes.UserProfile
	BroadcastFrequency time.Duration
	BroadcastAddr *net.UDPAddr
}

func NewP2PFinder(userProfile *usertypes.UserProfile, frequency time.Duration, broadcastAddr *net.UDPAddr) *P2PFinder {
	return &P2PFinder{
		UserProfile: userProfile,
		BroadcastFrequency: frequency,
		BroadcastAddr: broadcastAddr,
	}
}

func (pf *P2PFinder) broadCastToPeers(){
	fmt.Println("started broadcasting")
	udpConn, err := net.DialUDP("udp", nil, pf.BroadcastAddr)

	if err != nil{
		log.Fatal(err)
	}
	defer udpConn.Close()

	ticker := time.NewTicker(pf.BroadcastFrequency)

	for {
		<-ticker.C
		broadcastMsg := []byte(fmt.Sprintf("%s:%s:%s:", uniqueString, pf.UserProfile.Name, pf.UserProfile.DiffieHellman.ToString()))
		_, err := udpConn.Write(broadcastMsg)
		if err != nil{
			log.Fatal(err)
		}
	}
	
}

func (pf *P2PFinder) listenForPeers(){
	fmt.Println("start listening")
	udpConn, err := net.ListenMulticastUDP("udp", nil, pf.BroadcastAddr)

	if err != nil{
		log.Fatal(err)
	}

	defer udpConn.Close()

	for {
		packet, senderAddr, err := readBroadcastPacketFromUDPConnection(udpConn)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}

		msg, err := getBroadcastedMessageFromPeers(packet)

		if err != nil {
			log.Fatal(err)
		}

		// Process each packet in a separate goroutine
		go func(addr *net.UDPAddr, msg *UniqueBroadcastMessage) {
			// Process the broadcast message and add the peer to the list
			// Note: Implement the logic to add the peer to the list using a mutex
			// For example:
			// pf.peerList.AddPeer(msg)
			dh := pf.UserProfile.DiffieHellman
			_, valid := dh.StringToPublicKey(msg.PublicKey)
			if(valid){
				log.Fatal("Bad public key")
			}
			log.Printf("Received broadcast message from %s: %s", addr.String(), msg)
		}(senderAddr, msg)
	}
	
}

func (pf *P2PFinder) StartP2PDiscovery(){
	go pf.broadCastToPeers()
	go pf.listenForPeers()
}






