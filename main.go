package main

import (
	"fmt"
	"os"
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
	case "info":
	case "peers":
	case "handshake":
	}
}
