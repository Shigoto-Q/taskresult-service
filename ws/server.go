package ws

import (
    "net/http"
    "regexp"
    "strings"
    "time"
    "log"
    "golang.org/x/net/websocket"
    "github.com/SimeonAleksov/socket-service/config"
    "github.com/SimeonAleksov/socket-service/middleware"
)


var token string
var user_id int
var connectionUpgradeRegex = regexp.MustCompile("(^|.*,\\s*)upgrade($|\\s*,)")
var db = config.SetupDb()

func isWebsocketRequest(req *http.Request) bool {
    token = req.URL.Query().Get("token")
    user_id = middleware.GetUser(token)
    return connectionUpgradeRegex.MatchString(strings.ToLower(req.Header.Get("Connection"))) && strings.ToLower(req.Header.Get("Upgrade")) == "websocket"
}

func Handle(w http.ResponseWriter, r *http.Request) {
    if isWebsocketRequest(r) {
        websocket.Handler(handleWebSockets).ServeHTTP(w, r)
    }
    log.Println("Finished sending response...")
}

func handleWebSockets(ws *websocket.Conn) {
  for {
    results := config.FetchResults(user_id, db)
    err := websocket.JSON.Send(ws, results)
    if err != nil {
        return
  }
    time.Sleep(time.Second) 
  }
}

