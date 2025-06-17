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
	AfterApproval *afterApproval // 行后审批
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		AfterApproval: &afterApproval{option: option},
	}
}

type afterApproval struct {
	option *core.Option
}

// CreatePersonalReceipt 行后审批结果同步（行后授权）
func (afterApproval *afterApproval) CreatePersonalReceipt(ctx context.Context, req *CreatePersonalReceiptApiReq, reqOption *core.ReqOption) (*CreatePersonalReceiptApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/AfterApproval/createPersonalReceipt"

	apiResp, err := core.Request(ctx, apiReq, afterApproval.option, reqOption)
	if err != nil {
		return nil, err
	}
	createPersonalReceiptApiResp := CreatePersonalReceiptApiResp{
		ApiResp: apiResp,
	}
	createPersonalReceiptApiReply := CreatePersonalReceiptApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := afterApproval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if afterApproval.option.EnableEncryption && nil != afterApproval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch afterApproval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", afterApproval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(afterApproval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					afterApproval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createPersonalReceiptApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createPersonalReceiptApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createPersonalReceiptApiReply); err != nil {
				return nil, err
			}
		}
		createPersonalReceiptApiResp.CreatePersonalReceiptApiReply = &createPersonalReceiptApiReply
	}
	return &createPersonalReceiptApiResp, nil
}

// GetPersonalReceiptOrder 行后审批完订单～查询个人付款单
func (afterApproval *afterApproval) GetPersonalReceiptOrder(ctx context.Context, req *GetPersonalReceiptOrderApiReq, reqOption *core.ReqOption) (*GetPersonalReceiptOrderApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/AfterApproval/getPersonalReceiptOrder"

	apiResp, err := core.Request(ctx, apiReq, afterApproval.option, reqOption)
	if err != nil {
		return nil, err
	}
	getPersonalReceiptOrderApiResp := GetPersonalReceiptOrderApiResp{
		ApiResp: apiResp,
	}
	getPersonalReceiptOrderApiReply := GetPersonalReceiptOrderApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := afterApproval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if afterApproval.option.EnableEncryption && nil != afterApproval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch afterApproval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", afterApproval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(afterApproval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					afterApproval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getPersonalReceiptOrderApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getPersonalReceiptOrderApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getPersonalReceiptOrderApiReply); err != nil {
				return nil, err
			}
		}
		getPersonalReceiptOrderApiResp.GetPersonalReceiptOrderApiReply = &getPersonalReceiptOrderApiReply
	}
	return &getPersonalReceiptOrderApiResp, nil
}
