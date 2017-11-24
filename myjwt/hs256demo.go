package myjwt

import (
	"jwt-go"
	"fmt"
	"encoding/json"
)

const (
	secretSigningString = "TestString"
)

//Hash Message Authentication Code
//使用哈希方式进行加密运算
//jwt提供了256,384,512长度算法
func useSigningMethodHS256() {
	token := jwt.New(jwt.SigningMethodHS256)
	//头部为一个Map，可增加头部信息
	//token.Header["nice"] = "ok"
	token.Claims = jwt.MapClaims{
		"name": "Yunga",
		"age":  25,
	}

	tokenString := hs256encode(token)
	hs256decode(tokenString)
}

//编码
func hs256encode(token *jwt.Token) string {
	//获取完整的令牌
	tokenString, err := token.SignedString([]byte(secretSigningString))
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)

	//获取令牌中最重要的部分，头和payload
	//tokenString, err = token.SigningString()
	//fmt.Println(tokenString)

	return tokenString
}

//解码
func hs256decode(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		//首先进行jwt前两部的验证
		claims := token.Claims.(jwt.MapClaims)
		name := claims["name"]
		age := claims["age"]
		fmt.Println(name)
		fmt.Println(age)

		//最后返回key，在parse函数内验证第三部分·签名·
		return []byte(secretSigningString), nil
	})
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(token)
	fmt.Println(string(b))
	//fmt.Println(token.Valid)
}
