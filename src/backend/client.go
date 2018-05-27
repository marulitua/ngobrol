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

func (c *Client) readPumb() {
    defer func() {
        c.hub.unregister <- c
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
        _, message, err := c.conn.ReadMessage()
        //log.Printf("message size %T", c.conn.readRemaining)
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            break
        }
        message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
        log.Printf("got message => %s", message)
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
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            w, err := c.conn.NextWriter(websocket.TextMessage)
            if err != nil {
                return
            }
            w.Write(message)

            n := len(c.send)
            for i := 0; i < n; i++ {
                w.Write(newline)
                w.Write(<-c.send)
            }

            if err := w.Close(); err != nil {
                return
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
