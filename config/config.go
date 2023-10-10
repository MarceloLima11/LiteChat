package config

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/MarceloLima11/LiteChat/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

func Init() {
	flag.Parse()

	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 4 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
