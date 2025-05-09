package p2p

import (
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
)

func NewNode() host.Host {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", node.ID())
	fmt.Println("Адреса:", node.Addrs())

	return node
}

func HandleStream(stream network.Stream) {
	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Получено: %s\n", string(buf[:n]))
	stream.Close()
}
