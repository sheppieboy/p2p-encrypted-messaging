package p2pbroadcasting

import (
	"net"
)




type UniqueBroadcastMessage struct{
	UniqueIdentfier string
	Name string
	Port string
}





func readBroadcastPacketFromUDPConnection(conn *net.Conn, bufferSize int)


func AssembleUniqueMessageFromPeers(bytes []byte){
	
}

