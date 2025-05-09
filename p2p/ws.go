package p2p

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func StartWSServer(port string, handler func(*websocket.Conn)) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Ошибка подключения:", err)
			return
		}
		defer conn.Close()
		handler(conn)
	})

	log.Printf("WebSocket сервер запущен на :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type WebSocketConn = websocket.Conn
