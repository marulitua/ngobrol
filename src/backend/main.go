package main

import (
    "log"
    "net/http"
)

func main() {
    hub := hubGenerator()
    go hub.watch()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        if r.Method != "GET" {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        requestUrl := r.URL.Path
        log.Println(requestUrl)

        http.ServeFile(w, r, "public" + requestUrl)
    })

    http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(hub, w, r)
    })

    port := ":3000"
    log.Println("Server started on", port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
