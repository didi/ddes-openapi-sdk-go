package core

import (
	"sync"
	"time"
)

const (
	tokenKey = "access_token_cache"
)

var tokenCache = &localTokenCache{} // 本地的token缓存

func NewTokenCache(option *Option) {
	if nil != option.TokenCache {
		tokenManager = LocalTokenManager{
			tokenCache: option.TokenCache,
		}
	}
}

type TokenCache interface {
	SetToken(value string, ttl time.Duration)
	GetToken() (string, error)
}

// 本地Token缓存
type localTokenCache struct {
	m sync.Map
}

type CacheValue struct {
	Value      string
	ExpireTime time.Time
}

func NewLocalTokenCache() TokenCache {
	return &localTokenCache{}
}

func (cache *localTokenCache) SetToken(value string, ttl time.Duration) {
	// 根据ttl计算超时时间
	expireTime := time.Now().Add(ttl)
	cache.m.Store(tokenKey, &CacheValue{
		Value:      value,
		ExpireTime: expireTime,
	})
}
func (cache *localTokenCache) GetToken() (string, error) {
	// 根据key查询缓存数据
	value, ok := cache.m.Load(tokenKey)
	if !ok {
		return "", nil
	}
	cacheValue := value.(*CacheValue)
	// 判断是否过期,如果已经过期则返回空
	if cacheValue.ExpireTime.Before(time.Now()) {
		cache.m.Delete(tokenKey) // 删除已过期的数据
		return "", nil
	}
	return cacheValue.Value, nil
}
