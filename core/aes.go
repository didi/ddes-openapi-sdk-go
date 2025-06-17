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
