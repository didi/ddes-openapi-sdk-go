package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/http"
)

type V1 struct {
	Login *login //
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Login: &login{option: option},
	}
}

type login struct {
	option *core.Option
}

// GetLoginEncryptStr 单点请求
func (login *login) GetLoginEncryptStr(ctx context.Context, req *GetLoginEncryptStrApiReq, reqOption *core.ReqOption) (*GetLoginEncryptStrApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Login/getLoginEncryptStr"

	apiResp, err := core.Request(ctx, apiReq, login.option, reqOption)
	if err != nil {
		return nil, err
	}
	getLoginEncryptStrApiResp := GetLoginEncryptStrApiResp{
		ApiResp: apiResp,
	}
	getLoginEncryptStrApiReply := GetLoginEncryptStrApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := login.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if login.option.EnableEncryption && nil != login.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch login.option.EncryptionOption.Ent {
					case 1:
						decodedCiphertext128, err := base64.StdEncoding.DecodeString(encryptData)
						if err != nil {
							return nil, err
						}
						decodedCiphertext = decodedCiphertext128
					case 2:
						decodedCiphertext128, err := base64.URLEncoding.DecodeString(encryptData)
						if err != nil {
							return nil, err
						}
						decodedCiphertext = decodedCiphertext128
					default:
						return nil, fmt.Errorf("未支持的Ent：%d", login.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(login.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					login.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getLoginEncryptStrApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getLoginEncryptStrApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getLoginEncryptStrApiReply); err != nil {
				return nil, err
			}
		}
		getLoginEncryptStrApiResp.GetLoginEncryptStrApiReply = &getLoginEncryptStrApiReply
	}
	return &getLoginEncryptStrApiResp, nil
}
