package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
    //fmt.Println("Hello World")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        if r.Method != "GET" {
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                return
        }

        http.ServeFile(w, r, "public/index.html")
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

    port := ":3000"
    log.Println("Server started on", port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
