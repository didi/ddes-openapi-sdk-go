package core

import (
	"net/http"
	"net/url"
)

type ApiReq struct {
	HttpMethod  string            `json:"http_method"`  // http methods
	ApiPath     string            `json:"api_path"`     // 接口地址
	Body        interface{}       `json:"body"`         // 请求体
	QueryParams url.Values        `json:"query_params"` // 接口查询参数
	PathParams  map[string]string `json:"path_params"`  // 接口路径参数
}

//
//type QueryParams map[string][]string
//type PathParams map[string]string

type ApiResp struct {
	StatusCode int         `json:"-"` // 状态码
	Header     http.Header `json:"-"` // 响应头信息
	Body       []byte      `json:"-"` // 响应体
}
