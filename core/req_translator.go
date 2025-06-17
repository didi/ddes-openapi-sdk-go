package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type RequestTranslator struct{}

func (translator *RequestTranslator) translate(ctx context.Context, apiReq *ApiReq, option *Option, reqOption *ReqOption) (*http.Request, error) {
	// 组装接口请求参数(access_token,client_id,client_secret等)
	if err := ApiReqParamsAssembly(apiReq, option); err != nil {
		return nil, err
	}
	// 给请求参数进行签名
	var bodyBuffer io.Reader
	var fullUrl string // 完整请求地址
	// 序列化请求体，原http.MethodPost, http.MethodPut, http.MethodPatch内逻辑
	if nil != apiReq.Body {
		rawBody, err := json.Marshal(apiReq.Body)
		if err != nil {
			return nil, err
		}
		//option.Logger.Debug(ctx, "签名前：", string(rawBody))
		var toSignDataMap map[string]interface{}
		if err := json.Unmarshal(rawBody, &toSignDataMap); err != nil {
			return nil, err
		}
		if _, ok := toSignDataMap["timestamp"]; !ok {
			var timestamp = time.Now().Unix()
			toSignDataMap["timestamp"] = timestamp
		}
		// 添加body参数的加密逻辑
		if option.EnableEncryption && nil != option.EncryptionOption {
			toSignDataMap, err = translator.AssemblyBodyMapToEncryptionDataMap(apiReq, toSignDataMap, option)
			if err != nil {
				return nil, err
			}
		}
		newMap := signer.SignPostBody(option.signKey, toSignDataMap)
		marshal, err := json.Marshal(newMap)
		if err != nil {
			return nil, err
		}
		//option.Logger.Debug(ctx, "签名后：", string(marshal))
		bodyBuffer = bytes.NewBuffer(marshal)
	}
	// 对查询参数进行签名，原http.MethodGet, http.MethodDelete内逻辑
	if len(apiReq.QueryParams) > 0 {
		if !apiReq.QueryParams.Has("timestamp") {
			apiReq.QueryParams.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
		}
		if option.EnableEncryption && nil != option.EncryptionOption {
			toEncQueryParams, err := translator.AssemblyQueryParamsToEncryptionUrlValue(apiReq, apiReq.QueryParams, option)
			if err != nil {
				return nil, err
			}
			apiReq.QueryParams = toEncQueryParams
		}
		//log.Println("转换加密后：", apiReq.QueryParams.Encode())

		signeValue, err := signer.SignQueryParams(apiReq.QueryParams, option.signKey)
		if err != nil {
			return nil, err
		}

		apiReq.QueryParams.Set("sign", signeValue)
		if apiReq.QueryParams.Has("sign_key") {
			apiReq.QueryParams.Del("sign_key")
		}
	}
	//log.Println("签名后：", apiReq.QueryParams.Encode())
	// 拼接完整地址
	fullUrl = AssemblyFullUrl(option.BaseUrl, apiReq.ApiPath, apiReq.PathParams, apiReq.QueryParams)

	request, err := http.NewRequestWithContext(ctx, apiReq.HttpMethod, fullUrl, bodyBuffer)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	//option.Logger.Debug(ctx, request.Method, fullUrl)
	// 添加额外的header参数
	if nil != reqOption && len(reqOption.Header) > 0 {
		for key, values := range reqOption.Header {
			for _, value := range values {
				request.Header.Add(key, value)
			}
		}
	}
	// 往header中添加用于统计的字段；
	didiAgent := "didi-open-sdk-go" + "/" + version
	request.Header.Add("User-Agent", didiAgent)

	return request, nil
}

var (
	// 无需进行加密的字段
	notEncryptionKeys = map[string]struct{}{
		"client_id":    {},
		"access_token": {},
		//"client_secret": {},
		//"timestamp": {},
	}
	// 无需加密传输的接口
	notEncryptionApis = map[string]struct{}{
		"/river/Auth/authorize": {},
	}
)

const (
	EntKey        = "ent"             // 传输类型：0-普通传输，1-AES128加密传输，2-AES256加密传输
	EntContentKey = "encrypt_content" // 密文，入参加密后的密文在ent为1或2时必传
)

