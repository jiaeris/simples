package buf

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func SimpleDemo() {

	buf := IntToBytes(121014)
	fmt.Println(len(buf))

	Int := BytesToInt(buf)
	fmt.Println(Int)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	tmp := int32(n) //4个字节8位
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}
