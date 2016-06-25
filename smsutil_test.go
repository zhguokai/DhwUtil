package goutil

import "testing"

func TestSendYmMsgToMobile(t *testing.T) {
	var smsCdKey string = "****-***-****-*****"
	var smsPass string = "******"
	err := SendYmMsgToMobile("***********", "豆花网", "测试一下啦", smsCdKey, smsPass)
	if err != nil {
		t.Error(err.Error())
	}
}
