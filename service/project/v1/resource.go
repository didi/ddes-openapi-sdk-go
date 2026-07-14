package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

type V1 struct {
	Project *project // 项目
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Project: &project{option: option},
	}
}

type project struct {
	option *core.Option
}

// GetProjectDetail 查询项目下的人员关联信息
func (p *project) GetProjectDetail(ctx context.Context, req *GetProjectDetailApiReq, reqOption *core.ReqOption) (*GetProjectDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Project/detail"

	apiResp, err := core.Request(ctx, apiReq, p.option, reqOption)
	if err != nil {
		return nil, err
	}
	getProjectDetailApiResp := GetProjectDetailApiResp{
		ApiResp: apiResp,
	}
	getProjectDetailApiReply := GetProjectDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := p.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if p.option.EnableEncryption && nil != p.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch p.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", p.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(p.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					p.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getProjectDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getProjectDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getProjectDetailApiReply); err != nil {
				return nil, err
			}
		}
		getProjectDetailApiResp.GetProjectDetailApiReply = &getProjectDetailApiReply
	}
	return &getProjectDetailApiResp, nil
}

// OutTravelerList 查询项目外部出行人列表
func (p *project) OutTravelerList(ctx context.Context, req *OutTravelerListApiReq, reqOption *core.ReqOption) (*OutTravelerListApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/open-apis/v2/project/outTravelerList"

	apiResp, err := core.Request(ctx, apiReq, p.option, reqOption)
	if err != nil {
		return nil, err
	}
	outTravelerListApiResp := OutTravelerListApiResp{
		ApiResp: apiResp,
	}
	outTravelerListApiReply := OutTravelerListApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := p.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if p.option.EnableEncryption && nil != p.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch p.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", p.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(p.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					p.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &outTravelerListApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &outTravelerListApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &outTravelerListApiReply); err != nil {
				return nil, err
			}
		}
		outTravelerListApiResp.OutTravelerListApiReply = &outTravelerListApiReply
	}
	return &outTravelerListApiResp, nil
}

// UpdateMember 绑定项目与人员关系
func (p *project) UpdateMember(ctx context.Context, req *UpdateMemberApiReq, reqOption *core.ReqOption) (*UpdateMemberApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Project/updateMember"

	apiResp, err := core.Request(ctx, apiReq, p.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateMemberApiResp := UpdateMemberApiResp{
		ApiResp: apiResp,
	}
	updateMemberApiReply := UpdateMemberApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := p.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if p.option.EnableEncryption && nil != p.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch p.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", p.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(p.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					p.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateMemberApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateMemberApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateMemberApiReply); err != nil {
				return nil, err
			}
		}
		updateMemberApiResp.UpdateMemberApiReply = &updateMemberApiReply
	}
	return &updateMemberApiResp, nil
}

// DelMember 删除项目与人员关系
func (p *project) DelMember(ctx context.Context, req *DelMemberApiReq, reqOption *core.ReqOption) (*DelMemberApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Project/delMember"

	apiResp, err := core.Request(ctx, apiReq, p.option, reqOption)
	if err != nil {
		return nil, err
	}
	delMemberApiResp := DelMemberApiResp{
		ApiResp: apiResp,
	}
	delMemberApiReply := DelMemberApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := p.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if p.option.EnableEncryption && nil != p.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch p.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", p.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(p.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					p.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &delMemberApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &delMemberApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &delMemberApiReply); err != nil {
				return nil, err
			}
		}
		delMemberApiResp.DelMemberApiReply = &delMemberApiReply
	}
	return &delMemberApiResp, nil
}
