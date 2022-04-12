package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

const key = "AxRwdPuI3ZWonhtB"

func main() {
	e := enCrypt([]byte("123456789"), []byte(key))
	//r_base := base64.StdEncoding.EncodeToString(e)
	fmt.Printf("%v \n", e)
	//d_base, _ := base64.StdEncoding.DecodeString(r_base)
	d := deCrypt(e, []byte(key))
	fmt.Printf("%v \n", d)
}
func enCrypt(text []byte, key []byte) []byte {
	const iv = "ForgPLNMkzt1kI2B"
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	count := []byte(iv)
	blockMode := cipher.NewCTR(block, count)
	message := make([]byte, len(text))
	blockMode.XORKeyStream(message, text)
	return message
}


func deCrypt(text []byte, key []byte) []byte {
	const iv = "ForgPLNMkzt1kI2B"
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	count := []byte(iv)
	blockMode := cipher.NewCTR(block, count)
	message := make([]byte, len(text))
	blockMode.XORKeyStream(message, text)
	return message
}