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
	LegalEntity *legalEntity // 公司主体
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		LegalEntity: &legalEntity{option: option},
	}
}

type legalEntity struct {
	option *core.Option
}

// CreateLegalEntity 公司主体新增
func (legalEntity *legalEntity) CreateLegalEntity(ctx context.Context, req *CreateLegalEntityApiReq, reqOption *core.ReqOption) (*CreateLegalEntityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/LegalEntity/add"

	apiResp, err := core.Request(ctx, apiReq, legalEntity.option, reqOption)
	if err != nil {
		return nil, err
	}
	createLegalEntityApiResp := CreateLegalEntityApiResp{
		ApiResp: apiResp,
	}
	createLegalEntityApiReply := CreateLegalEntityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := legalEntity.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if legalEntity.option.EnableEncryption && nil != legalEntity.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch legalEntity.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", legalEntity.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(legalEntity.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					legalEntity.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createLegalEntityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createLegalEntityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createLegalEntityApiReply); err != nil {
				return nil, err
			}
		}
		createLegalEntityApiResp.CreateLegalEntityApiReply = &createLegalEntityApiReply
	}
	return &createLegalEntityApiResp, nil
}

// DelLegalEntity 公司主体停用
func (legalEntity *legalEntity) DelLegalEntity(ctx context.Context, req *DelLegalEntityApiReq, reqOption *core.ReqOption) (*DelLegalEntityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/LegalEntity/del"

	apiResp, err := core.Request(ctx, apiReq, legalEntity.option, reqOption)
	if err != nil {
		return nil, err
	}
	delLegalEntityApiResp := DelLegalEntityApiResp{
		ApiResp: apiResp,
	}
	delLegalEntityApiReply := DelLegalEntityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := legalEntity.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if legalEntity.option.EnableEncryption && nil != legalEntity.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch legalEntity.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", legalEntity.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(legalEntity.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					legalEntity.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &delLegalEntityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &delLegalEntityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &delLegalEntityApiReply); err != nil {
				return nil, err
			}
		}
		delLegalEntityApiResp.DelLegalEntityApiReply = &delLegalEntityApiReply
	}
	return &delLegalEntityApiResp, nil
}

// GetLegalEntity 公司主体查询
func (legalEntity *legalEntity) GetLegalEntity(ctx context.Context, req *GetLegalEntityApiReq, reqOption *core.ReqOption) (*GetLegalEntityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/LegalEntity/get"

	apiResp, err := core.Request(ctx, apiReq, legalEntity.option, reqOption)
	if err != nil {
		return nil, err
	}
	getLegalEntityApiResp := GetLegalEntityApiResp{
		ApiResp: apiResp,
	}
	getLegalEntityApiReply := GetLegalEntityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := legalEntity.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if legalEntity.option.EnableEncryption && nil != legalEntity.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch legalEntity.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", legalEntity.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(legalEntity.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					legalEntity.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getLegalEntityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getLegalEntityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getLegalEntityApiReply); err != nil {
				return nil, err
			}
		}
		getLegalEntityApiResp.GetLegalEntityApiReply = &getLegalEntityApiReply
	}
	return &getLegalEntityApiResp, nil
}

// UpdateLegalEntity 公司主体修改
func (legalEntity *legalEntity) UpdateLegalEntity(ctx context.Context, req *UpdateLegalEntityApiReq, reqOption *core.ReqOption) (*UpdateLegalEntityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/LegalEntity/edit"

	apiResp, err := core.Request(ctx, apiReq, legalEntity.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateLegalEntityApiResp := UpdateLegalEntityApiResp{
		ApiResp: apiResp,
	}
	updateLegalEntityApiReply := UpdateLegalEntityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := legalEntity.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if legalEntity.option.EnableEncryption && nil != legalEntity.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch legalEntity.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", legalEntity.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(legalEntity.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					legalEntity.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateLegalEntityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateLegalEntityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateLegalEntityApiReply); err != nil {
				return nil, err
			}
		}
		updateLegalEntityApiResp.UpdateLegalEntityApiReply = &updateLegalEntityApiReply
	}
	return &updateLegalEntityApiResp, nil
}
