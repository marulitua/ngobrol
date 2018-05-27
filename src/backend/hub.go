package main

type Hub struct {
    // Incoming message
    broadcast chan []byte

    // List of connected clients
    clients map[*Client]bool

    // register new client
    register chan *Client

    // client terminating connection
    unregister chan *Client
}

func hubGenerator() *Hub {
    return &Hub{
        clients:         make(map[*Client]bool),
        broadcast:       make(chan []byte),
        register:        make(chan *Client),
        unregister:      make(chan *Client),
    }
}

func (h *Hub) watch() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
        case message := <-h.broadcast:
            for client := range h.clients {
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
