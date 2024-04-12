package cryptomethods

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
	"log"
)

type DiffieHellman struct{
	PublicKey *ecdh.PublicKey
	PrivateKey *ecdh.PrivateKey
}

func NewDiffieHellman()*DiffieHellman{
	curve := ecdh.P256()
	privateKey, err := ecdh.Curve.GenerateKey(curve, rand.Reader)

	//pk generation fails
	if err != nil {
		log.Fatalf("Error in Private Key Generation %v", err)
	}

	publicKey := privateKey.PublicKey()
	
	return &DiffieHellman{
		PublicKey: publicKey,
		PrivateKey: privateKey,
	}
}

func (dh *DiffieHellman) Print(){
	fmt.Println(dh.PrivateKey)
	fmt.Println(dh.PublicKey)
}
