package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5(str string) (encode string) {
	md5 := md5.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}

func main() {
	encode := Md5("coopers")
	fmt.Println("golang用md5加密32位：", encode) //22c1158d867b1f855c8105127f4a284a
}

//ref: https://blog.csdn.net/HelloWorldYangSong/article/details/99703499
