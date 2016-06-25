package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//func:GetRequest
//desc:发送GET请求
//param:url string  仅支持http请求 例:http://www.baidu.com
//return:[]byte 返回值:可以转为String,JSON,XML
//return:error 请求错误
func GetRequest(url string) ([]byte, error) {
	oid, err := http.Get(url)
	defer oid.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	oids, err := ioutil.ReadAll(oid.Body)
	return oids, err
}

// JsonPostRequest用来向指定的URL发送POST请求,
// 参数格式为JSON格式,
// 参数值为map[string]interface{}
// 执行完成之后,返回结果集为[]byte数组和error信息。
// []byte可以根据实际情况转换成string,JSOn,struct
// errors不为空代表,请求过程失败
// BUG(author:)
func JsonPostRequest(url string, data map[string]interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json err:", err.Error())
		return []byte{}, err
	}

	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}
	return result, nil
}

// XmlBytePostRequest用来向指定的URL发送POST请求,
// 参数格式为XML格式,
// 参数值为XML格式的字符串
// 执行完成之后,返回结果集为[]byte数组和error信息。
// []byte可以根据实际情况转换成string,JSOn,struct
// errors不为空代表,请求过程失败
// BUG(author:)
func XmlStringPostRequest(url string, xmlStringData string) ([]byte, error) {

	body := bytes.NewBuffer([]byte(xmlStringData))
	res, err := http.Post(url, "text/xml;charset=utf-8", body)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

// XmlBytePostRequest用来向指定的URL发送POST请求,
// 参数格式为XML格式,
// 参数值为struct进行xml.Marshal()后的字节数组
// 执行完成之后,返回结果集为[]byte数组和error信息。
// []byte可以根据实际情况转换成string,JSOn,struct
// errors不为空代表,请求过程失败
// BUG(author:)
func XmlBytePostRequest(url string, xmlData []byte) ([]byte, error) {
	body := bytes.NewBuffer(xmlData)
	res, err := http.Post(url, "text/xml;charset=utf-8", body)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}
