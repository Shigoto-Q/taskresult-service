package main

import (
  "log"
  "net/http"
  "encoding/json"
  "time"
  "github.com/gorilla/websocket"
  "github.com/SimeonAleksov/socket-service/config"
  "github.com/SimeonAleksov/socket-service/middleware"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}


func main() { 
    log.Println("Starting server")
    db := config.SetupDb()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

      upgrader.CheckOrigin = func(r *http.Request) bool { return true }
      ws, err := upgrader.Upgrade(w, r, nil)
       if err != nil {
         log.Println(err)
      }
       log.Println("client successfully connected")
       token := r.URL.Query().Get("token")
       user_id := middleware.GetUser(token)
       users := config.FetchResults(user_id, db)
       usersJson, err := json.Marshal(users)
       if err != nil {
         log.Println(err)
       }
       err = ws.WriteMessage(1, usersJson)
       if err != nil {
         log.Println(err)
       }
       func() {
         for{
             time.Sleep(time.Second)
             users := config.FetchResults(user_id, db)
             usersJson, err := json.Marshal(users)
             if err != nil {
               log.Println(err)
             }
             err = ws.WriteMessage(1, usersJson)
             if err != nil {
               log.Println(err)
             }
           }
       }()
       defer ws.Close()
   },
 )
 log.Fatal(http.ListenAndServe(":5000", nil))

}
