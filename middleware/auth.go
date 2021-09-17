package middleware

import (
  "bytes"
  "log"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

type User struct {
  User_id int
}

func GetUser(tokenString string) int {
  postBody, _ := json.Marshal(map[string]string{
    "token": tokenString,
  })
  responseBody := bytes.NewBuffer(postBody)
  resp, err := http.Post("http://shigoto.live/api/v1/jwt/verify/", "application/json", responseBody)
  if err != nil {
      log.Fatalf("An Error Occured %v", err)
   }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
      log.Fatalf("An Error Occured %v", err)
   }
  var result User
  if err := json.Unmarshal(body, &result); err != nil {
      log.Fatalf("An Error Occured %v", err)
  }
  return result.User_id
}
