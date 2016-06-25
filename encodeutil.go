package dhwutil

import (
	"crypto/sha1"
	"io"
	"fmt"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"strings"
)


// EncodeSha1 对字符串进行哈希计算,产生Sha1的加密值
func EncodeSha1(data string) string {
	sh := sha1.New()
	io.WriteString(sh, data)
	return fmt.Sprintf("%x", sh.Sum(nil))
}

// EncodeMd5 对字符串进行MD5加密
func EncodeMd5(data string) string {
	m5 := md5.New()
	io.WriteString(m5, data)
	return fmt.Sprintf("%x", m5.Sum(nil))
}


// EncodeSha512 对字符串进行哈希计算,产生Sha1的加密值
func EncodeSha512(oldStr string) (newStr string) {
	hash := sha512.New()
	oldStrByte := []byte(oldStr)
	hash.Write(oldStrByte)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return strings.ToUpper(mdStr)
}