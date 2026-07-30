package core

import (
	"strconv"
)

// RawMessageToString 将 JSON 原始字节（json.RawMessage）统一转为字符串。
//
// 用于 reply struct 的自定义 UnmarshalJSON：真实流量中部分 ID/枚举字段类型
// 不稳定（number/string/空串/非数字混存），标准 json.Unmarshal 对 *string 收
// number 会失败。用 json.RawMessage 先接收原始字节，再经本函数统一转 string：
//   - number → 字符串形式（123 → "123"，1.5 → "1.5"）
//   - string → 去引号后的原值
//   - null/空 → ""（调用方按需判断是否赋值）
//   - bool → "true"/"false"
//   - 其他（array/object）→ 原始 JSON 文本
//
// 零外部依赖。
func RawMessageToString(data []byte) (string, error) {
	if len(data) == 0 || string(data) == "null" {
		return "", nil
	}
	// string：去引号
	if data[0] == '"' {
		s, err := strconv.Unquote(string(data))
		if err != nil {
			return string(data[1 : len(data)-1]), nil
		}
		return s, nil
	}
	// bool
	if string(data) == "true" || string(data) == "false" {
		return string(data), nil
	}
	// number：int 优先，再 float
	if i, err := strconv.ParseInt(string(data), 10, 64); err == nil {
		return strconv.FormatInt(i, 10), nil
	}
	if fl, err := strconv.ParseFloat(string(data), 64); err == nil {
		return strconv.FormatFloat(fl, 'f', -1, 64), nil
	}
	// 兜底：原始文本
	return string(data), nil
}
