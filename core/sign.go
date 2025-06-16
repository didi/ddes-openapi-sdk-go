package core

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type Signer struct{}

// SignQueryParams 对查询参数进行签名
func (signer *Signer) SignQueryParams(params url.Values, signKey string) (string, error) {
	// 忽略sign字段
	if params.Has("sign") {
		params.Del("sign")
	}
	// 添加signKey
	if !params.Has("sign_key") {
		params.Set("sign_key", signKey)
	}
	// TODO 处理__obj__字段，实现对象自动转json字符串
	//for key, value := range params {
	//	if !strings.HasSuffix(key, "__obj__") {
	//		continue
	//	}
	//	realKey := strings.TrimSuffix(key, "__obj__")
	//	if _, ok := params[realKey]; !ok {
	//		for _, v := range value {
	//			bytes, err := json.Marshal(&v)
	//			if err != nil {
	//				return "", err
	//			}
	//			params.Add(realKey, string(bytes))
	//		}
	//	}
	//	delete(params, key)
	//}

	// 获取key并对key进行排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接参数键值对
	var paramStr strings.Builder
	for _, key := range keys {
		values := params[key]
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
	toSignStr := paramStr.String()

	//log.Printf("toSignStr: %s", toSignStr)
	h := md5.New()
	h.Write([]byte(toSignStr))
	encodeToString := hex.EncodeToString(h.Sum(nil))
	return encodeToString, nil
	//bytes, err := json.Marshal(params)
	//if err != nil {
	//	return "", nil
	//}
	//var m map[string]interface{}
	//if err := json.Unmarshal(bytes, &m); err != nil {
	//	return "", err
	//}
	////m["sign_key"] = signKey
	//if _, ok := m["timestamp"]; !ok {
	//	m["timestamp"] = time.Now().Unix()
	//}
	//var keys []string
	//// 提取所有的键
	//for key := range m {
	//	keys = append(keys, key)
	//}
	//// 对键进行排序
	//sort.Strings(keys)
	//
	//var parts []string
	//// 遍历排序后的键，将键值对按 key=Value 格式组装+
	//for _, key := range keys {
	//	// 忽律sign字段本身
	//	if key == "sign" {
	//		continue
	//	}
	//	if Value, ok := m[key]; ok {
	//		var valueStr string
	//		switch v := Value.(type) {
	//		case string:
	//			valueStr = v
	//		case int:
	//			valueStr = strconv.Itoa(v)
	//		case int64:
	//			valueStr = strconv.FormatInt(v, 10)
	//		case float64:
	//			valueStr = strconv.FormatFloat(v, 'f', -1, 64)
	//		default:
	//			valueStr = fmt.Sprintf("%v", v)
	//		}
	//		parts = append(parts, fmt.Sprintf("%s=%s", key, fmt.Sprintf("%v", valueStr)))
	//	}
	//}
	//// 使用 & 连接所有的键值对
	//toSignStr := strings.Join(parts, "&")
}

// SignPostBody 对post请求体进行签名
func (signer *Signer) SignPostBody(signKey string, m map[string]interface{}) map[string]interface{} {
	m["sign_key"] = signKey

	// 处理__obj__字段,实现对象自动替换
	for key, value := range m {
		if !strings.HasSuffix(key, "__obj__") {
			continue
		}
		realKey := strings.TrimSuffix(key, "__obj__")
		if _, ok := m[realKey]; !ok {
			bytes, err := json.Marshal(&value)
			if err != nil {
				return nil
			}
			m[realKey] = string(bytes)
		}
		delete(m, key)
	}

	var keys []string
	// 提取所有的键
	for key := range m {
		keys = append(keys, key)
	}
	// 对键进行排序
	sort.Strings(keys)

	var parts []string
	// 遍历排序后的键，将键值对按 key=Value 格式组装+
	for _, key := range keys {
		// 忽律sign字段本身
		if key == "sign" {
			continue
		}
		if value, ok := m[key]; ok {
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
			case []interface{}, interface{}:
				xvl, err := json.Marshal(&v)
				if err != nil {
					return nil
				}
				valueStr = string(xvl)
			default:
				valueStr = fmt.Sprintf("%v", v)
			}
			parts = append(parts, fmt.Sprintf("%s=%s", key, fmt.Sprintf("%v", valueStr)))
		}
	}
	// 使用 & 连接所有的键值对
	toSignStr := strings.Join(parts, "&")
	//log.Printf("toSignStr: %s", toSignStr)
	h := md5.New()
	h.Write([]byte(toSignStr))
	encodeToString := hex.EncodeToString(h.Sum(nil))
	m["sign"] = encodeToString
	delete(m, "sign_key")
	return m
}

func SignRequest(signKey string, body interface{}) (string, error) {
	// 签名
	bytes, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(bytes, &m); err != nil {
		panic(err)
	}
	m["sign_key"] = signKey
	var keys []string
	// 提取所有的键
	for key := range m {
		keys = append(keys, key)
	}
	// 对键进行排序
	sort.Strings(keys)

	var parts []string
	// 遍历排序后的键，将键值对按 key=Value 格式组装+
	for _, key := range keys {
		// 忽律sign字段本身
		if key == "sign" {
			continue
		}
		if value, ok := m[key]; ok {
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
			parts = append(parts, fmt.Sprintf("%s=%s", key, fmt.Sprintf("%v", valueStr)))
		}
	}
	// 使用 & 连接所有的键值对
	toSignStr := strings.Join(parts, "&")
	//log.Printf("toSignStr: %s", toSignStr)
	h := md5.New()
	h.Write([]byte(toSignStr))
	encodeToString := hex.EncodeToString(h.Sum(nil))
	return encodeToString, nil
}
