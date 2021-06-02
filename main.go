package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/websocket"
  "github.com/SimeonAleksov/socket-service/config"
  "github.com/SimeonAleksov/socket-service/middleware"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}
var User_id int 


func main() { 
    fmt.Println("Starting server")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

      upgrader.CheckOrigin = func(r *http.Request) bool { return true }
      ws, err := upgrader.Upgrade(w, r, nil)
       if err != nil {
         log.Println(err)
      }
       log.Println("client successfully connected")
       err = ws.WriteMessage(1, []byte("Hello from the other side"))
       token := r.URL.Query().Get("token")
       db := config.SetupDb()
       User_id = middleware.GetUser(token)
       config.FetchResults(User_id, db)

       if err != nil {
         log.Println(err)
       }    
       defer func() {
         ws.Close()
       }()
   },
 )
 log.Fatal(http.ListenAndServe(":5000", nil))

}
