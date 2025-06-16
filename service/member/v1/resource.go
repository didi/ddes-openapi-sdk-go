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
	Member *member // 用户
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Member: &member{option: option},
	}
}

type member struct {
	option *core.Option
}

// CreateMember 用户新增
func (member *member) CreateMember(ctx context.Context, req *CreateMemberApiReq, reqOption *core.ReqOption) (*CreateMemberApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Member/single"

	apiResp, err := core.Request(ctx, apiReq, member.option, reqOption)
	if err != nil {
		return nil, err
	}
	createMemberApiResp := CreateMemberApiResp{
		ApiResp: apiResp,
	}
	createMemberApiReply := CreateMemberApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := member.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if member.option.EnableEncryption && nil != member.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch member.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", member.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(member.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					member.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createMemberApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createMemberApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createMemberApiReply); err != nil {
				return nil, err
			}
		}
		createMemberApiResp.CreateMemberApiReply = &createMemberApiReply
	}
	return &createMemberApiResp, nil
}

// DelMember 用户删除
func (member *member) DelMember(ctx context.Context, req *DelMemberApiReq, reqOption *core.ReqOption) (*DelMemberApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Member/del"

	apiResp, err := core.Request(ctx, apiReq, member.option, reqOption)
	if err != nil {
		return nil, err
	}
	delMemberApiResp := DelMemberApiResp{
		ApiResp: apiResp,
	}
	delMemberApiReply := DelMemberApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := member.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if member.option.EnableEncryption && nil != member.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch member.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", member.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(member.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					member.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
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

// GetMemberDetail 员工明细
func (member *member) GetMemberDetail(ctx context.Context, req *GetMemberDetailApiReq, reqOption *core.ReqOption) (*GetMemberDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Member/detail"

	apiResp, err := core.Request(ctx, apiReq, member.option, reqOption)
	if err != nil {
		return nil, err
	}
	getMemberDetailApiResp := GetMemberDetailApiResp{
		ApiResp: apiResp,
	}
	getMemberDetailApiReply := GetMemberDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := member.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if member.option.EnableEncryption && nil != member.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch member.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", member.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(member.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					member.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getMemberDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getMemberDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getMemberDetailApiReply); err != nil {
				return nil, err
			}
		}
		getMemberDetailApiResp.GetMemberDetailApiReply = &getMemberDetailApiReply
	}
	return &getMemberDetailApiResp, nil
}

// GetMemberQuota 员工限额查询
func (member *member) GetMemberQuota(ctx context.Context, req *GetMemberQuotaApiReq, reqOption *core.ReqOption) (*GetMemberQuotaApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Member/getQuota"

	apiResp, err := core.Request(ctx, apiReq, member.option, reqOption)
	if err != nil {
		return nil, err
	}
	getMemberQuotaApiResp := GetMemberQuotaApiResp{
		ApiResp: apiResp,
	}
	getMemberQuotaApiReply := GetMemberQuotaApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := member.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if member.option.EnableEncryption && nil != member.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch member.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", member.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(member.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					member.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getMemberQuotaApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getMemberQuotaApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getMemberQuotaApiReply); err != nil {
				return nil, err
			}
		}
		getMemberQuotaApiResp.GetMemberQuotaApiReply = &getMemberQuotaApiReply
	}
	return &getMemberQuotaApiResp, nil
}

// ListMember 员工列表（批量查询）
func (member *member) ListMember(ctx context.Context, req *ListMemberApiReq, reqOption *core.ReqOption) (*ListMemberApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Member/get"

	apiResp, err := core.Request(ctx, apiReq, member.option, reqOption)
	if err != nil {
		return nil, err
	}
	listMemberApiResp := ListMemberApiResp{
		ApiResp: apiResp,
	}
	listMemberApiReply := ListMemberApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := member.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if member.option.EnableEncryption && nil != member.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch member.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", member.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(member.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					member.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listMemberApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listMemberApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listMemberApiReply); err != nil {
				return nil, err
			}
		}
		listMemberApiResp.ListMemberApiReply = &listMemberApiReply
	}
	return &listMemberApiResp, nil
}

// UpdateMember 用户修改
func (member *member) UpdateMember(ctx context.Context, req *UpdateMemberApiReq, reqOption *core.ReqOption) (*UpdateMemberApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Member/edit"

	apiResp, err := core.Request(ctx, apiReq, member.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateMemberApiResp := UpdateMemberApiResp{
		ApiResp: apiResp,
	}
	updateMemberApiReply := UpdateMemberApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := member.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if member.option.EnableEncryption && nil != member.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch member.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", member.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(member.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					member.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
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
