package core

import (
	"net/http"
	"net/url"
	"time"
)

type Option struct {
	clientId         string            // client_id
	clientSecret     string            // client_secret
	signKey          string            // 签名KEY
	BaseUrl          string            // 服务请求地址
	RequestTimeOut   time.Duration     // HTTP请求超时设置
	EnableTokenCache bool              // 是否启用Token缓存，true：用户无需主动传入access_token，sdk自行处理；false：sdk不缓存token，由用户自行传入；
	TokenCache       TokenCache        // TOKEN缓存自定义实现，不传则使用默认缓存
	EnableEncryption bool              // 是否启用加密传输
	EncryptionOption *EncryptionOption // 加密相关参数
	SignMethod       int               // 签名方式 1:MD5 2:SHA256 (默认MD5)
	Serializer       Serializer        // 序列化
	LogLevel         LogLevel
	Logger           Logger
}
type EncryptionOption struct {
	Ent int    // 传输类型：0 普通传输，1 AES128位加密传输， 2 AES256位加密传输
	Key string // 密钥
}

func (option *Option) GetClientId() string {
	return option.clientId
}
func (option *Option) GetClientSecret() string {
	return option.clientSecret
}
func (option *Option) GetSignKey() string {
	return option.signKey
}

type Encryption struct {
	Type   int    // 加密方式
	PubKey string // 公钥
}

const (
	baseUrl = "https://api.es.xiaojukeji.com"
)

// NewDefaultOption 初始化默认配置
func NewDefaultOption(appId, appSecret, signKey string) *Option {
	return &Option{
		clientId:         appId,
		clientSecret:     appSecret,
		signKey:          signKey,
		BaseUrl:          baseUrl,
		RequestTimeOut:   10 * time.Second,
		EnableTokenCache: false,
		TokenCache:       NewLocalTokenCache(),
		EnableEncryption: false,
		EncryptionOption: nil,
		SignMethod:       1,
		Serializer:       &DefaultSerializer{},
		LogLevel:         LogLevelInfo,
		Logger:           NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}
}

type ReqOption struct {
	Header         http.Header       // 请求头信息
	RequestTimeOut time.Duration     // HTTP请求超时设置(仅对当前请求有效，会覆盖全局超时设置)
	QueryParams    url.Values        // 接口查询参数
	PathParams     map[string]string // 接口路径参数
	Serializer     Serializer        // 序列化,
}
