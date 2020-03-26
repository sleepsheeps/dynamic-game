package utils

import "testing"

func TestGetString(t *testing.T) {
	InitCache(":6379")
	res, err := GetString("test")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(res)
}

func TestSetString(t *testing.T) {
	InitCache(":6379")
	err := SetString("test", "test")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestDel(t *testing.T) {
	InitCache(":6379")
	err := Del("test")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
