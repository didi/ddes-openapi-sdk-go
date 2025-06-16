package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var tokenManager = LocalTokenManager{tokenCache: tokenCache} //默认的token管理器

type LocalTokenManager struct {
	tokenCache TokenCache
}

func NewLocalTokenManager(tokenCache TokenCache) *LocalTokenManager {
	tokenManager = LocalTokenManager{tokenCache: tokenCache}
	return &tokenManager
}

func (tokenManager *LocalTokenManager) SetToken(value string, ttl time.Duration) {
	tokenManager.tokenCache.SetToken(value, ttl)
}

func (tokenManager *LocalTokenManager) GetAccessToken(option *Option) (string, error) {
	if nil == option {
		return "", errors.New("option is nil")
	}
	if nil == tokenManager.tokenCache {
		return "", errors.New("token manager tokenCache is nil")
	}
	// 从本地缓存读取token数据
	token, err := tokenManager.tokenCache.GetToken()
	if err != nil {
		return "", err
	}
	if "" != token {
		return token, nil
	}
	// 从服务器获取token数据
	body := make(map[string]interface{})
	body["client_id"] = option.clientId
	body["client_secret"] = option.clientSecret
	body["grant_type"] = "client_credentials"
	body["timestamp"] = time.Now().Unix()
	apiReq := &ApiReq{
		HttpMethod:  http.MethodPost,
		ApiPath:     "/river/Auth/authorize",
		Body:        body,
		QueryParams: nil,
		PathParams:  nil,
	}
	apiResp, err := Request(context.Background(), apiReq, option, nil)
	if err != nil {
		return "", err
	}
	if http.StatusOK != apiResp.StatusCode {
		if nil == apiResp.Body {
			return "", fmt.Errorf("status code: %d", apiResp.StatusCode)
		}
		return "", fmt.Errorf("status code: %d,%s", apiResp.StatusCode, UnescapeUnicode(string(apiResp.Body)))
	}
	if nil == apiResp.Body {
		return "", fmt.Errorf("status code: %d", apiResp.StatusCode)
	}
	/**
	{
	  "errno" : 500,
	  "errmsg" : "access_token生成失败，请稍后重试",
	  "request_id" : "i+2SqRl3AckcUcP6Q6vAkR+1SZGqZoe0U16d3I5G7WcQIhI87V++1x5wN/FABZ80"
	}
	*/
	respBody := make(map[string]interface{})
	if err := json.Unmarshal(apiResp.Body, &respBody); err != nil {
		return "", err
	}
	accessTokenObj, hasToken := respBody["access_token"]
	expiresInObj, hasExpiresIn := respBody["expires_in"]
	if hasToken && hasExpiresIn {
		//if nil != apiResp.AuthorizeApiReply && "" != apiResp.AuthorizeApiReply.AccessToken && apiResp.AuthorizeApiReply.ExpiresIn > 0 {
		// 获取token成功
		accessToken, ok := accessTokenObj.(string)
		if !ok {
			return "", errors.New("access token is invalid")
		}
		expiresIn, eOk := expiresInObj.(float64)
		if !eOk {
			return "", errors.New("expires_in is invalid")
		}
		token = accessToken
		// 缓存新的token数据
		tokenManager.tokenCache.SetToken(token, time.Duration(expiresIn)*time.Second)
		return token, nil
	}

	return "", fmt.Errorf("status code: %d,%s", apiResp.StatusCode, UnescapeUnicode(string(apiResp.Body)))
}

// ApiReqParamsAssembly 接口请求参数组装，包括client_id,client_secret,access_token等
func ApiReqParamsAssembly(apiReq *ApiReq, option *Option) error {
	if nil == apiReq || "" == apiReq.ApiPath || nil == option {
		return nil
	}
	// 如果启用了token缓存则自动赋值access_token字段
	if option.EnableTokenCache {
		// 判断当前接口是否在白名单中
		if _, ok := noneAccessTokenParamPathMap[apiReq.ApiPath]; !ok {
			token, err := tokenManager.GetAccessToken(option)
			if err != nil {
				return err
			}
			if "" == token {
				return fmt.Errorf("get token failed,token is empty")
			}
			if nil != apiReq.QueryParams && len(apiReq.QueryParams) > 0 {
				apiReq.QueryParams.Set("access_token", token)
			}
			if nil != apiReq.Body {
				bodyMap := make(map[string]interface{})
				bytes, err := json.Marshal(apiReq.Body)
				if err != nil {
					return err
				}
				if err := json.Unmarshal(bytes, &bodyMap); err != nil {
					return err
				}
				// 检查 bodyMap 中是否包含 token 字段
				//if _, exists := bodyMap["access_token"]; !exists {
				//	// 若不包含 token 字段，则添加
				//	bodyMap["access_token"] = token
				//}
				// 启用了token缓存则直接默认用缓存token
				bodyMap["access_token"] = token
				apiReq.Body = bodyMap
			}
		}
	}

	// 自动赋值client_id,client_secret
	if nil != apiReq.QueryParams && len(apiReq.QueryParams) > 0 {
		if !apiReq.QueryParams.Has("client_id") {
			apiReq.QueryParams.Set("client_id", option.clientId)
		}
		if _, ok := noneAccessTokenParamPathMap[apiReq.ApiPath]; ok {
			if !apiReq.QueryParams.Has("client_secret") {
				apiReq.QueryParams.Set("client_secret", option.clientSecret)
			}
		}
	}
	if nil != apiReq.Body {
		bodyMap := make(map[string]interface{})
		bytes, err := json.Marshal(apiReq.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			return err
		}

		if _, exists := bodyMap["client_id"]; !exists {
			bodyMap["client_id"] = option.clientId
		}
		if _, ok := noneAccessTokenParamPathMap[apiReq.ApiPath]; ok {
			if _, exists := bodyMap["client_secret"]; !exists {
				bodyMap["client_secret"] = option.clientSecret
			}
		}
		apiReq.Body = bodyMap
	}
	return nil
}
