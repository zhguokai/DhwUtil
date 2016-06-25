package goutil

import (
	"time"
	"strings"
	"math/rand"
	"math/big"
	cr "crypto/rand"
)
//生成随机数字
func GetRandomNum(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return strings.ToUpper(string(result))
}
//生成两个数之间的随机数
func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := cr.Int(cr.Reader, maxBigInt)

	for i.Int64() < min {
		i, _ = cr.Int(cr.Reader, maxBigInt)
	}
	return i.Int64()
}