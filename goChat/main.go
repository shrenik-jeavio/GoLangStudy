package main

import (
	"flag"
	"log"
	"net/http"
	//"github.com/gorilla/websocket"
)


var addr = flag.String("addr", ":8000", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request){
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", 404)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
	}
	http.ServeFile(w, r, "home.html")
}

func main(){
	flag.Parse()
	hub := newHub()

	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}