package cmd

import (
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

func ReadTorrentFile(torrentFileName string) (Torrent, error) {
	file, err := os.ReadFile(torrentFileName)
	if err != nil {
		return Torrent{}, err
	}

	var torrent Torrent
	decodedFile, _, err := DecodeBenDictionary(string(file), 0)
	if err != nil {
		return Torrent{}, err
	}

	torrent.Announce = decodedFile["announce"].(string)
	torrent.Info = decodedFile["info"].(Info)

	return torrent, nil
}
