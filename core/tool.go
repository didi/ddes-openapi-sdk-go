package core

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// 错误类型
var (
	ErrUnsupportedType = fmt.Errorf("unsupported type")
	ErrInvalidValue    = fmt.Errorf("invalid value")
)

type SmartDecoder struct {
	// 配置选项
	//StrictMode bool // 严格模式：遇到无法转换的值时返回错误
}

// NewSmartDecoder 创建新的智能解码器
func NewSmartDecoder() *SmartDecoder {
	return &SmartDecoder{}
}

// Decode 解析JSON数据到结构体
func (d *SmartDecoder) Decode(data []byte, v interface{}) error {
	// 先尝试标准JSON解析
	if err := json.Unmarshal(data, v); err == nil {
		return nil
	}
	// 如果标准解析失败，使用中间映射进行转换
	var intermediate map[string]interface{}
	if err := json.Unmarshal(data, &intermediate); err != nil {
		return err
	}
	// 将中间映射转换到目标结构体
	return d.mapToStruct(intermediate, reflect.ValueOf(v).Elem())
}
func LowercaseFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func (d *SmartDecoder) mapToStruct(m map[string]interface{}, v reflect.Value) error {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}
		if jsonTag == "" {
			jsonTag = LowercaseFirst(fieldType.Name)
		}
		tagName := strings.Split(jsonTag, ",")[0]
		value, exists := m[tagName]
		if !exists {
			continue
		}
		fieldValue := v.Field(i)
		if !fieldValue.CanSet() {
			continue
		}
		if err := d.setValue(fieldValue, value); err != nil {
			// 均按严格模式执行
			return fmt.Errorf("field %s: %w", tagName, err)
			//if d.StrictMode {
			//	return fmt.Errorf("field %s: %w", tagName, err)
			//}
			//// 在非严格模式下忽略错误
			//continue
		}
	}
	return nil
}

type s struct {
	boolValue     bool  `json:"bool_value"`  // false
	bytesValuePrt *bool `json:"bytes_value"` // nil
}

// setValue 设置字段值，处理类型转换
func (d *SmartDecoder) setValue(field reflect.Value, value interface{}) error {
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		return d.setValue(field.Elem(), value)
	}
	switch field.Kind() {
	case reflect.Bool: // 【checked】
		return d.setBool(field, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64: // 【checked】
		return d.setInt(field, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64: //【checked】
		return d.setUint(field, value)
	case reflect.Float32, reflect.Float64: //【checked】
		return d.setFloat(field, value)
	case reflect.Map:
		if m, ok := value.(map[string]interface{}); ok {
			return d.setMap(field, m)
		}
		return fmt.Errorf("%w: expected map, got %T", ErrInvalidValue, value)
	case reflect.Slice:
		if s, ok := value.([]interface{}); ok {
			return d.setSlice(field, s)
		}
		var x complex64 = 0
		fmt.Println(x)
		return fmt.Errorf("%w: expected slice, got %T", ErrInvalidValue, value)
	case reflect.String: //【checked】
		return d.setString(field, value)
	case reflect.Struct: //【checked】
		if m, ok := value.(map[string]interface{}); ok {
			return d.mapToStruct(m, field)
		}
		return fmt.Errorf("%w: expected map for struct, got %T", ErrInvalidValue, value)
	case reflect.Invalid, reflect.Complex64, reflect.Complex128, reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer: // 忽略的类型
		return fmt.Errorf("%w: %v", ErrUnsupportedType, field.Kind())
	default:
		return fmt.Errorf("%w: %v", ErrUnsupportedType, field.Kind())
	}
}

func (d *SmartDecoder) setSlice(field reflect.Value, slice []interface{}) error {
	newSlice := reflect.MakeSlice(field.Type(), len(slice), len(slice))
	for i, elem := range slice {
		elemValue := newSlice.Index(i)
		if err := d.setValue(elemValue, elem); err != nil {
			return fmt.Errorf("element %d: %w", i, err)
		}
	}
	// 设置切片值
	field.Set(newSlice)
	return nil
}

func (d *SmartDecoder) setMap(field reflect.Value, m map[string]interface{}) error {
	// 创建新映射
	newMap := reflect.MakeMap(field.Type())
	// 获取键和值的类型
	keyType := field.Type().Key()
	valueType := field.Type().Elem()
	// 遍历映射元素
	for k, v := range m {
		// 创建键值
		key := reflect.ValueOf(k)
		// 如果键类型不是string，尝试转换
		if keyType.Kind() != reflect.String {
			var err error
			key, err = d.convertValue(key, keyType)
			if err != nil {
				return fmt.Errorf("key %q: %w", k, err)
			}
		}
		// 创建值
		value := reflect.New(valueType).Elem()
		if err := d.setValue(value, v); err != nil {
			return fmt.Errorf("value for key %q: %w", k, err)
		}
		// 设置映射项
		newMap.SetMapIndex(key, value)
	}
	// 设置映射值
	field.Set(newMap)
	return nil
}

// 设置整数值 interface{} -> int,允许整形、字符串、布尔值尝试转换为reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64
func (d *SmartDecoder) setInt(field reflect.Value, value interface{}) error {
	var intValue int64
	var err error
	switch v := value.(type) {
	case float64: // json的number类型转int(后面判断是否溢出)
		intValue = int64(v)
	case string: // 能转则转，不能转直接报错
		intValue, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidValue, err)
		}
	default:
		return fmt.Errorf("%w: cannot convert %T to %s", ErrInvalidValue, value, field.Type().Name())
	}
	// 检查是否超出范围
	if !isIntInRange(intValue, field.Kind()) {
		return fmt.Errorf("%w: value %d overflows type %v", ErrInvalidValue, intValue, field.Kind())
	}
	// 设置值
	field.SetInt(intValue)
	return nil
}

