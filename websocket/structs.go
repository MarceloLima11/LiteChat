package websocket

import "github.com/gorilla/websocket"

type (
	Client struct {
		hub *Hub

		conn *websocket.Conn

		send chan []byte
	}

	Hub struct {
		clients map[*Client]bool

		broadcast chan []byte

		register chan *Client

		unregister chan *Client
	}
)
