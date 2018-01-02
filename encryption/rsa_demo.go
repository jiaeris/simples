package encryption

import (
	"encoding/pem"
	"errors"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"fmt"
)

func EnDecryptDemo(data []byte) ([]byte) {

	priBlock, _ := pem.Decode(PKCS1WithPassPrivateBytes)
	if priBlock == nil {
		fmt.Println(errors.New("private key error"))
	}

	pubBlock, _ := pem.Decode(PKCS1WithPassPublicBytes)
	if pubBlock == nil {
		fmt.Println(errors.New("public key error"))
	}

	//generate public key
	pub, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	pubKey := pub.(*rsa.PublicKey)
	//generate private key
	//with password
	priBlockBytes, err := x509.DecryptPEMBlock(priBlock, []byte("1203"))
	if err != nil {
		fmt.Println(err)
	}
	pri, err := x509.ParsePKCS1PrivateKey(priBlockBytes)
	//with out password
	//pri, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	//pri, err := x509.ParsePKCS8PrivateKey(priBlock.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	priKey := pri /*.(*rsa.PrivateKey)*/

	//encryption
	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("encrypt: ", cipherData)
	//decryption
	plainData, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherData)
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypt: ", string(plainData))

	return cipherData
}

func RSADriverDemo() {
	rd := NewRSA()
	err := rd.InitByByte(PKCS1WithPassPrivateBytes, PKCS1WithPassPublicBytes, []string{"1203"})
	if err != nil {
		fmt.Println(err)
	}
	cipherData, err := rd.Encrypt([]byte("hello world"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cipherData)
	plainData, err := rd.Decrypt(cipherData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plainData))

}
