package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	signer = Signer{}
	// 请求转换
	requestTranslator = &RequestTranslator{}
	//无需组装access_token参数的接口，除此之外，如果用户设置了token缓存接口，需要自动将token写入请求数据中；
	noneAccessTokenParamPathMap = map[string]struct{}{
		"/river/Auth/authorize": {},
	}
)

// Request 发起http请求。1、将apiReq转换为http.Request; 2、http client发送请求；
func Request(ctx context.Context, apiReq *ApiReq, option *Option, reqOption *ReqOption) (*ApiResp, error) {
	// 将apiReq转换为http request
	request, err := requestTranslator.translate(ctx, apiReq, option, reqOption)
	if err != nil {
		return nil, err
	}
	option.Logger.Debug(ctx, request.Method, request.URL.String())
	if nil != request.Body {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, request.Body); err != nil {
			return nil, err
		}
		request.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		option.Logger.Debug(ctx, "Request Body:", string(buf.Bytes()))
	}

	client := &http.Client{}
	if option.RequestTimeOut > 0 { // 获取全局超时设置
		client.Timeout = option.RequestTimeOut
	}
	if nil != reqOption && reqOption.RequestTimeOut > 0 { // 获取当前请求超时设置
		client.Timeout = option.RequestTimeOut
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		if nil != body {
			_ = body.Close()
		}
	}(response.Body)
	// 组装响应参数
	apiResp := ApiResp{
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       nil,
	}
	if nil != response.Body {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		apiResp.Body = body
		option.Logger.Debug(ctx, "Response Body:", UnescapeUnicode(string(apiResp.Body)))
	}

	return &apiResp, nil
}

// AssemblyFullUrl 根据baseUrl、接口地址、路径参数、查询参数等组装完整请求地址
func AssemblyFullUrl(baseUrl string, apiPath string, pathParams map[string]string, queryParams map[string][]string) string {
	// 路径参数
	for key, value := range pathParams {
		placeholder := fmt.Sprintf("{%s}", key)
		apiPath = strings.ReplaceAll(apiPath, placeholder, url.PathEscape(value))
	}
	// 拼接baseUrl和apiPath
	fullUrl := strings.TrimRight(baseUrl, "/") + apiPath
	// 查询参数
	if len(queryParams) > 0 {
		query := url.Values{}
		for key, values := range queryParams {
			for _, value := range values {
				query.Add(key, value)
			}
		}
		if len(query.Encode()) > 0 {
			fullUrl += "?" + query.Encode()
		}
	}
	return fullUrl
}