// AssemblyBodyMapToEncryptionDataMap 将待签名的Body Map数据转换为待签名的加密传输map数据
func (translator *RequestTranslator) AssemblyBodyMapToEncryptionDataMap(apiReq *ApiReq, dataMap map[string]interface{}, option *Option) (map[string]interface{}, error) {
	// 过滤不支持加密传输的接口
	if _, ok := notEncryptionApis[apiReq.ApiPath]; ok {
		return dataMap, nil
	}
	toEncMap := make(map[string]interface{})
	toEncMap[EntKey] = option.EncryptionOption.Ent
	// 获取待加密字符串
	toEncParamsArr := make([]string, 0) // 待加密的字段列表
	for key, value := range dataMap {
		if _, ok := notEncryptionKeys[key]; ok {
			toEncMap[key] = value
		} else {
			// 拼接待加密字符串
			var valueStr string
			switch v := value.(type) {
			case string:
				valueStr = v
			case int:
				valueStr = strconv.Itoa(v)
			case int64:
				valueStr = strconv.FormatInt(v, 10)
			case float64:
				valueStr = strconv.FormatFloat(v, 'f', -1, 64)
			default:
				valueStr = fmt.Sprintf("%v", v)
			}
			toEncParamsArr = append(toEncParamsArr, fmt.Sprintf("%s=%s", key, fmt.Sprintf("%v", valueStr)))
		}
	}
	toEnvStr := strings.Join(toEncParamsArr, "&")
	//log.Printf("toEncryption value: %s", toEnvStr)
	//
	// 网站提供的加密逻辑
	//aesTool := NewAESTool(option.EncryptionOption.Key)
	//encryptECB, err := aesTool.ECBEncrypt128(toEnvStr)
	//if err != nil {
	//	return nil, err
	//}
	//toEncMap[EntContentKey] = encryptECB

	encryptECB, err := AESEncryptECB([]byte(toEnvStr), []byte(option.EncryptionOption.Key))
	if err != nil {
		return nil, err
	}
	toEncMap[EntContentKey] = base64.StdEncoding.EncodeToString(encryptECB)

	return toEncMap, nil
}

// AssemblyQueryParamsToEncryptionUrlValue 将查询参数转换为待签名的加密传输格式的查询参数
func (translator *RequestTranslator) AssemblyQueryParamsToEncryptionUrlValue(apiReq *ApiReq, queryParams url.Values, option *Option) (url.Values, error) {
	// 过滤不支持加密传输的接口
	if _, ok := notEncryptionApis[apiReq.ApiPath]; ok {
		return queryParams, nil
	}
	toEncValues := url.Values{}
	toEncValues.Set(EntKey, strconv.FormatInt(int64(option.EncryptionOption.Ent), 10))
	// 组装待加密字符串
	var paramStr strings.Builder

	for key, values := range queryParams {
		if _, ok := notEncryptionKeys[key]; ok {
			toEncValues[key] = values
		} else {
			if len(values) > 1 {
				// 说明：当查询参数一个key有多个值的时候，目前后端是取最后一个值进行签名计算；
				value := values[len(values)-1]
				if paramStr.Len() > 0 {
					paramStr.WriteByte('&')
				}
				paramStr.WriteString(key)
				paramStr.WriteByte('=')
				paramStr.WriteString(value)
			} else {
				for _, v := range values {
					if paramStr.Len() > 0 {
						paramStr.WriteByte('&')
					}
					paramStr.WriteString(key)
					paramStr.WriteByte('=')
					paramStr.WriteString(v)
				}
			}
		}
	}
	toEncStr := paramStr.String()
	//log.Printf("待加密数据: %s", toEncStr)

	// 网站提供的加密逻辑
	//aesTool := NewAESTool(option.EncryptionOption.Key)
	//encryptECB, err := aesTool.ECBEncrypt128(toEncStr)
	//if err != nil {
	//	return nil, err
	//}
	//toEncValues.Set(EntContentKey, encryptECB)

	encryptECB, err := AESEncryptECB([]byte(toEncStr), []byte(option.EncryptionOption.Key))
	if err != nil {
		return nil, err
	}

	encryptECBBase64 := base64.StdEncoding.EncodeToString(encryptECB)
	//queryEscape := url.QueryEscape(encryptECBBase64)
	toEncValues.Set(EntContentKey, encryptECBBase64)
	return toEncValues, nil
}

const (
	ParamAccessTokenKeyName = "access_token"
)

//// 需要添加access_token的接口中如果用户未主动赋值则默认添加；
//func (translator *RequestTranslator) addAccessTokenParamToBody(m map[string]interface{}, apiReq *ApiReq, option *Option) {
//	if "" == apiReq.ApiPath || nil == option.TokenCache {
//		return
//	}
//	// 判断接口是否需要添加access_token
//	if _, ok := noneAccessTokenParamPathMap[apiReq.ApiPath]; ok {
//		return
//	}
//	// 判断用户是否传递了token参数
//	if value, ok := m[ParamAccessTokenKeyName]; ok {
//		valueStr, ok := value.(string)
//		if ok && valueStr != "" {
//			return
//		}
//	}
//	// 从本地缓存中获取access_token，如果不存在则发起token认证请求获取；
//	cacheToken := option.TokenCache.GetToken()
//	if "" == cacheToken {
//		translator.SendAuthReq(option)
//	}
//	m[ParamAccessTokenKeyName] = cacheToken
//}
//
//func (translator *RequestTranslator) SendAuthReq(option *Option) {
//
//}
