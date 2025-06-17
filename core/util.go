package core

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func GenMD5EncryptionKey(companyId string) string {
	return MD5([]byte(fmt.Sprintf("es_traveler_%s", companyId)))
}
func GenSha256EncryptionKey(companyId string) string {
	return Sha256([]byte(fmt.Sprintf("es_traveler_%s", companyId)))
}

func MD5Base64(bytes []byte) string {
	hash := md5.New()
	hash.Write(bytes)
	md5Value := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(md5Value)
}
func MD5(bytes []byte) string {
	hash := md5.New()
	hash.Write(bytes)
	md5Value := hash.Sum(nil)
	return hex.EncodeToString(md5Value)
}
func Sha256(bytes []byte) string {
	hash := sha256.New()
	hash.Write([]byte(bytes))
	hashedData := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedData)
	return hashedString
}

func UnescapeUnicode(s string) string {
	var result []rune
	i := 0
	for i < len(s) {
		if s[i] == '\\' && i+5 < len(s) && s[i+1] == 'u' {
			var r rune
			_, err := fmt.Sscanf(s[i+2:i+6], "%x", &r)
			if err == nil {
				result = append(result, r)
				i += 6
			} else {
				result = append(result, rune(s[i]))
				i++
			}
		} else {
			result = append(result, rune(s[i]))
			i++
		}
	}
	return string(result)
}
