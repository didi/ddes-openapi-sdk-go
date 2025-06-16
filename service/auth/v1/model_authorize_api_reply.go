package v1

// AuthorizeApiReply struct for AuthorizeApiReply
type AuthorizeApiReply struct {
	AccessToken *string `json:"access_token,omitempty"` // 接口获取授权后的access token
	ExpiresIn   *int32  `json:"expires_in,omitempty"`   // access_token的生命周期，单位是秒数
	TokenType   *string `json:"token_type,omitempty"`   // access_token类型
	Scope       *string `json:"scope,omitempty"`        // 权限范围
	Errno       *int32  `json:"errno,omitempty"`        // 错误码
	Errmsg      *string `json:"errmsg,omitempty"`       // 错误文案
	RequestId   *string `json:"request_id,omitempty"`   // 请求ID(该字段一定要保留，方便排查问题)
}
