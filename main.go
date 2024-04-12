package main

import (
	"github.com/sheppieboy/p2p-encrypted-messaging/user"
	"github.com/sheppieboy/p2p-encrypted-messaging/usertypes"
)


func main(){
	name := "luke"
	port := "5454"
	userProfile := usertypes.NewUserProfile(name, port)

	user := user.NewUser(userProfile)

	user.Start()
}