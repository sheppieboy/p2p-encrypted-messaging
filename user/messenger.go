package user

import (
	"fmt"
	"log"
	"net"

	"github.com/sheppieboy/p2p-encrypted-messaging/usertypes"
)

type Message struct {
	SenderPubKey string
	PeerMessage string
}

type Messenger struct{
	ListeningUDPAddr *net.UDPAddr
	UserProfile *usertypes.UserProfile
}

func getIPAddress()string{
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatalf("Error getting IP Address: %v", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String()
}

func NewMessenger(userProfile *usertypes.UserProfile) *Messenger {
	IPAddr := getIPAddress()
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(fmt.Sprintf("%s:%s", IPAddr, userProfile.Port)))
	if err !=nil{
		log.Fatal(err)
	}

	return &Messenger{
		ListeningUDPAddr: udpAddr,
		UserProfile: userProfile,
	}
}

func (m *Messenger) ListenForMessages(){
	udpConn, err := net.ListenUDP("udp", m.ListeningUDPAddr)

	if err != nil {
		log.Fatal(err)
	}

	defer udpConn.Close()

	for {
		buffer := make([]byte, 1024)
		_, senderAddr, err := udpConn.ReadFromUDP(buffer)

		if err != nil{
			log.Println("error reading from UDP:", err)
		}

		go func(addr *net.UDPAddr, message []byte){
			fmt.Printf("Received message from %s: %s\n", addr.String(), string(message))
		}(senderAddr, buffer)
	}
}






