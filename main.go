package main

import (
    "flag"
    "net/http"
    "log"
    "github.com/SimeonAleksov/socket-service/ws"
)


func main() {
    flag.Parse()
    http.HandleFunc("/results", ws.Handle)
    http.HandleFunc("/status", ws.HandleStatus)
    log.Println("Serving...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
