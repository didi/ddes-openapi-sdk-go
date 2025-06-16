package core

import (
	"testing"
	"time"
)

func Test_localCache_Set(t *testing.T) {
	cache := NewLocalTokenCache()
	cache.SetToken("name", 3*time.Second)
	value, err := cache.GetToken()
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
	time.Sleep(3 * time.Second)
	value, err = cache.GetToken()
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
}
