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
	Extend *extend // 拓展信息
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Extend: &extend{option: option},
	}
}

type extend struct {
	option *core.Option
}

// CreateExtendBatch 拓展信息批量创建
func (extend *extend) CreateExtendBatch(ctx context.Context, req *CreateExtendBatchApiReq, reqOption *core.ReqOption) (*CreateExtendBatchApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/ExtendInfo/BatchSync"

	apiResp, err := core.Request(ctx, apiReq, extend.option, reqOption)
	if err != nil {
		return nil, err
	}
	createExtendBatchApiResp := CreateExtendBatchApiResp{
		ApiResp: apiResp,
	}
	createExtendBatchApiReply := CreateExtendBatchApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := extend.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if extend.option.EnableEncryption && nil != extend.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch extend.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", extend.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(extend.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					extend.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createExtendBatchApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createExtendBatchApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createExtendBatchApiReply); err != nil {
				return nil, err
			}
		}
		createExtendBatchApiResp.CreateExtendBatchApiReply = &createExtendBatchApiReply
	}
	return &createExtendBatchApiResp, nil
}

// ListExtend 拓展信息查询
func (extend *extend) ListExtend(ctx context.Context, req *ListExtendApiReq, reqOption *core.ReqOption) (*ListExtendApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/ExtendInfo/Get"

	apiResp, err := core.Request(ctx, apiReq, extend.option, reqOption)
	if err != nil {
		return nil, err
	}
	listExtendApiResp := ListExtendApiResp{
		ApiResp: apiResp,
	}
	listExtendApiReply := ListExtendApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := extend.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if extend.option.EnableEncryption && nil != extend.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch extend.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", extend.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(extend.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					extend.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listExtendApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listExtendApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listExtendApiReply); err != nil {
				return nil, err
			}
		}
		listExtendApiResp.ListExtendApiReply = &listExtendApiReply
	}
	return &listExtendApiResp, nil
}

// UpdateExtendStatus 档案状态处理
func (extend *extend) UpdateExtendStatus(ctx context.Context, req *UpdateExtendStatusApiReq, reqOption *core.ReqOption) (*UpdateExtendStatusApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/ExtendInfo/Status"

	apiResp, err := core.Request(ctx, apiReq, extend.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateExtendStatusApiResp := UpdateExtendStatusApiResp{
		ApiResp: apiResp,
	}
	updateExtendStatusApiReply := UpdateExtendStatusApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := extend.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if extend.option.EnableEncryption && nil != extend.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch extend.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", extend.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(extend.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					extend.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateExtendStatusApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateExtendStatusApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateExtendStatusApiReply); err != nil {
				return nil, err
			}
		}
		updateExtendStatusApiResp.UpdateExtendStatusApiReply = &updateExtendStatusApiReply
	}
	return &updateExtendStatusApiResp, nil
}
