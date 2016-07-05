package goutil

import (
	"testing"
	"log"
)

func TestStrToUnixTime(t *testing.T) {
	ts := "2016-06-08 12:11:23"
	tu, err := StrToUnixTime("2006-01-02 15:04:05",ts)
	if err != nil {
		t.Error(err.Error())
	}
	println(tu)
}

func TestGetFirstDayTimeInMonth(t *testing.T) {
	log.Println(GetFormatTime(YYYYMMDDHIMMSSsss))
}