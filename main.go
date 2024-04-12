package main

import "github.com/sheppieboy/p2p-encrypted-messaging/usertypes"


func main(){
	name := "luke"
	port := "5454"
	userProfile := usertypes.NewUserProfile(name, port)

	user := NewUser(userProfile)

	user.P2PFinder.StartP2PDiscovery()
}