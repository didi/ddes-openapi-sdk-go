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
	BudgetCenter *budgetCenter // 部门或项目
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		BudgetCenter: &budgetCenter{option: option},
	}
}

type budgetCenter struct {
	option *core.Option
}

// CreateBudgetCenter 部门或项目新增
func (budgetCenter *budgetCenter) CreateBudgetCenter(ctx context.Context, req *CreateBudgetCenterApiReq, reqOption *core.ReqOption) (*CreateBudgetCenterApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/BudgetCenter/add"

	apiResp, err := core.Request(ctx, apiReq, budgetCenter.option, reqOption)
	if err != nil {
		return nil, err
	}
	createBudgetCenterApiResp := CreateBudgetCenterApiResp{
		ApiResp: apiResp,
	}
	createBudgetCenterApiReply := CreateBudgetCenterApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := budgetCenter.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if budgetCenter.option.EnableEncryption && nil != budgetCenter.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch budgetCenter.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", budgetCenter.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(budgetCenter.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					budgetCenter.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createBudgetCenterApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createBudgetCenterApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createBudgetCenterApiReply); err != nil {
				return nil, err
			}
		}
		createBudgetCenterApiResp.CreateBudgetCenterApiReply = &createBudgetCenterApiReply
	}
	return &createBudgetCenterApiResp, nil
}

// DelBudgetCenter 部门或项目停用
func (budgetCenter *budgetCenter) DelBudgetCenter(ctx context.Context, req *DelBudgetCenterApiReq, reqOption *core.ReqOption) (*DelBudgetCenterApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/BudgetCenter/del"

	apiResp, err := core.Request(ctx, apiReq, budgetCenter.option, reqOption)
	if err != nil {
		return nil, err
	}
	delBudgetCenterApiResp := DelBudgetCenterApiResp{
		ApiResp: apiResp,
	}
	delBudgetCenterApiReply := DelBudgetCenterApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := budgetCenter.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if budgetCenter.option.EnableEncryption && nil != budgetCenter.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch budgetCenter.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", budgetCenter.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(budgetCenter.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					budgetCenter.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &delBudgetCenterApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &delBudgetCenterApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &delBudgetCenterApiReply); err != nil {
				return nil, err
			}
		}
		delBudgetCenterApiResp.DelBudgetCenterApiReply = &delBudgetCenterApiReply
	}
	return &delBudgetCenterApiResp, nil
}

// GetBudgetCenter 部门或项目查询
func (budgetCenter *budgetCenter) GetBudgetCenter(ctx context.Context, req *GetBudgetCenterApiReq, reqOption *core.ReqOption) (*GetBudgetCenterApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/BudgetCenter/get"

	apiResp, err := core.Request(ctx, apiReq, budgetCenter.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBudgetCenterApiResp := GetBudgetCenterApiResp{
		ApiResp: apiResp,
	}
	getBudgetCenterApiReply := GetBudgetCenterApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := budgetCenter.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if budgetCenter.option.EnableEncryption && nil != budgetCenter.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch budgetCenter.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", budgetCenter.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(budgetCenter.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					budgetCenter.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBudgetCenterApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBudgetCenterApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBudgetCenterApiReply); err != nil {
				return nil, err
			}
		}
		getBudgetCenterApiResp.GetBudgetCenterApiReply = &getBudgetCenterApiReply
	}
	return &getBudgetCenterApiResp, nil
}

// UpdateBudgetCenter 部门或项目修改
func (budgetCenter *budgetCenter) UpdateBudgetCenter(ctx context.Context, req *UpdateBudgetCenterApiReq, reqOption *core.ReqOption) (*UpdateBudgetCenterApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/BudgetCenter/edit"

	apiResp, err := core.Request(ctx, apiReq, budgetCenter.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateBudgetCenterApiResp := UpdateBudgetCenterApiResp{
		ApiResp: apiResp,
	}
	updateBudgetCenterApiReply := UpdateBudgetCenterApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := budgetCenter.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if budgetCenter.option.EnableEncryption && nil != budgetCenter.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch budgetCenter.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", budgetCenter.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(budgetCenter.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					budgetCenter.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateBudgetCenterApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateBudgetCenterApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateBudgetCenterApiReply); err != nil {
				return nil, err
			}
		}
		updateBudgetCenterApiResp.UpdateBudgetCenterApiReply = &updateBudgetCenterApiReply
	}
	return &updateBudgetCenterApiResp, nil
}
