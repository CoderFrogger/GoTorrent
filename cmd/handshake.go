package cmd

import (
	"crypto/rand"
	"fmt"
	"net"
	"os"
)

type TCPHandshake struct {
	Length       byte
	ProtocolName string
	Reserved     [8]byte
	InfoHash     []byte
	PeerID       []byte
}

func CreateTCPHandshakeMessage(handshake TCPHandshake) []byte {
	var msg []byte
	msg = append(msg, handshake.Length)
	msg = append(msg, handshake.ProtocolName...)
	msg = append(msg, handshake.Reserved[:]...)
	msg = append(msg, handshake.InfoHash[:]...)
	msg = append(msg, handshake.PeerID[:]...)
	return msg
}

func ConnectWithPeer(
	peerAdr string,
	msg []byte,
) (net.Conn, TCPHandshake, error) {
	conn, err := net.Dial("tcp", peerAdr)
	if err != nil {
		return nil, TCPHandshake{}, err
	}

	_, err = conn.Write(msg)
	if err != nil {
		return nil, TCPHandshake{}, err
	}

	resp := make([]byte, 68)
	_, err = conn.Read(resp)
	if err != nil {
		return nil, TCPHandshake{}, err
	}

	return conn, TCPHandshake{
		Length:       resp[0],
		ProtocolName: string(resp[1:20]),
		Reserved:     [8]byte{},
		InfoHash:     resp[28:48],
		PeerID:       resp[48:68],
	}, nil
}

func GenerateRandomPeerID() []byte {
	randomBytes := make([]byte, 0)

	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Printf("Error generating PeerID: %s\n", err)
		os.Exit(1)
	}
	return randomBytes
}
