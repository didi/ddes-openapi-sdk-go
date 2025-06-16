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
	Auth *auth // 接口认证
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Auth: &auth{option: option},
	}
}

type auth struct {
	option *core.Option
}

// Authorize 授权认证
func (auth *auth) Authorize(ctx context.Context, req *AuthorizeApiReq, reqOption *core.ReqOption) (*AuthorizeApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Auth/authorize"

	apiResp, err := core.Request(ctx, apiReq, auth.option, reqOption)
	if err != nil {
		return nil, err
	}
	authorizeApiResp := AuthorizeApiResp{
		ApiResp: apiResp,
	}
	authorizeApiReply := AuthorizeApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := auth.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if auth.option.EnableEncryption && nil != auth.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch auth.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", auth.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(auth.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					auth.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &authorizeApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &authorizeApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &authorizeApiReply); err != nil {
				return nil, err
			}
		}
		authorizeApiResp.AuthorizeApiReply = &authorizeApiReply
	}
	return &authorizeApiResp, nil
}
