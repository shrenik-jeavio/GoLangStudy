package main

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// register client
	clients map[*Client]bool

	// Inbound message from client
	broadcast chan []byte

	//register requests from the clients
	register chan *Client

	//unregister requests from the clients
	unregister chan *Client
}


func newHub() *Hub {
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for{
		select{
			case client := <-h.register:
				h.clients[client] = true
			case client := <-h.unregister:
				if _, ok := h.clients[client]; ok{
					delete(h.clients, client)
					close(client.send)
				}
			case message := <-h.broadcast:
				for client := range h.clients{
					select {
						case client.send <- message:
						default:
							close(client.send)
							delete(h.clients, client)

					}
				}
		}
	}
}