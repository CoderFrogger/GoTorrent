package main

import (
	"fmt"
	"os"
)

func main() {
	switch command := os.Args[1]; command {
	case "decode":
	case "info":
	case "peers":
	case "handshake":
	}
}
