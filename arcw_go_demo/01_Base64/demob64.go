package main

import (
	"encoding/base64"
	"fmt"
)

/**
Base64编解码
*/
func main() {
	// 待编码字符串
	data := "abc123!?$*&()'-=@-"

	//  1. 标准的base64编码
	// 编码字符串
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码字符串
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	fmt.Println()

	// 2. 兼容URL的base64编码
	// 编码字符串
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码字符串
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
