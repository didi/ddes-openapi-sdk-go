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
	OutApproval *outApproval // 外部审批
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		OutApproval: &outApproval{option: option},
	}
}

type outApproval struct {
	option *core.Option
}

// UpdateOutApprovalStatus 外部通知审批单状态变更接口【此接口文档未标注返回值，待验证】
func (outApproval *outApproval) UpdateOutApprovalStatus(ctx context.Context, req *UpdateOutApprovalStatusApiReq, reqOption *core.ReqOption) (*UpdateOutApprovalStatusApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/OutApproval/Status"

	apiResp, err := core.Request(ctx, apiReq, outApproval.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateOutApprovalStatusApiResp := UpdateOutApprovalStatusApiResp{
		ApiResp: apiResp,
	}
	updateOutApprovalStatusApiReply := UpdateOutApprovalStatusApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := outApproval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if outApproval.option.EnableEncryption && nil != outApproval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch outApproval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", outApproval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(outApproval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					outApproval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateOutApprovalStatusApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateOutApprovalStatusApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateOutApprovalStatusApiReply); err != nil {
				return nil, err
			}
		}
		updateOutApprovalStatusApiResp.UpdateOutApprovalStatusApiReply = &updateOutApprovalStatusApiReply
	}
	return &updateOutApprovalStatusApiResp, nil
}
