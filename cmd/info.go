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

	torrent.Info.Name = decodedFile["info"].(map[string]interface{})["name"].(string)
	torrent.Info.Length = decodedFile["info"].(map[string]interface{})["length"].(int)
	torrent.Info.PieceLength = decodedFile["info"].(map[string]interface{})["piece length"].(int)
	torrent.Info.Pieces = decodedFile["info"].(map[string]interface{})["pieces"].(string)

	return torrent, nil
}
