package myjwt

//jwt demo
//Header.Payload.VerifySignature
//Header头部：包含了，类型和加密方式"Header":{"alg":"HS256","typ":"JWT"}
//Payload包含了自定义信息
//第三部分包含加密信息
func HSTest() {
	useSigningMethodHS256()
}

func RSTest() {
	useSigningMethodRS256()
}
