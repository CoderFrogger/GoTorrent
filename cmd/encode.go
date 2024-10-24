package cmd

import (
	"crypto/sha1"
	"fmt"
)

func bencodeInfo(n Info) string {
	return fmt.Sprintf(
		"d6:lengthi%de4:name%d:%s12:piece lengthi%de6:pieces%d:%se",
		n.Length,
		len(n.Name),
		n.Name,
		n.PieceLength,
		len(n.Pieces),
		n.Pieces,
	)
}

func (info Info) HexHash() []byte {
	infoHash := sha1.New()
	bencode := bencodeInfo(info)
	infoHash.Write([]byte(string(bencode)))
	return infoHash.Sum(nil)
}
