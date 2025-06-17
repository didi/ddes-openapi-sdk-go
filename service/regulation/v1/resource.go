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
	Regulation *regulation // 制度
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Regulation: &regulation{option: option},
	}
}

type regulation struct {
	option *core.Option
}

// GetRegulation 制度详情【原始JSON示例与文档对不上，待验证】
func (regulation *regulation) GetRegulation(ctx context.Context, req *GetRegulationApiReq, reqOption *core.ReqOption) (*GetRegulationApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Regulation/detail"

	apiResp, err := core.Request(ctx, apiReq, regulation.option, reqOption)
	if err != nil {
		return nil, err
	}
	getRegulationApiResp := GetRegulationApiResp{
		ApiResp: apiResp,
	}
	getRegulationApiReply := GetRegulationApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := regulation.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if regulation.option.EnableEncryption && nil != regulation.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch regulation.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", regulation.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(regulation.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					regulation.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getRegulationApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getRegulationApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getRegulationApiReply); err != nil {
				return nil, err
			}
		}
		getRegulationApiResp.GetRegulationApiReply = &getRegulationApiReply
	}
	return &getRegulationApiResp, nil
}

// ListRegulation 制度列表
func (regulation *regulation) ListRegulation(ctx context.Context, req *ListRegulationApiReq, reqOption *core.ReqOption) (*ListRegulationApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Regulation/get"

	apiResp, err := core.Request(ctx, apiReq, regulation.option, reqOption)
	if err != nil {
		return nil, err
	}
	listRegulationApiResp := ListRegulationApiResp{
		ApiResp: apiResp,
	}
	listRegulationApiReply := ListRegulationApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := regulation.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if regulation.option.EnableEncryption && nil != regulation.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch regulation.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", regulation.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(regulation.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					regulation.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listRegulationApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listRegulationApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listRegulationApiReply); err != nil {
				return nil, err
			}
		}
		listRegulationApiResp.ListRegulationApiReply = &listRegulationApiReply
	}
	return &listRegulationApiResp, nil
}
