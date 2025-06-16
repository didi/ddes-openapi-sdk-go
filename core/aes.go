package core

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

// PKCS7Padding 填充函数，将数据填充到块大小的整数倍
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7UnPadding 去填充函数，去除填充的数据
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// AESEncryptECB 使用 AES ECB 模式进行加密
func AESEncryptECB(plaintext []byte, key []byte) ([]byte, error) {
	if len(key) == 64 {
		//keyBytes, err := GenerateSecretKeyBytes(2, string(key))
		//if err != nil {
		//	return nil, err
		//}
		//key = keyBytes

		newKey, err := hex.DecodeString(string(key))
		if err != nil {
			return nil, err
		}
		key = newKey
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))
	mode := NewECBEncryptor(block)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// AESDecryptECB 使用 AES ECB 模式进行解密
func AESDecryptECB(ciphertext []byte, key []byte) ([]byte, error) {
	if len(key) == 64 {
		newKey, err := hex.DecodeString(string(key))
		if err != nil {
			return nil, err
		}
		key = newKey
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	mode := NewECBDecryptor(block)
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = PKCS7UnPadding(plaintext)
	return plaintext, nil
}

// NewECBEncryptor 创建 ECB 模式的加密器
func NewECBEncryptor(b cipher.Block) cipher.BlockMode {
	return ecbEncryptor{block: b}
}

type ecbEncryptor struct {
	block cipher.Block
}

func (x ecbEncryptor) BlockSize() int { return x.block.BlockSize() }

func (x ecbEncryptor) CryptBlocks(dst, src []byte) {
	for len(src) > 0 {
		x.block.Encrypt(dst, src[:x.block.BlockSize()])
		src = src[x.block.BlockSize():]
		dst = dst[x.block.BlockSize():]
	}
}

// NewECBDecryptor 创建 ECB 模式的解密器
func NewECBDecryptor(b cipher.Block) cipher.BlockMode {
	return ecbDecryptor{block: b}
}

type ecbDecryptor struct {
	block cipher.Block
}

func (x ecbDecryptor) BlockSize() int { return x.block.BlockSize() }

func (x ecbDecryptor) CryptBlocks(dst, src []byte) {
	for len(src) > 0 {
		x.block.Decrypt(dst, src[:x.block.BlockSize()])
		src = src[x.block.BlockSize():]
		dst = dst[x.block.BlockSize():]
	}
}

// ################# 222
//
//	type AesTool struct {
//		key string
//	}
//
//	func NewAESTool(k string) *AesTool {
//		return &AesTool{
//			key: k,
//		}
//	}
//
//	func (a *AesTool) ECBEncrypt128(data interface{}) (s string, err error) {
//		defer func() {
//			if e := recover(); e != nil {
//				err = errors.New(e.(string))
//			}
//		}()
//		if data == nil {
//			return "", nil
//		}
//		// 类型转换处理器
//		var src []byte
//		switch v := data.(type) {
//		case []byte:
//			src = v
//		case string:
//			src = []byte(v)
//		default:
//			rv := reflect.Indirect(reflect.ValueOf(v))
//			if rv.Kind() == reflect.Struct {
//				jsonData, err := json.Marshal(v)
//				if err != nil {
//					return "", errors.New("json序列化失败")
//				}
//				src = jsonData
//			} else {
//				return "", errors.New("不支持的数据类型")
//			}
//		}
//		keyBytes, err := GenerateSecretKeyBytes(1, a.key)
//		if err != nil {
//			return "", err
//		}
//
//		ciphertext := PKCS7Padding([]byte(src), 16)
//
//		//ECB加密
//		crypted, err := AesEncrypt(ciphertext, keyBytes)
//		if err != nil {
//			return "", err
//		}
//
//		response := base64.StdEncoding.EncodeToString(crypted)
//		return response, nil
//	}
//
//	func (a *AesTool) ECBDecrypt128(src string) (s string, err error) {
//		defer func() {
//			if e := recover(); e != nil {
//				err = errors.New(e.(string))
//			}
//		}()
//		replacer := strings.NewReplacer("/", "_", "+", "-")
//		src = replacer.Replace(src)
//		// 自动补全填充符
//		if len(src)%4 != 0 {
//			src += strings.Repeat("=", (4-len(src)%4)%4)
//		}
//		enc, err := base64.URLEncoding.DecodeString(src)
//		if err != nil {
//			return "", err
//		}
//		keyBytes, err := GenerateSecretKeyBytes(1, a.key)
//		if err != nil {
//			return "", err
//		}
//		// ECB解密
//		origin, err := AesDecrypt(enc, []byte(keyBytes))
//		if err != nil {
//			return "", err
//		}
//
//		// 使用PKCS#7对解密后的内容去除填充
//		response := string(PKCS7UnPadding(origin))
//		return response, nil
//	}
//
//	func AesEncrypt(src, key []byte) ([]byte, error) {
//		Block, err := aes.NewCipher(key)
//		if err != nil {
//			return nil, err
//		}
//		if len(src) == 0 {
//			return nil, errors.New("plaintext empty")
//		}
//		mode := NewECBEncrypter(Block)
//		ciphertext := src
//		mode.CryptBlocks(ciphertext, ciphertext)
//		return ciphertext, nil
//	}
//
// // AesDecrypt Aes/ECB模式的解密方法，PKCS7填充方式
//
//	func AesDecrypt(src, key []byte) ([]byte, error) {
//		Block, err := aes.NewCipher(key)
//		if err != nil {
//			return nil, err
//		}
//		if len(src) == 0 {
//			return nil, errors.New("plaintext empty")
//		}
//		mode := NewECBDecrypter(Block)
//		ciphertext := src
//		mode.CryptBlocks(ciphertext, ciphertext)
//		return ciphertext, nil
//	}
//
// // ECB模式结构体
//
//	type ecb struct {
//		b         cipher.Block
//		blockSize int
//	}
//
// // 实例化ECB对象
//
//	func newECB(b cipher.Block) *ecb {
//		return &ecb{
//			b:         b,
//			blockSize: b.BlockSize(),
//		}
//	}
//
// // ECB加密类
// type ecbEncrypter ecb
//
//	func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
//		return (*ecbEncrypter)(newECB(b))
//	}
//
//	func (x *ecbEncrypter) BlockSize() int {
//		return x.blockSize
//	}
//
//	func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
//		if len(src)%x.blockSize != 0 {
//			panic("crypto/cipher: input not full blocks")
//		}
//		if len(dst) < len(src) {
//			panic("crypto/cipher: output smaller than input")
//		}
//		for len(src) > 0 {
//			x.b.Encrypt(dst, src[:x.blockSize])
//			dst = dst[x.blockSize:]
//			src = src[x.blockSize:]
//		}
//	}
//
// // ECB解密类
// type ecbDecrypter ecb
//
//	func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
//		return (*ecbDecrypter)(newECB(b))
//	}
//
//	func (x *ecbDecrypter) BlockSize() int {
//		return x.blockSize
//	}
//
//	func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
//		if len(src)%x.blockSize != 0 {
//			panic("crypto/cipher: input not full blocks")
//		}
//		if len(dst) < len(src) {
//			panic("crypto/cipher: output smaller than input")
//		}
//		for len(src) > 0 {
//			x.b.Decrypt(dst, src[:x.blockSize])
//			dst = dst[x.blockSize:]
//			src = src[x.blockSize:]
//		}
//	}
//
// // PKCS7Padding PKCS7填充
//
//	func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
//		padding := blockSize - len(ciphertext)%blockSize
//		padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//		return append(ciphertext, padtext...)
//	}
//
// // PKCS7UnPadding PKCS7去除
//
//	func PKCS7UnPadding(ciphertext []byte) []byte {
//		length := len(ciphertext)
//		unpadding := int(ciphertext[length-1])
//		step := length - unpadding
//		if step < 0 {
//			return []byte{}
//		}
//		return ciphertext[:step]
//	}

// GenerateSecretKeyBytes 生成符合AES标准的密钥 ent=0返回空字符串，ent=1使用MD5(128bit)，ent=2使用SHA256(256bit)
func GenerateSecretKeyBytes(ent int64, key string) ([]byte, error) {
	switch ent {
	case 0:
		return []byte{}, nil
	case 1:
		return []byte(key), nil
	case 2:
		return HexToBin(key)
	default:
		return nil, errors.New("生成秘钥失败：ent只能是0、1、2")
	}
}

func HexToBin(hexStr string) ([]byte, error) {
	if len(hexStr)%2 != 0 {
		return nil, errors.New("密钥长度必须是128或者256")
	}
	binData, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, errors.New("密钥长度必须是128或者256")
	}
	return binData, nil
}
