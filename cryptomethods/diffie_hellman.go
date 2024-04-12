package cryptomethods

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
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

func (dh *DiffieHellman) SharedSecret(peerPubKey *ecdh.PublicKey) ([]byte){
	sharedSecret, err := dh.PrivateKey.ECDH(peerPubKey)

	if err != nil{
		log.Fatalf("Error in shared secret generation %v:", err)
	}
	return sharedSecret
}

func (dh *DiffieHellman) ToString() string {
	bytesPubKey := dh.PublicKey.Bytes()
	intPubKey := new(big.Int).SetBytes(bytesPubKey)
	strPublicKey := intPubKey.String()
	return strPublicKey
}

func (dh *DiffieHellman) StringToPublicKey(strPublicKey string)(*ecdh.PublicKey, bool){
	intPubKey := new(big.Int)
	intPubKey.SetString(strPublicKey, 10) //base 10
	bytesPubKey := intPubKey.Bytes()
	pubKey, err := dh.PublicKey.Curve().NewPublicKey(bytesPubKey)

	if err != nil{
		log.Fatalf("Invalid Public key %v", err)
	}
	valid := dh.PublicKey.Equal(pubKey)
	return pubKey, valid
}


func (dh *DiffieHellman) Print(){
	pk := dh.PrivateKey.Bytes()
	valid, err := dh.PrivateKey.Curve().NewPrivateKey(pk)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(dh.PrivateKey.Equal(valid))
}
