package jwttoken

import "hello/orthodonticsAdmin/middleware/token"

func CreateToken(claims token.CustomClaims)(string, error)  {
	jwtSign := token.CreateJwt()
	tokenstr, err :=jwtSign.CreateToken(claims)
	if err != nil {
		return "", err
	}
	return tokenstr,  nil
}

func RefreshToken(tokenstr string)string {
	JwtSign := token.CreateJwt()
	str, _ := JwtSign.RefreshToken(tokenstr,60*60)
	return str
}
