package main

import (
	"fmt"
	"log"

	"github.com/Uso1and/p2p-chat/crypto"
	"github.com/Uso1and/p2p-chat/p2p"
	"github.com/Uso1and/p2p-chat/wal"
)

func main() {

	if err := wal.InitDB(); err != nil {
		log.Fatal("WAL error:", err)
	}

	_, pubKey, err := crypto.GenerateKeyPair()
	if err != nil {
		log.Fatal("Keygen error:", err)
	}
	fmt.Printf("Public key: %x\n", pubKey)

	node := p2p.NewNode()
	defer node.Close()

	go p2p.StartWSServer("8080", func(conn *p2p.WebSocketConn) {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			fmt.Printf("Received: %s\n", msg)
		}
	})

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
