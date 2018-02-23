package rsadriver

import (
	"os"
	"errors"
	"io/ioutil"
	"crypto/rsa"
	"encoding/pem"
	"crypto/x509"
	"crypto/rand"
)

type Type int

//default PKCS1
const (
	PKCS1 = iota
	PKCS8
)

type RSADriver interface {
	InitByFile(priKeyPath, pubKeyPath string, pass ...string) *RSA
	InitByByte(priKeyBytes, pubKeyBytes []byte, pass string) *RSA
	SetPKCSVersion(t Type)
	Encrypt(plainData []byte) ([]byte, error)
	Decrypt(cipherData []byte) ([]byte, error)
}

type RSA struct {
	Type   Type
	PriKey *rsa.PrivateKey
	PubKey *rsa.PublicKey
	isInit bool
}

//use this method to create driver object.
func NewRSA() RSADriver {
	rsa := &RSA{
		Type:   PKCS1,
		isInit: false,
	}
	return rsa
}

//
func (r *RSA) SetPKCSVersion(t Type) {
	r.Type = t
}

func (r *RSA) InitByFile(priKeyPath, pubKeyPath string, pass ...string) *RSA {
	priKeyFile, err := os.Open(priKeyPath)
	if err != nil {
		panic(errors.New("open private key file error"))
	}
	priKeyBytes, err := ioutil.ReadAll(priKeyFile)
	if err != nil {
		panic(errors.New("read private key file error"))
	}
	pubKeyFile, err := os.Open(pubKeyPath)
	if err != nil {
		panic(errors.New("open public key file error"))
	}
	pubKeyBytes, err := ioutil.ReadAll(pubKeyFile)
	if err != nil {
		panic(errors.New("read public key file error"))
	}
	if len(pass) > 0 {
		return r.InitByByte(priKeyBytes, pubKeyBytes, pass[0])
	} else {
		return r.InitByByte(priKeyBytes, pubKeyBytes, "")
	}
}

func (r *RSA) InitByByte(priKeyBytes, pubKeyBytes []byte, pass string) *RSA {
	//decode public key
	pubBlock, _ := pem.Decode(pubKeyBytes)
	if pubBlock == nil {
		panic(errors.New("public key error"))
	}
	pubKeyI, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic(err)
	}
	r.PubKey = pubKeyI.(*rsa.PublicKey)

	//decode private key
	priBlock, _ := pem.Decode(priKeyBytes)
	if priBlock == nil {
		panic(errors.New("private key error"))
	}
	var priBlockBytes []byte
	if len(pass) > 0 {
		var err error
		priBlockBytes, err = x509.DecryptPEMBlock(priBlock, []byte(pass))
		if err != nil {
			panic(errors.New("verify password error,please check your password"))
		}
	} else {
		priBlockBytes = priBlock.Bytes
	}
	var priKey *rsa.PrivateKey
	var pppkerr error
	switch r.Type {
	case PKCS1:
		priKey, pppkerr = x509.ParsePKCS1PrivateKey(priBlockBytes)
		break
	case PKCS8:
		priKeyI, pppkerrI := x509.ParsePKCS8PrivateKey(priBlockBytes)
		priKey = priKeyI.(*rsa.PrivateKey)
		pppkerr = pppkerrI
		break
	default:
		panic(errors.New("have no this type, types(pkcs1,pkcs8)"))
	}
	if pppkerr != nil {
		panic(pppkerr)
	}
	r.PriKey = priKey
	r.isInit = true
	return r
}

func (r *RSA) Encrypt(plainData []byte) ([]byte, error) {
	if !r.isInit {
		return nil, errors.New("please init rsa driver")
	}
	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, r.PubKey, plainData)
	if err != nil {
		return nil, err
	}
	return cipherData, nil
}

func (r *RSA) Decrypt(cipherData []byte) ([]byte, error) {
	if !r.isInit {
		return nil, errors.New("please init rsa driver")
	}
	plainData, err := rsa.DecryptPKCS1v15(rand.Reader, r.PriKey, cipherData)
	if err != nil {
		return nil, err
	}
	return plainData, nil
}
