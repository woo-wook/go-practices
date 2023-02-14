package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := "QWERTYUIOPASDFGH"
	s := "테스트 테스트 테스트 문구입니다 테스트 문구!"

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block, ciphertext)
	fmt.Println(string(plaintext))
}

func encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(b, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("오류!")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(b, iv)

	mode.CryptBlocks(plaintext, ciphertext)
	return plaintext
}
