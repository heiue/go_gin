package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Crypt(text []byte, key []byte) []byte {
	fmt.Println(text)
	const iv = "ForgPLNMkzt1kI2B"
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	count := []byte(iv)
	blockMode := cipher.NewCTR(block, count)
	message := make([]byte, len(text))
	fmt.Println(message)
	blockMode.XORKeyStream(message, text)
	fmt.Println(message)
	return message
}