//This file contains code Encryption
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

const (
	KEYBITS = 1024
)

//GenerateRsaKeyPair, Generate Key Pair returns PEM Encoded private key and Public key
func GenerateRsaKeyPair() ([]byte, []byte) {
	pr, err := rsa.GenerateKey(rand.Reader, KEYBITS)
	checkerr(err)
	var b = pem.Block{}
	b.Type = "PRIVATE KEY"
	b.Bytes = x509.MarshalPKCS1PrivateKey(pr)
	var bb = pem.Block{}
	bb.Type = "PUBLIC KEY"
	bb.Bytes = x509.MarshalPKCS1PublicKey(&pr.PublicKey)
	return pem.EncodeToMemory(&b), pem.EncodeToMemory(&bb)
}

//WritePrivateKeyTOPemFile, This function create's private key pem file in given folder , name of file is determine by hash of public key corrosponding to private key
func WritePrivateKeyTOPemFile(pr, pu []byte, folder string) []byte {
	n := sha256.New()
	b, _ := pem.Decode(pu)
	n.Write(b.Bytes)
	h := n.Sum(nil)

	f, _ := os.Create(fmt.Sprintf("%s/%s.pem", folder, hex.EncodeToString(h))) //Write's file in name of hash of public key
	f.Write(pr)
	return h
}
func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
