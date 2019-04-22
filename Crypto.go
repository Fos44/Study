package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	DenisPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	DenisPublicKey := &DenisPrivateKey.PublicKey

	SerjPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	SerjPublicKey := &SerjPrivateKey.PublicKey


	fmt.Println("Denis PrivateKey:", DenisPrivateKey)
	fmt.Println("Denis PublicKey:", DenisPublicKey)
	fmt.Println("Serj PrivateKey:", SerjPrivateKey)
	fmt.Println("Serj PublicKey:", SerjPublicKey)

	message := []byte("The code must be good")
	label := []byte("")
	hash := sha256.New()


	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, SerjPublicKey, message, label)

	if err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(message),ciphertext)

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // Simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, DenisPrivateKey, newhash, hashed, &opts)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("PSS signature : %x\n", signature)

	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, SerjPrivateKey, ciphertext, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("OAEP decrypted [%x] to \n[%x]\n", ciphertext, plainText)

	err = rsa.VerifyPSS(DenisPublicKey, newhash, hashed, signature, &opts)

	if err != nil {
		fmt.Println("Who are you? Signature failed")
		os.Exit(1)
	}else {
		fmt.Println("Signature succsessful")
	}
}
