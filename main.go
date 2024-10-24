package main

import (
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
	case "handshake":
	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
