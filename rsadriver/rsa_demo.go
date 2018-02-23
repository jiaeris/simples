package rsadriver

//
//how to use it
//
//import (
//	"encoding/pem"
//	"errors"
//	"crypto/x509"
//	"crypto/rsa"
//	"crypto/rand"
//	"fmt"
//)
//
//func EnDecryptDemo(data []byte) ([]byte) {
//
//	priBlock, _ := pem.Decode(PKCS1WithPassPrivateBytes)
//	if priBlock == nil {
//		fmt.Println(errors.New("private key error"))
//	}
//
//	pubBlock, _ := pem.Decode(PKCS1WithPassPublicBytes)
//	if pubBlock == nil {
//		fmt.Println(errors.New("public key error"))
//	}
//
//	//generate public key
//	pub, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
//	if err != nil {
//		fmt.Println(err)
//	}
//	pubKey := pub.(*rsa.PublicKey)
//	//generate private key
//	//with password
//	priBlockBytes, err := x509.DecryptPEMBlock(priBlock, []byte("1203"))
//	if err != nil {
//		fmt.Println(err)
//	}
//	pri, err := x509.ParsePKCS1PrivateKey(priBlockBytes)
//	//with out password
//	//pri, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
//	//pri, err := x509.ParsePKCS8PrivateKey(priBlock.Bytes)
//	if err != nil {
//		fmt.Println(err)
//	}
//	priKey := pri /*.(*rsa.PrivateKey)*/
//
//	//encryption
//	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("encrypt: ", cipherData)
//	//decryption
//	plainData, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherData)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("decrypt: ", string(plainData))
//
//	return cipherData
//}
//
//const (
//	rootPath       = "./pems"
//	pkcs1AesPriKey = rootPath + "/aes_private_key.pem"
//	pkcs1AesPubKey = rootPath + "/aes_public_key.pem"
//
//	pkcs8PriKey = rootPath + "/pkcs8_private_key.pem"
//	pkcs8PubKey = rootPath + "/pkcs8_public_key.pem"
//
//	pkcs1NoPassPriKey = rootPath + "/private_key_nopass.pem"
//	pkcs1NoPassPubKey = rootPath + "/public_key_nopass.pem"
//)
//
//func RSADriverDemo() {
//	rd := NewRSA().InitByByte(PKCS1WithPassPrivateBytes, PKCS1WithPassPublicBytes, "1203")
//	cipherData, err := rd.Encrypt([]byte("hello world"))
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(cipherData)
//	plainData, err := rd.Decrypt(cipherData)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(plainData))
//
//	//use pem file
//	rd2 := NewRSA()
//	//rd2.InitByFile(pkcs1AesPriKey, pkcs1AesPubKey, "1203")
//	rd2.SetPKCSVersion(PKCS8)
//	rd2.InitByFile(pkcs8PriKey, pkcs8PubKey)
//	cipherData, err = rd2.Encrypt([]byte("hei hei hei "))
//	fmt.Println(cipherData)
//	plainData, err = rd2.Decrypt(cipherData)
//	fmt.Println(string(plainData))
//}
