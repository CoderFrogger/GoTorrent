package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"GoTorrent/cmd"
)

func main() {
	switch command := os.Args[1]; command {
	case "decode":
		benString := os.Args[2]

		decodedInput, _, err := cmd.DecodeBencode(benString, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		jsonOutput, _ := json.Marshal(decodedInput)
		fmt.Println(string(jsonOutput))

	case "info":
		torrentFileName := os.Args[2]

		decodedTorrent, err := cmd.ReadTorrentFile(torrentFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Tracker URL: %v\n", decodedTorrent.Announce)
		fmt.Printf("Length: %v\n", decodedTorrent.Info.Length)
		fmt.Printf("Info Hash: %x\n", decodedTorrent.Info.HexHash())
		fmt.Printf("Piece Length: %v\n", decodedTorrent.Info.PieceLength)
		fmt.Printf("Piece Hashes: \n")

		piecesHashes := []byte(decodedTorrent.Info.Pieces)
		for i := 0; i <= len(piecesHashes)-20; i += 20 {
			fmt.Printf("%x\n", piecesHashes[i:i+20])
		}

	case "peers":
		torrentFileName := os.Args[2]

		decodedTorrent, err := cmd.ReadTorrentFile(torrentFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		trackerResponse, err := cmd.DiscoverPeers(decodedTorrent)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		peers := make([]string, 0)
		for i := 0; i < len(trackerResponse.Peers); i += 6 {
			peerIP := fmt.Sprintf(
				"%d.%d.%d.%d",
				trackerResponse.Peers[i],
				trackerResponse.Peers[i+1],
				trackerResponse.Peers[i+2],
				trackerResponse.Peers[i+3],
			)
			peerPort := int(
				trackerResponse.Peers[i+4],
			)<<8 | int(
				trackerResponse.Peers[i+5],
			)
			peerAdr := fmt.Sprintf("%s:%d", peerIP, peerPort)

			peers = append(peers, peerAdr)
			fmt.Printf("%s\n", peerAdr)
		}

	case "handshake":
		torrentFileName := os.Args[2]

		decodedTorrent, err := cmd.ReadTorrentFile(torrentFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		peerAddress := os.Args[3]
		if peerAddress == "" {
			fmt.Println("Peer ip and port required")
			return
		}

		tcpHandshake := cmd.CreateTCPHandshakeMessage(cmd.TCPHandshake{
			Length:       byte(19),
			ProtocolName: "BitTorrent protocol",
			Reserved:     [8]byte{},
			InfoHash:     decodedTorrent.Info.HexHash(),
			PeerID:       cmd.GenerateRandomPeerID(),
		})

		conn, tcpResponse, err := cmd.ConnectWithPeer(peerAddress, tcpHandshake)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer conn.Close()
		fmt.Printf("Peer ID: %s\n", hex.EncodeToString(tcpResponse.PeerID))

	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
