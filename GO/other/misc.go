package other

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Atoi string转int
func Atoi(str string) int{
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
//Itoa int转string
func Itoa(num int) string{
	return strconv.Itoa(num)
}

//I64ToA int64转string
func I64ToA(num int64)string{
	return strconv.FormatInt(num,10)
}

//AtoInt64 string转int64
func AtoInt64(str string) int64{
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

//Ifs 伪三目运算符
func Ifs(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// 识别手机号码
func IsMobile(mobile string) bool{
	result, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if result {
		return true
	} else {
		return false
	}
}

//生成x位随机验证码
func GenValidateCode(width int) string {
	numeric := []byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}


//RandChar 生成随机长度字符串
func RandChar(size int) string {
	var char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i ++ {
		s.WriteByte(char[rand.Int63() % int64(len(char))])
	}
	return s.String()
}
