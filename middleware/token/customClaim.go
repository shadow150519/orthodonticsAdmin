package token

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Id int `json:"id"`
	Username string `json:"username"`
	WorkerId string `json:"worker_id"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}
