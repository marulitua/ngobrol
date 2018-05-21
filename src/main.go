package main

import (
    //"fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
    //fmt.Println("Hello World")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
        var conn, _ = upgrader.Upgrade(w, r, nil)
        go func(conn *websocket.Conn) {
            for {
                mType, msg, _ := conn.ReadMessage()

                conn.WriteMessage(mType, msg)
            }
        }(conn)
    })

    http.ListenAndServe(":3000", nil)//handler http.Handler)
}
