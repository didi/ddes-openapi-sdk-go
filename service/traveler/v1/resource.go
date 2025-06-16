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
	Traveler *traveler // 外部出行人
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Traveler: &traveler{option: option},
	}
}

type traveler struct {
	option *core.Option
}

// CreateTraveler 外部出行人新增【请求参数组装比较繁琐】
func (traveler *traveler) CreateTraveler(ctx context.Context, req *CreateTravelerApiReq, reqOption *core.ReqOption) (*CreateTravelerApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/traveler/create"

	apiResp, err := core.Request(ctx, apiReq, traveler.option, reqOption)
	if err != nil {
		return nil, err
	}
	createTravelerApiResp := CreateTravelerApiResp{
		ApiResp: apiResp,
	}
	createTravelerApiReply := CreateTravelerApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := traveler.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if traveler.option.EnableEncryption && nil != traveler.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch traveler.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", traveler.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(traveler.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					traveler.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createTravelerApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createTravelerApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createTravelerApiReply); err != nil {
				return nil, err
			}
		}
		createTravelerApiResp.CreateTravelerApiReply = &createTravelerApiReply
	}
	return &createTravelerApiResp, nil
}

// DelTraveler 外部出行人删除
func (traveler *traveler) DelTraveler(ctx context.Context, req *DelTravelerApiReq, reqOption *core.ReqOption) (*DelTravelerApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/traveler/del"

	apiResp, err := core.Request(ctx, apiReq, traveler.option, reqOption)
	if err != nil {
		return nil, err
	}
	delTravelerApiResp := DelTravelerApiResp{
		ApiResp: apiResp,
	}
	delTravelerApiReply := DelTravelerApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := traveler.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if traveler.option.EnableEncryption && nil != traveler.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch traveler.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", traveler.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(traveler.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					traveler.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &delTravelerApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &delTravelerApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &delTravelerApiReply); err != nil {
				return nil, err
			}
		}
		delTravelerApiResp.DelTravelerApiReply = &delTravelerApiReply
	}
	return &delTravelerApiResp, nil
}

// UpdateTraveler 外部出行人修改
func (traveler *traveler) UpdateTraveler(ctx context.Context, req *UpdateTravelerApiReq, reqOption *core.ReqOption) (*UpdateTravelerApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/traveler/update"

	apiResp, err := core.Request(ctx, apiReq, traveler.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateTravelerApiResp := UpdateTravelerApiResp{
		ApiResp: apiResp,
	}
	updateTravelerApiReply := UpdateTravelerApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := traveler.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if traveler.option.EnableEncryption && nil != traveler.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch traveler.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", traveler.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(traveler.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					traveler.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateTravelerApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateTravelerApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateTravelerApiReply); err != nil {
				return nil, err
			}
		}
		updateTravelerApiResp.UpdateTravelerApiReply = &updateTravelerApiReply
	}
	return &updateTravelerApiResp, nil
}
