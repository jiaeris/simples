package myjwt

import (
	"jwt-go"
	"fmt"
	"crypto/rsa"
	//"time"
	"encoding/json"
	"crypto/rand"
)

var prk *rsa.PrivateKey
var puk *rsa.PublicKey

func useSigningMethodRS256() {
	//创建公私钥对象
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	var err error
	prk, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
	}
	puk = &prk.PublicKey
	fmt.Println("prk: " + prk.D.String())
	fmt.Println("puk: " + puk.N.String())
	fmt.Println(len(prk.D.String()))
	fmt.Println(len(puk.N.String()))

	token := jwt.New(jwt.SigningMethodRS256)
	//头部为一个Map，可增加头部信息
	//token.Header["nice"] = "ok"
	token.Claims = jwt.MapClaims{
		"name": "Yunga",
		"age":  25,
	}

	tokenString := rs256encode(token)
	rs256decode(tokenString)
}

func rs256encode(token *jwt.Token) string {
	tokenString, err := token.SignedString(prk)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
	return tokenString
}

func rs256decode(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return puk, nil
	})
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(token)
	fmt.Println(string(b))
	fmt.Println(token.Valid)
}

//读取
//func readKey(keyPath string) ([]byte, error) {
//	file, err := os.Open(keyPath)
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	buf, err := ioutil.ReadAll(file)
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	return buf, nil
//}
