package goutil

import (
	"errors"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"log"
)


//发送自定义短信
func SendYmMsgToMobile(mobile, signName, msg, smsCdKey, smsPass string) error {
	smsGateWayUrl := "http://sdk4rptws.eucp.b2m.cn:8080/sdkproxy/sendsms.action"
	if "" == mobile {
		return errors.New("手机号不能为空！")
	}
	v := url.Values{}
	v.Set("cdkey", smsCdKey)
	v.Set("password", smsPass)
	v.Set("phone", mobile)
	v.Set("message", "【" + signName + "】" + msg)
	if resp, err := http.PostForm(smsGateWayUrl, v); err == nil {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//xml.Unmarshal(data,&)
		var resp YmSmsResp
		err := xml.Unmarshal(data, &resp)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		if resp.Error == "0" {
			return nil
		}
		return errors.New(resp.Message)
	} else {
		defer resp.Body.Close()
		return err
	}
}

type YmSmsResp struct {
	XMLName xml.Name `xml:"response"`
	Error   string   `xml:"error"`
	Message string   `xml:"message"`
}