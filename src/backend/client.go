package main

import (
    "bytes"
    "time"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
)

const (
    writeWait       = 10 * time.Second

    pongWait        = 10 * time.Second

    pingPeriod      = (pongWait * 9) / 10

    maxMessageSize  = 512
)

var (
    newline = []byte{'\n'}
    space   = []byte{' '} 
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

type Client struct {
    hub *Hub

    conn *websocket.Conn

    send chan []byte
}

type Message struct {
  // the json tag means this will serialize as a lowercased field
  Message string `json:"message"`
  Event string `json:"event"`
  User int `json:"user"`
}

func (c *Client) readPumb() {
    defer func() {
        c.hub.unregister <- c
        log.Println("terminating connection...")
        c.conn.Close()
    }()
    //log.Printf("Set read limit => %d", maxMessageSize)
    c.conn.SetReadLimit(maxMessageSize)
    c.conn.SetReadDeadline(time.Now().Add(pongWait))
    c.conn.SetPongHandler(func (string) error {
        //log.Printf("got pong")
        c.conn.SetReadDeadline(time.Now().Add(pongWait));
        return nil
    })

    for {
        m := Message{}

        // receive a message using the codec
        err := c.conn.ReadJSON(&m)
        if err != nil {
            log.Println("Error reading json.", err)
                if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            return
        }

        log.Printf("Received message: %v", m)
        message := bytes.TrimSpace(bytes.Replace([]byte(m.Message), newline, space, -1))
        c.hub.broadcast <- message
    }
}

func (c *Client) writePumb() {
    ticker := time.NewTicker(pingPeriod)
    defer func() {
        ticker.Stop()
        c.conn.Close()
    }()
    for {
        select {
        case message, ok := <- c.send:
            c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if !ok {
                log.Println("cannot set write deadline", websocket.CloseMessage)
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            response := Message{
                string(message),
                "user.registered",
                4,
            }
            log.Printf("response %v", response)

            if err := c.conn.WriteJSON(&response); err != nil {
                log.Println("cannot write", err)
            }

        case <-ticker.C:
            c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                return
            }
        }
    }
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
    client.hub.register <- client

    go client.writePumb()
    go client.readPumb()
}
