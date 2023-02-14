package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey

	s := "테스트 테스트 테스트 입니다. 테스트에요~~!"

	ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	fmt.Println(string(plaintext))

	hash := md5.New()
	hash.Write([]byte(s))
	digest := hash.Sum(nil)

	var h1 crypto.Hash
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, h1, digest)

	fmt.Println(string(signature))

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15(publicKey, h2, digest, signature)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("검증 성공!")
	}

}
