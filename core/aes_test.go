package core

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestAES(t *testing.T) {
	companyId := "1125938266433647"
	// 生成128位加密需要的密钥
	toEncStr := fmt.Sprintf("es_traveler_%s", companyId)
	log.Println("待加密数据：", toEncStr)
	md5Key := MD5([]byte(toEncStr))
	t.Log("ent = 1, md5 key =", md5Key)
	// 创建一个新的SHA-256哈希对象
	hash := sha256.New()
	// 向哈希对象写入数据
	hash.Write([]byte(toEncStr))
	// 计算哈希值
	hashedData := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashedString := hex.EncodeToString(hashedData)
	// 输出结果
	fmt.Printf("ent = 2, sha256 key: %s\n", hashedString)

}

// 测试128位加密和解密
func TestAESEncryptECB(t *testing.T) {
	// 明文
	plaintext := []byte("Hello, AES ECB encryption!")

	// 128位密钥（16 字节）
	key128 := []byte("0123456789abcdef")

	// 128位加密
	ciphertext128, err := AESEncryptECB(plaintext, key128)
	if err != nil {
		fmt.Println("128位加密出错:", err)
		return
	}
	encryptedBase64_128 := base64.StdEncoding.EncodeToString(ciphertext128)
	fmt.Println("128位加密后的 Base64 字符串:", encryptedBase64_128)

	// 128位解密
	decodedCiphertext128, err := base64.StdEncoding.DecodeString(encryptedBase64_128)
	if err != nil {
		fmt.Println("128位 Base64 解码出错:", err)
		return
	}
	decryptedText128, err := AESDecryptECB(decodedCiphertext128, key128)
	if err != nil {
		fmt.Println("128位解密出错:", err)
		return
	}
	fmt.Println("128位解密后的明文:", string(decryptedText128))
}

// 测试解密
func TestAESDecryptECB(t *testing.T) {
	key := []byte("mVx22PwqebjHcDFNhStIyw==")
	encryptedBase64 := "bWYNcf3HhrboIUKRSuJ9muU0MBv6M7Y1qOybY+SsJ1TV5OhmGC9AUpZVwbVVFHNKbJDkukGlfAjXjv/WT9VYJf6ppk1gliiA9GtBKSk1JrIr1j6E57x9cXTCk/KKFY/uNhj7fTgca2czJzwH2+tbVTidbyTJ3CGGRn/JhA3DqaGfdHB/dvm5fU9Rn0BPbBQRfvz7f3RIuUIaiFj0NeZOcQ=="
	fmt.Println("加密后的 Base64 字符串:", encryptedBase64)

	// 解密
	decodedCiphertext128, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		fmt.Println(" Base64 解码出错:", err)
		return
	}
	decryptedText128, err := AESDecryptECB(decodedCiphertext128, key)
	if err != nil {
		fmt.Println("解密出错:", err)
		return
	}
	fmt.Println("解密后的明文:", string(decryptedText128))
}

func TestGenerateSecretKeyBytes(t *testing.T) {
	//key := "dc9a4b891d3907ad6c0159c2bf87f3b560d199ef006707ca670bb4054cbb767c"
	key := "b542e944ed911235b2e5bb27a349943508a4c995dabaad14e665ecc9c392583a"

	t.Log(len([]byte(key)))

	got, err := GenerateSecretKeyBytes(2, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(got))
	toString := hex.EncodeToString(got)
	t.Log(toString)
}
