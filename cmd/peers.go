package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type TrackerResponse struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func DiscoverPeers(torrent Torrent) (TrackerResponse, error) {
	query := url.Values{}
	query.Add(("info_hash"), string(torrent.Info.HexHash()))
	query.Add("peer_id", string(GenerateRandomPeerID()))
	query.Add("port", "6881")
	query.Add("uploaded", "0")
	query.Add("downloaded", "0")
	query.Add("left", strconv.Itoa(torrent.Info.Length))
	query.Add("compact", "1")

	response, err := http.Get(torrent.Announce + "?" + query.Encode())
	if err != nil {
		return TrackerResponse{}, fmt.Errorf(
			"Error getting response from peer: %s\n",
			err,
		)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return TrackerResponse{}, fmt.Errorf(
			"Error reading response: %s\n",
			err,
		)
	}

	buffer, _, err := DecodeBenDictionary(string(body), 0)
	if err != nil {
		return TrackerResponse{}, fmt.Errorf(
			"Error decoding dictionary: %s\n",
			err,
		)
	}

	var tracker TrackerResponse
	tracker.Interval = buffer["interval"].(int)
	tracker.Peers = buffer["peers"].(string)

	return tracker, nil
}
