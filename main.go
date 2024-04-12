package main

import (
	"github.com/sheppieboy/p2p-encrypted-messaging/user"
	"github.com/sheppieboy/p2p-encrypted-messaging/usertypes"
)


func main(){
	name := "luke"
	userProfile := usertypes.NewUserProfile(name)
	userProfile.DiffieHellman.Print()

	user := user.NewUser(userProfile)
	
	user.Start()
}