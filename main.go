package main

import (
    "flag"
    "net/http"
    "log"
    "github.com/SimeonAleksov/socket-service/ws"
)


func main() {
    flag.Parse()
    http.HandleFunc("/ws/results", ws.Handle)
    http.HandleFunc("/ws/status", ws.HandleStatus)
    log.Println("Serving...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
