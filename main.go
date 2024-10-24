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

		// decodedTorrent, err := cmd.ReadTorrentFile(torrentFileName)

	case "peers":
	case "handshake":
	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
