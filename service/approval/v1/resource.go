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
	Approval *approval // 审批(行前申请单、行后授权)
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Approval: &approval{option: option},
	}
}

type approval struct {
	option *core.Option
}

// ApprovalPass 外部审批处理接口
func (approval *approval) ApprovalPass(ctx context.Context, req *ApprovalPassApiReq, reqOption *core.ReqOption) (*ApprovalPassApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/pass"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	approvalPassApiResp := ApprovalPassApiResp{
		ApiResp: apiResp,
	}
	approvalPassApiReply := ApprovalPassApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &approvalPassApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &approvalPassApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &approvalPassApiReply); err != nil {
				return nil, err
			}
		}
		approvalPassApiResp.ApprovalPassApiReply = &approvalPassApiReply
	}
	return &approvalPassApiResp, nil
}

// CancelApproval 取消申请单
func (approval *approval) CancelApproval(ctx context.Context, req *CancelApprovalApiReq, reqOption *core.ReqOption) (*CancelApprovalApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/cancel"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	cancelApprovalApiResp := CancelApprovalApiResp{
		ApiResp: apiResp,
	}
	cancelApprovalApiReply := CancelApprovalApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &cancelApprovalApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &cancelApprovalApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &cancelApprovalApiReply); err != nil {
				return nil, err
			}
		}
		cancelApprovalApiResp.CancelApprovalApiReply = &cancelApprovalApiReply
	}
	return &cancelApprovalApiResp, nil
}

// CreateTravelApproval 创建[差旅]申请单(行前申请单)
func (approval *approval) CreateTravelApproval(ctx context.Context, req *CreateTravelApprovalApiReq, reqOption *core.ReqOption) (*CreateTravelApprovalApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/create"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	createTravelApprovalApiResp := CreateTravelApprovalApiResp{
		ApiResp: apiResp,
	}
	createApprovalApiReply := CreateApprovalApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createApprovalApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createApprovalApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createApprovalApiReply); err != nil {
				return nil, err
			}
		}
		createTravelApprovalApiResp.CreateApprovalApiReply = &createApprovalApiReply
	}
	return &createTravelApprovalApiResp, nil
}

// CreateBusinessByTimesApproval 创建[行前审批-按日期]申请单(行前申请单)
func (approval *approval) CreateBusinessByTimesApproval(ctx context.Context, req *CreateBusinessByTimesApprovalApiReq, reqOption *core.ReqOption) (*CreateBusinessByTimesApprovalApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/create"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	createBusinessByTimesApprovalApiResp := CreateBusinessByTimesApprovalApiResp{
		ApiResp: apiResp,
	}
	createApprovalApiReply := CreateApprovalApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createApprovalApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createApprovalApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createApprovalApiReply); err != nil {
				return nil, err
			}
		}
		createBusinessByTimesApprovalApiResp.CreateApprovalApiReply = &createApprovalApiReply
	}
	return &createBusinessByTimesApprovalApiResp, nil
}

// CreateBusinessByDateApproval 创建[行前审批-按次数]申请单(行前申请单)
func (approval *approval) CreateBusinessByDateApproval(ctx context.Context, req *CreateBusinessByDateApprovalApiReq, reqOption *core.ReqOption) (*CreateBusinessByDateApprovalApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/create"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	createBusinessByDateApprovalApiResp := CreateBusinessByDateApprovalApiResp{
		ApiResp: apiResp,
	}
	createApprovalApiReply := CreateApprovalApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &createApprovalApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &createApprovalApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &createApprovalApiReply); err != nil {
				return nil, err
			}
		}
		createBusinessByDateApprovalApiResp.CreateApprovalApiReply = &createApprovalApiReply
	}
	return &createBusinessByDateApprovalApiResp, nil
}

// GetApprovalDetail 申请单详情接口
func (approval *approval) GetApprovalDetail(ctx context.Context, req *GetApprovalDetailApiReq, reqOption *core.ReqOption) (*GetApprovalDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/open-apis/v1/approval/detail"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	getApprovalDetailApiResp := GetApprovalDetailApiResp{
		ApiResp: apiResp,
	}
	getApprovalDetailApiReply := GetApprovalDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getApprovalDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getApprovalDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getApprovalDetailApiReply); err != nil {
				return nil, err
			}
		}
		getApprovalDetailApiResp.GetApprovalDetailApiReply = &getApprovalDetailApiReply
	}
	return &getApprovalDetailApiResp, nil
}

// ListApprovalOrder 审批单查询订单（用车）
func (approval *approval) ListApprovalOrder(ctx context.Context, req *ListApprovalOrderApiReq, reqOption *core.ReqOption) (*ListApprovalOrderApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Approval/getOrder"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	listApprovalOrderApiResp := ListApprovalOrderApiResp{
		ApiResp: apiResp,
	}
	listApprovalOrderApiReply := ListApprovalOrderApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listApprovalOrderApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listApprovalOrderApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listApprovalOrderApiReply); err != nil {
				return nil, err
			}
		}
		listApprovalOrderApiResp.ListApprovalOrderApiReply = &listApprovalOrderApiReply
	}
	return &listApprovalOrderApiResp, nil
}

// UpdateApproval 修改申请单(行前申请单)
func (approval *approval) UpdateApproval(ctx context.Context, req *UpdateApprovalApiReq, reqOption *core.ReqOption) (*UpdateApprovalApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/update"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateApprovalApiResp := UpdateApprovalApiResp{
		ApiResp: apiResp,
	}
	updateApprovalApiReply := UpdateApprovalApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateApprovalApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateApprovalApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateApprovalApiReply); err != nil {
				return nil, err
			}
		}
		updateApprovalApiResp.UpdateApprovalApiReply = &updateApprovalApiReply
	}
	return &updateApprovalApiResp, nil
}

// UpdateBusinessByTimesApproval 修改[用车按次数]申请单(行前申请单)
func (approval *approval) UpdateBusinessByTimesApproval(ctx context.Context, req *UpdateBusinessByTimesApprovalApiReq, reqOption *core.ReqOption) (*UpdateBusinessByTimesApprovalApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/update"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateBusinessByTimesApprovalApiResp := UpdateBusinessByTimesApprovalApiResp{
		ApiResp: apiResp,
	}
	updateApprovalApiReply := UpdateApprovalApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateApprovalApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateApprovalApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateApprovalApiReply); err != nil {
				return nil, err
			}
		}
		updateBusinessByTimesApprovalApiResp.UpdateApprovalApiReply = &updateApprovalApiReply
	}
	return &updateBusinessByTimesApprovalApiResp, nil
}

// UpdateBusinessByDateApproval 修改[用车按日期]申请单(行前申请单)
func (approval *approval) UpdateBusinessByDateApproval(ctx context.Context, req *UpdateBusinessByDateApprovalApiReq, reqOption *core.ReqOption) (*UpdateBusinessByDateApprovalApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Approval/update"

	apiResp, err := core.Request(ctx, apiReq, approval.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateBusinessByDateApprovalApiResp := UpdateBusinessByDateApprovalApiResp{
		ApiResp: apiResp,
	}
	updateApprovalApiReply := UpdateApprovalApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := approval.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if approval.option.EnableEncryption && nil != approval.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch approval.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", approval.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(approval.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					approval.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateApprovalApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateApprovalApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateApprovalApiReply); err != nil {
				return nil, err
			}
		}
		updateBusinessByDateApprovalApiResp.UpdateApprovalApiReply = &updateApprovalApiReply
	}
	return &updateBusinessByDateApprovalApiResp, nil
}
