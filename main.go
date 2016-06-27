package main

import "fmt"

//import crypto "virgil_crypto"
//import "golang.org/x/crypto/curve25519"
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	// "io"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
)

func main() {
	//fmt.Println(crypto.VirgilVersionAsString());
	var private, public [32]byte
	var private1, public1 [32]byte

	var common, common1 [32]byte
	b := make([]byte, 32)
	rand.Read(b)
	copy(private[:], b)

	rand.Read(b)
	copy(private1[:], b)

	curve25519.ScalarBaseMult(&public, &private)
	curve25519.ScalarBaseMult(&public1, &private1)

	curve25519.ScalarMult(&common, &private, &public1)
	curve25519.ScalarMult(&common1, &private1, &public)

	fmt.Println(common)
	fmt.Println(common1)

	dk := pbkdf2.Key(common1[:], common1[:], 4096, 32, sha1.New)
	fmt.Println(dk)

	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("exampleplainqerfhq9re8hfq-8ref-qr7fqr78 ftext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	/*if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	    panic(err.Error())
	}*/

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)
}
