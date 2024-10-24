package main

import (
	"encoding/json"
	"fmt"
	"os"

	"GoTorrent/cmd"
)

type Torrent struct {
	Announce string `bencode:"announce"`
	Info     Info   `bencode:"info"`
}

type Info struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

func main() {
	switch command := os.Args[1]; command {
	case "decode":
		benString := os.Args[2]

		decodedInput, _, err := cmd.DecodeBencode(benString, 0)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decodedInput)
		fmt.Println(string(jsonOutput))

	case "info":
	case "peers":
	case "handshake":
	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
