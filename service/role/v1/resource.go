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
	Role *role // 角色
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Role: &role{option: option},
	}
}

type role struct {
	option *core.Option
}

// ListRole 角色查询【返回data字段为数组，特殊处理】
func (role *role) ListRole(ctx context.Context, req *ListRoleApiReq, reqOption *core.ReqOption) (*ListRoleApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Role/get"

	apiResp, err := core.Request(ctx, apiReq, role.option, reqOption)
	if err != nil {
		return nil, err
	}
	listRoleApiResp := ListRoleApiResp{
		ApiResp: apiResp,
	}
	listRoleApiReply := ListRoleApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := role.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if role.option.EnableEncryption && nil != role.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch role.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", role.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(role.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					role.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listRoleApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listRoleApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listRoleApiReply); err != nil {
				return nil, err
			}
		}
		listRoleApiResp.ListRoleApiReply = &listRoleApiReply
	}
	return &listRoleApiResp, nil
}