func (d *SmartDecoder) setUint(field reflect.Value, value interface{}) error {
	var uintValue uint64
	var err error

	switch v := value.(type) {
	case float64:
		if v < 0 {
			return fmt.Errorf("%w: negative value %f cannot be converted to uint", ErrInvalidValue, v)
		}
		uintValue = uint64(v)
	case string:
		uintValue, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidValue, err)
		}
	default:
		return fmt.Errorf("%w: cannot convert %T to uint", ErrInvalidValue, value)
	}

	if !isUintInRange(uintValue, field.Kind()) {
		return fmt.Errorf("%w: value %d overflows type %v", ErrInvalidValue, uintValue, field.Kind())
	}

	field.SetUint(uintValue)
	return nil
}

func (d *SmartDecoder) setFloat(field reflect.Value, value interface{}) error {
	var floatValue float64
	var err error

	switch v := value.(type) {
	case float64:
		floatValue = v
	case string:
		floatValue, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidValue, err)
		}
	default:
		return fmt.Errorf("%w: cannot convert %T to float", ErrInvalidValue, value)
	}

	field.SetFloat(floatValue)
	return nil
}

func (d *SmartDecoder) setBool(field reflect.Value, value interface{}) error {
	var boolValue bool
	switch v := value.(type) {
	case bool:
		boolValue = v
	case string: // 仅允许true|false转bool
		switch strings.ToLower(v) {
		case "true":
			boolValue = true
		case "false":
			boolValue = false
		default:
			return fmt.Errorf("%w: invalid boolean string %q", ErrInvalidValue, v)
		}
	default:
		// json: cannot unmarshal [string] into Go struct field [ListAirportCityReply.data.flight_station] of type [[]v1.AirCityFlightStation]
		return fmt.Errorf("cannot unmarshal %T(%v) into Go struct field type %s", value, v, field.Type())
	}
	field.SetBool(boolValue)
	return nil
}

func (d *SmartDecoder) setString(field reflect.Value, value interface{}) error {
	var stringValue string

	switch v := value.(type) {
	case nil:
		stringValue = ""
	case string:
		stringValue = v
	case float64:
		stringValue = strconv.FormatFloat(v, 'f', -1, 64) // 不添加多余的零
	case bool:
		stringValue = strconv.FormatBool(v)
	case []interface{}, map[string]interface{}:
		if bytes, err := json.Marshal(&v); err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidValue, err)
		} else {
			stringValue = string(bytes)
		}
	default:
		fmt.Printf("%T转string", value)
		stringValue = fmt.Sprintf("%v", v)
	}
	field.SetString(stringValue)
	return nil
}

func (d *SmartDecoder) convertValue(value reflect.Value, targetType reflect.Type) (reflect.Value, error) {
	newValue := reflect.New(targetType).Elem()
	if err := d.setValue(newValue, value.Interface()); err != nil {
		return reflect.Value{}, err
	}

	return newValue, nil
}

func isIntInRange(value int64, kind reflect.Kind) bool {
	switch kind {
	case reflect.Int8:
		return value >= math.MinInt8 && value <= math.MaxInt8
	case reflect.Int16:
		return value >= math.MinInt16 && value <= math.MaxInt16
	case reflect.Int32:
		return value >= math.MinInt32 && value <= math.MaxInt32
	case reflect.Int:
		return value >= math.MinInt && value <= math.MaxInt
	case reflect.Int64:
		return true
	default:
		return false
	}
}

func isUintInRange(value uint64, kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint8:
		return value <= math.MaxUint8
	case reflect.Uint16:
		return value <= math.MaxUint16
	case reflect.Uint32:
		return value <= math.MaxUint32
	case reflect.Uint:
		// 取决于平台，简化处理
		return value <= math.MaxUint
	case reflect.Uint64:
		return true
	default:
		return false
	}
}
