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
	Rank *rank // 职级
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Rank: &rank{option: option},
	}
}

type rank struct {
	option *core.Option
}

// CreateRank 职级新增
func (rank *rank) CreateRank(ctx context.Context, req *CreateRankApiReq, reqOption *core.ReqOption) (*CreateRankApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/rank/create"

	apiResp, err := core.Request(ctx, apiReq, rank.option, reqOption)
	if err != nil {
		return nil, err
	}
	createRankApiResp := CreateRankApiResp{
		ApiResp: apiResp,
	}
	createRankApiReply := CreateRankApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := rank.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if rank.option.EnableEncryption && nil != rank.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch rank.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", rank.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(rank.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					rank.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createRankApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createRankApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createRankApiReply); err != nil {
				return nil, err
			}
		}
		createRankApiResp.CreateRankApiReply = &createRankApiReply
	}
	return &createRankApiResp, nil
}

// DelRank 职级删除
func (rank *rank) DelRank(ctx context.Context, req *DelRankApiReq, reqOption *core.ReqOption) (*DelRankApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/rank/del"

	apiResp, err := core.Request(ctx, apiReq, rank.option, reqOption)
	if err != nil {
		return nil, err
	}
	delRankApiResp := DelRankApiResp{
		ApiResp: apiResp,
	}
	delRankApiReply := DelRankApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := rank.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if rank.option.EnableEncryption && nil != rank.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch rank.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", rank.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(rank.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					rank.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &delRankApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &delRankApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &delRankApiReply); err != nil {
				return nil, err
			}
		}
		delRankApiResp.DelRankApiReply = &delRankApiReply
	}
	return &delRankApiResp, nil
}

// ListRank 职级查询
func (rank *rank) ListRank(ctx context.Context, req *ListRankApiReq, reqOption *core.ReqOption) (*ListRankApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Rank/getRanks"

	apiResp, err := core.Request(ctx, apiReq, rank.option, reqOption)
	if err != nil {
		return nil, err
	}
	listRankApiResp := ListRankApiResp{
		ApiResp: apiResp,
	}
	listRankApiReply := ListRankApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := rank.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if rank.option.EnableEncryption && nil != rank.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch rank.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", rank.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(rank.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					rank.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listRankApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listRankApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listRankApiReply); err != nil {
				return nil, err
			}
		}
		listRankApiResp.ListRankApiReply = &listRankApiReply
	}
	return &listRankApiResp, nil
}

// UpdateRank 职级修改
func (rank *rank) UpdateRank(ctx context.Context, req *UpdateRankApiReq, reqOption *core.ReqOption) (*UpdateRankApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/rank/update"

	apiResp, err := core.Request(ctx, apiReq, rank.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateRankApiResp := UpdateRankApiResp{
		ApiResp: apiResp,
	}
	updateRankApiReply := UpdateRankApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := rank.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if rank.option.EnableEncryption && nil != rank.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch rank.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", rank.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(rank.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					rank.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateRankApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateRankApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateRankApiReply); err != nil {
				return nil, err
			}
		}
		updateRankApiResp.UpdateRankApiReply = &updateRankApiReply
	}
	return &updateRankApiResp, nil
}
