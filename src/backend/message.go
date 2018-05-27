package main

import (
  "log"
  "bytes"
  "github.com/gorilla/websocket"
)

type Message struct {
  // the json tag means this will serialize as a lowercased field
  Message string `json:"message"`
  Event string `json:"event"`
  User int `json:"user"`
}

func checkRequest(c *Client) {
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

/*
  messageType, message, err := c.conn.ReadMessage()
  log.Printf("message type %v", messageType)
  if err != nil {
    if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
      log.Printf("error: %v", err)
    }
  }
*/
  message := bytes.TrimSpace(bytes.Replace([]byte(m.Message), newline, space, -1))

  c.hub.broadcast <- message
  log.Printf("Received message: %v", m)
}
