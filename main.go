package main

import (
  "context"
  "fmt"
  "log"
  "net/http"
  "github.com/georgysavva/scany/pgxscan"
  "github.com/gorilla/websocket"
  "github.com/jackc/pgx/v4"
  "github.com/SimeonAleksov/socket-service/config"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}

func wsStart(w http.ResponseWriter, r *http.Request) {
  upgrader.CheckOrigin = func(r *http.Request) bool { return true }
   ws, err := upgrader.Upgrade(w, r, nil)
   if err != nil {
     log.Println(err)
   }

   log.Println("client successfully connected")
   err = ws.WriteMessage(1, []byte("Hello from the other side"))
   if err != nil {
     log.Println(err)
   }
}

func setup_routes() {
  http.HandleFunc("/", wsStart)
}

func main() { 
    conn, err := pgx.Connect(context.Background(), config.DBUrl)
    if err != nil {
      log.Println(err)
    }
    defer conn.Close(context.Background())
    var taskResults []*config.TaskResult

    err = pgxscan.Select(context.Background(), conn, &taskResults, config.Query)
  
    if err != nil {
      log.Println(err)
    }
    fmt.Printf("%v\n", taskResults)
    fmt.Println("Starting server")
    setup_routes()
    log.Fatal(http.ListenAndServe(":5000", nil))

}
