package middleware

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
)

func GetUser(tokenString string) int {
  claims := jwt.MapClaims{}
  _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error)        {
          return []byte("<YOUR VERIFICATION KEY>"), nil
        })
  if err != nil {
    fmt.Println(err)
  }
  userId := claims["user_id"].(float64)
  returnVal := int(userId)
  return returnVal
}
