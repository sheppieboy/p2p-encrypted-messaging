/*
This file will be the peer finder, it uses udp broadcasting to find peers and broastcast itself so peers can find the user.

On the networking level the broadcaster will broadcast a udp message that contains a unique string to help peers identify themselves.  It will broadcast the message on a frequent basis so
that peers will be quickly discoverable.  To do this the broacaster will read from a udp connection in a udp listener, parse the peer data and add it to the users peerlist along with the relevant
peer information.  The broadcaster will also broadcast the users relevant information so the peers on the network can discover the user and the user to their peer list.

The broadcaster will be a struct that is used by the user struct.  It will use composition to add the functions to broadcast itself, listen for peers, and start itself.
*/