package myencrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassWord(password string)(encryptStr string, err error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashAndPassword(hash string, password string)bool{
	err :=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}
