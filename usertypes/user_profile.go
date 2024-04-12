package usertypes

import "github.com/sheppieboy/p2p-encrypted-messaging/cryptomethods"

type UserProfile struct {
	Name string
	DiffieHellman cryptomethods.DiffieHellman
}

func NewUserProfile(name string) * UserProfile{
	return &UserProfile{
		Name: name,
		DiffieHellman: *cryptomethods.NewDiffieHellman(),
	}
}