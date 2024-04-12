package user

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/sheppieboy/p2p-encrypted-messaging/p2pbroadcasting"
	"github.com/sheppieboy/p2p-encrypted-messaging/usertypes"
)


const (
	BroadcastIP = "224.0.0.1"
	Frequency = 1 * time.Second
)

type User struct{
	UserProfile *usertypes.UserProfile
	P2PFinder *p2pbroadcasting.P2PFinder
}

func NewUser(userProfile *usertypes.UserProfile) *User {
	broadcastAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", BroadcastIP, userProfile.Port))
	if err != nil{
		log.Fatal(err)
	}

	return &User{
		UserProfile: userProfile,
		P2PFinder: p2pbroadcasting.NewP2PFinder(userProfile, Frequency, broadcastAddr),
	}
}

func (u *User) Start(){
	go u.P2PFinder.StartP2PDiscovery()
	select{}
}
