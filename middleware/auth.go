package middleware

import (
  "log"
  "github.com/dgrijalva/jwt-go"
)

func GetUser(tokenString string) int {
  claims := jwt.MapClaims{}
  _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error)        {
          return []byte(""), nil
        })
  if err != nil {
    log.Println(err)
  }
  userId := claims["user_id"].(float64)
  returnVal := int(userId)
  return returnVal
}
