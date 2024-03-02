package jwt

import (
	"testing"
	"time"
)

func Test_Sign(t *testing.T) {
	token, err := Sign("test001")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Log(token)
	time.Sleep(2 * time.Second)
	data, e := Verify(token)
	if e != nil {
		t.Errorf("%v\n", e)
	}
	t.Log(data.Username)
}
