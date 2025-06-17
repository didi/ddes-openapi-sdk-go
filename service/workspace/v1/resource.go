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
	Workspace *workspace //
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Workspace: &workspace{option: option},
	}
}

type workspace struct {
	option *core.Option
}

// CreateWorkplace 地点新增
func (workspace *workspace) CreateWorkplace(ctx context.Context, req *CreateWorkplaceApiReq, reqOption *core.ReqOption) (*CreateWorkplaceApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/workplace/create"

	apiResp, err := core.Request(ctx, apiReq, workspace.option, reqOption)
	if err != nil {
		return nil, err
	}
	createWorkplaceApiResp := CreateWorkplaceApiResp{
		ApiResp: apiResp,
	}
	createWorkplaceApiReply := CreateWorkplaceApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := workspace.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if workspace.option.EnableEncryption && nil != workspace.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch workspace.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", workspace.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(workspace.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					workspace.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createWorkplaceApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createWorkplaceApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createWorkplaceApiReply); err != nil {
				return nil, err
			}
		}
		createWorkplaceApiResp.CreateWorkplaceApiReply = &createWorkplaceApiReply
	}
	return &createWorkplaceApiResp, nil
}

// DeleteWorkplace 地点删除
func (workspace *workspace) DeleteWorkplace(ctx context.Context, req *DeleteWorkplaceApiReq, reqOption *core.ReqOption) (*DeleteWorkplaceApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/workplace/del"

	apiResp, err := core.Request(ctx, apiReq, workspace.option, reqOption)
	if err != nil {
		return nil, err
	}
	deleteWorkplaceApiResp := DeleteWorkplaceApiResp{
		ApiResp: apiResp,
	}
	deleteWorkplaceApiReply := DeleteWorkplaceApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := workspace.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if workspace.option.EnableEncryption && nil != workspace.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch workspace.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", workspace.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(workspace.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					workspace.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &deleteWorkplaceApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &deleteWorkplaceApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &deleteWorkplaceApiReply); err != nil {
				return nil, err
			}
		}
		deleteWorkplaceApiResp.DeleteWorkplaceApiReply = &deleteWorkplaceApiReply
	}
	return &deleteWorkplaceApiResp, nil
}

// UpdateWorkplace 地点修改
func (workspace *workspace) UpdateWorkplace(ctx context.Context, req *UpdateWorkplaceApiReq, reqOption *core.ReqOption) (*UpdateWorkplaceApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/workplace/update"

	apiResp, err := core.Request(ctx, apiReq, workspace.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateWorkplaceApiResp := UpdateWorkplaceApiResp{
		ApiResp: apiResp,
	}
	updateWorkplaceApiReply := UpdateWorkplaceApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := workspace.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if workspace.option.EnableEncryption && nil != workspace.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch workspace.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", workspace.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(workspace.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					workspace.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateWorkplaceApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateWorkplaceApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateWorkplaceApiReply); err != nil {
				return nil, err
			}
		}
		updateWorkplaceApiResp.UpdateWorkplaceApiReply = &updateWorkplaceApiReply
	}
	return &updateWorkplaceApiResp, nil
}
