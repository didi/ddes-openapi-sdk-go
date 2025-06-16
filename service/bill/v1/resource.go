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
	Bill *bill // 财税信息(账单管理、账单查询)
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Bill: &bill{option: option},
	}
}

type bill struct {
	option *core.Option
}

// BillConfirm 商旅、网约车账单确认
func (bill *bill) BillConfirm(ctx context.Context, req *BillConfirmApiReq, reqOption *core.ReqOption) (*BillConfirmApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Bill/confirm"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	billConfirmApiResp := BillConfirmApiResp{
		ApiResp: apiResp,
	}
	billConfirmApiReply := BillConfirmApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &billConfirmApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &billConfirmApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &billConfirmApiReply); err != nil {
				return nil, err
			}
		}
		billConfirmApiResp.BillConfirmApiReply = &billConfirmApiReply
	}
	return &billConfirmApiResp, nil
}

// GetAdjustBillDataResult 调账查询结果
func (bill *bill) GetAdjustBillDataResult(ctx context.Context, req *GetAdjustBillDataResultApiReq, reqOption *core.ReqOption) (*GetAdjustBillDataResultApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Bill/queryAdjustBillDataResult"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getAdjustBillDataResultApiResp := GetAdjustBillDataResultApiResp{
		ApiResp: apiResp,
	}
	getAdjustBillDataResultApiReply := GetAdjustBillDataResultApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getAdjustBillDataResultApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getAdjustBillDataResultApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getAdjustBillDataResultApiReply); err != nil {
				return nil, err
			}
		}
		getAdjustBillDataResultApiResp.GetAdjustBillDataResultApiReply = &getAdjustBillDataResultApiReply
	}
	return &getAdjustBillDataResultApiResp, nil
}

// GetBillDetailOfDaiJia 已出账单(按业务类型拆分 - 代驾)
func (bill *bill) GetBillDetailOfDaiJia(ctx context.Context, req *GetBillDetailOfDaiJiaApiReq, reqOption *core.ReqOption) (*GetBillDetailOfDaiJiaApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfDaiJiaApiResp := GetBillDetailOfDaiJiaApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfDaiJiaApiReply := GetBillDetailOfDaiJiaApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfDaiJiaApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfDaiJiaApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfDaiJiaApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfDaiJiaApiResp.GetBillDetailOfDaiJiaApiReply = &getBillDetailOfDaiJiaApiReply
	}
	return &getBillDetailOfDaiJiaApiResp, nil
}

// GetBillDetailOfDomesticFlight 已出账单(按业务类型拆分 - 国内机票)
func (bill *bill) GetBillDetailOfDomesticFlight(ctx context.Context, req *GetBillDetailOfDomesticFlightApiReq, reqOption *core.ReqOption) (*GetBillDetailOfDomesticFlightApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfDomesticFlightApiResp := GetBillDetailOfDomesticFlightApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfDomesticFlightApiReply := GetBillDetailOfDomesticFlightApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfDomesticFlightApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfDomesticFlightApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfDomesticFlightApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfDomesticFlightApiResp.GetBillDetailOfDomesticFlightApiReply = &getBillDetailOfDomesticFlightApiReply
	}
	return &getBillDetailOfDomesticFlightApiResp, nil
}

// GetBillDetailOfDomesticHotel 已出账单(按业务类型拆分 - 国内酒店)
func (bill *bill) GetBillDetailOfDomesticHotel(ctx context.Context, req *GetBillDetailOfDomesticHotelApiReq, reqOption *core.ReqOption) (*GetBillDetailOfDomesticHotelApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfDomesticHotelApiResp := GetBillDetailOfDomesticHotelApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfDomesticHotelApiReply := GetBillDetailOfDomesticHotelApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfDomesticHotelApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfDomesticHotelApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfDomesticHotelApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfDomesticHotelApiResp.GetBillDetailOfDomesticHotelApiReply = &getBillDetailOfDomesticHotelApiReply
	}
	return &getBillDetailOfDomesticHotelApiResp, nil
}

// GetBillDetailOfInterFlight 已出账单(按业务类型拆分 - 国际机票)
func (bill *bill) GetBillDetailOfInterFlight(ctx context.Context, req *GetBillDetailOfInterFlightApiReq, reqOption *core.ReqOption) (*GetBillDetailOfInterFlightApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfInterFlightApiResp := GetBillDetailOfInterFlightApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfInterFlightApiReply := GetBillDetailOfInterFlightApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfInterFlightApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfInterFlightApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfInterFlightApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfInterFlightApiResp.GetBillDetailOfInterFlightApiReply = &getBillDetailOfInterFlightApiReply
	}
	return &getBillDetailOfInterFlightApiResp, nil
}

// GetBillDetailOfInterHotel 已出账单(按业务类型拆分 - 海外酒店)
func (bill *bill) GetBillDetailOfInterHotel(ctx context.Context, req *GetBillDetailOfInterHotelApiReq, reqOption *core.ReqOption) (*GetBillDetailOfInterHotelApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfInterHotelApiResp := GetBillDetailOfInterHotelApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfInterHotelApiReply := GetBillDetailOfInterHotelApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfInterHotelApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfInterHotelApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfInterHotelApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfInterHotelApiResp.GetBillDetailOfInterHotelApiReply = &getBillDetailOfInterHotelApiReply
	}
	return &getBillDetailOfInterHotelApiResp, nil
}

// GetBillDetailOfManualOrder 已出账单(按业务类型拆分 - 增值手工单)
func (bill *bill) GetBillDetailOfManualOrder(ctx context.Context, req *GetBillDetailOfManualOrderApiReq, reqOption *core.ReqOption) (*GetBillDetailOfManualOrderApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfManualOrderApiResp := GetBillDetailOfManualOrderApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfManualOrderApiReply := GetBillDetailOfManualOrderApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfManualOrderApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfManualOrderApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfManualOrderApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfManualOrderApiResp.GetBillDetailOfManualOrderApiReply = &getBillDetailOfManualOrderApiReply
	}
	return &getBillDetailOfManualOrderApiResp, nil
}

// GetBillDetailOfTaxi 已出账单(按业务类型拆分 - 出租车)
func (bill *bill) GetBillDetailOfTaxi(ctx context.Context, req *GetBillDetailOfTaxiApiReq, reqOption *core.ReqOption) (*GetBillDetailOfTaxiApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfTaxiApiResp := GetBillDetailOfTaxiApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfTaxiApiReply := GetBillDetailOfTaxiApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfTaxiApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfTaxiApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfTaxiApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfTaxiApiResp.GetBillDetailOfTaxiApiReply = &getBillDetailOfTaxiApiReply
	}
	return &getBillDetailOfTaxiApiResp, nil
}

// GetBillDetailOfTrainTicket 已出账单(按业务类型拆分 - 火车票)
func (bill *bill) GetBillDetailOfTrainTicket(ctx context.Context, req *GetBillDetailOfTrainTicketApiReq, reqOption *core.ReqOption) (*GetBillDetailOfTrainTicketApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfTrainTicketApiResp := GetBillDetailOfTrainTicketApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfTrainTicketApiReply := GetBillDetailOfTrainTicketApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfTrainTicketApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfTrainTicketApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfTrainTicketApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfTrainTicketApiResp.GetBillDetailOfTrainTicketApiReply = &getBillDetailOfTrainTicketApiReply
	}
	return &getBillDetailOfTrainTicketApiResp, nil
}

// GetBillDetailOfWangYC 已出账单(按业务类型拆分 - 网约车)
func (bill *bill) GetBillDetailOfWangYC(ctx context.Context, req *GetBillDetailOfWangYCApiReq, reqOption *core.ReqOption) (*GetBillDetailOfWangYCApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/detail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillDetailOfWangYCApiResp := GetBillDetailOfWangYCApiResp{
		ApiResp: apiResp,
	}
	getBillDetailOfWangYCApiReply := GetBillDetailOfWangYCApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillDetailOfWangYCApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfWangYCApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillDetailOfWangYCApiReply); err != nil {
				return nil, err
			}
		}
		getBillDetailOfWangYCApiResp.GetBillDetailOfWangYCApiReply = &getBillDetailOfWangYCApiReply
	}
	return &getBillDetailOfWangYCApiResp, nil
}

// GetBillStructure 网约车、商旅账单树查询
func (bill *bill) GetBillStructure(ctx context.Context, req *GetBillStructureApiReq, reqOption *core.ReqOption) (*GetBillStructureApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getBillStructure"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillStructureApiResp := GetBillStructureApiResp{
		ApiResp: apiResp,
	}
	getBillStructureApiReply := GetBillStructureApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillStructureApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillStructureApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillStructureApiReply); err != nil {
				return nil, err
			}
		}
		getBillStructureApiResp.GetBillStructureApiReply = &getBillStructureApiReply
	}
	return &getBillStructureApiResp, nil
}

// GetBillSummary 账单汇总查询-商旅、网约车、出租车
func (bill *bill) GetBillSummary(ctx context.Context, req *GetBillSummaryApiReq, reqOption *core.ReqOption) (*GetBillSummaryApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/summary"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getBillSummaryApiResp := GetBillSummaryApiResp{
		ApiResp: apiResp,
	}
	getBillSummaryApiReply := GetBillSummaryApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getBillSummaryApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getBillSummaryApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getBillSummaryApiReply); err != nil {
				return nil, err
			}
		}
		getBillSummaryApiResp.GetBillSummaryApiReply = &getBillSummaryApiReply
	}
	return &getBillSummaryApiResp, nil
}

// GetNotGenBillDetailOfDaiJia 未出账单(按业务类型拆分 - 代驾)
func (bill *bill) GetNotGenBillDetailOfDaiJia(ctx context.Context, req *GetNotGenBillDetailOfDaiJiaApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfDaiJiaApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfDaiJiaApiResp := GetNotGenBillDetailOfDaiJiaApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfDaiJiaApiReply := GetNotGenBillDetailOfDaiJiaApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfDaiJiaApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfDaiJiaApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfDaiJiaApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfDaiJiaApiResp.GetNotGenBillDetailOfDaiJiaApiReply = &getNotGenBillDetailOfDaiJiaApiReply
	}
	return &getNotGenBillDetailOfDaiJiaApiResp, nil
}

// GetNotGenBillDetailOfFlight 未出账单(按业务类型拆分 - 国内机票)
func (bill *bill) GetNotGenBillDetailOfFlight(ctx context.Context, req *GetNotGenBillDetailOfFlightApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfFlightApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfFlightApiResp := GetNotGenBillDetailOfFlightApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfFlightApiReply := GetNotGenBillDetailOfFlightApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfFlightApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfFlightApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfFlightApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfFlightApiResp.GetNotGenBillDetailOfFlightApiReply = &getNotGenBillDetailOfFlightApiReply
	}
	return &getNotGenBillDetailOfFlightApiResp, nil
}

// GetNotGenBillDetailOfHotel 未出账单(按业务类型拆分 - 国内酒店)
func (bill *bill) GetNotGenBillDetailOfHotel(ctx context.Context, req *GetNotGenBillDetailOfHotelApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfHotelApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfHotelApiResp := GetNotGenBillDetailOfHotelApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfHotelApiReply := GetNotGenBillDetailOfHotelApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfHotelApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfHotelApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfHotelApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfHotelApiResp.GetNotGenBillDetailOfHotelApiReply = &getNotGenBillDetailOfHotelApiReply
	}
	return &getNotGenBillDetailOfHotelApiResp, nil
}

// GetNotGenBillDetailOfInterFlight 未出账单(按业务类型拆分 - 国际机票)
func (bill *bill) GetNotGenBillDetailOfInterFlight(ctx context.Context, req *GetNotGenBillDetailOfInterFlightApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfInterFlightApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfInterFlightApiResp := GetNotGenBillDetailOfInterFlightApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfInterFlightApiReply := GetNotGenBillDetailOfInterFlightApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfInterFlightApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfInterFlightApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfInterFlightApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfInterFlightApiResp.GetNotGenBillDetailOfInterFlightApiReply = &getNotGenBillDetailOfInterFlightApiReply
	}
	return &getNotGenBillDetailOfInterFlightApiResp, nil
}

// GetNotGenBillDetailOfInterHotel 未出账单(按业务类型拆分 - 海外酒店)
func (bill *bill) GetNotGenBillDetailOfInterHotel(ctx context.Context, req *GetNotGenBillDetailOfInterHotelApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfInterHotelApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfInterHotelApiResp := GetNotGenBillDetailOfInterHotelApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfInterHotelApiReply := GetNotGenBillDetailOfInterHotelApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfInterHotelApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfInterHotelApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfInterHotelApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfInterHotelApiResp.GetNotGenBillDetailOfInterHotelApiReply = &getNotGenBillDetailOfInterHotelApiReply
	}
	return &getNotGenBillDetailOfInterHotelApiResp, nil
}

// GetNotGenBillDetailOfManualOrder 未出账单(按业务类型拆分 - 增值手工单)
func (bill *bill) GetNotGenBillDetailOfManualOrder(ctx context.Context, req *GetNotGenBillDetailOfManualOrderApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfManualOrderApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfManualOrderApiResp := GetNotGenBillDetailOfManualOrderApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfManualOrderApiReply := GetNotGenBillDetailOfManualOrderApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfManualOrderApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfManualOrderApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfManualOrderApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfManualOrderApiResp.GetNotGenBillDetailOfManualOrderApiReply = &getNotGenBillDetailOfManualOrderApiReply
	}
	return &getNotGenBillDetailOfManualOrderApiResp, nil
}

// GetNotGenBillDetailOfTaxi 未出账单(按业务类型拆分 - 出租车)
func (bill *bill) GetNotGenBillDetailOfTaxi(ctx context.Context, req *GetNotGenBillDetailOfTaxiApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfTaxiApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfTaxiApiResp := GetNotGenBillDetailOfTaxiApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfTaxiApiReply := GetNotGenBillDetailOfTaxiApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfTaxiApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfTaxiApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfTaxiApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfTaxiApiResp.GetNotGenBillDetailOfTaxiApiReply = &getNotGenBillDetailOfTaxiApiReply
	}
	return &getNotGenBillDetailOfTaxiApiResp, nil
}

// GetNotGenBillDetailOfTrain 未出账单(按业务类型拆分 - 火车票)
func (bill *bill) GetNotGenBillDetailOfTrain(ctx context.Context, req *GetNotGenBillDetailOfTrainApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfTrainApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfTrainApiResp := GetNotGenBillDetailOfTrainApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfTrainApiReply := GetNotGenBillDetailOfTrainApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfTrainApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfTrainApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfTrainApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfTrainApiResp.GetNotGenBillDetailOfTrainApiReply = &getNotGenBillDetailOfTrainApiReply
	}
	return &getNotGenBillDetailOfTrainApiResp, nil
}

// GetNotGenBillDetailOfWangYC 未出账单(按业务类型拆分 - 网约车)
func (bill *bill) GetNotGenBillDetailOfWangYC(ctx context.Context, req *GetNotGenBillDetailOfWangYCApiReq, reqOption *core.ReqOption) (*GetNotGenBillDetailOfWangYCApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/getNotGeneratedBillDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getNotGenBillDetailOfWangYCApiResp := GetNotGenBillDetailOfWangYCApiResp{
		ApiResp: apiResp,
	}
	getNotGenBillDetailOfWangYCApiReply := GetNotGenBillDetailOfWangYCApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getNotGenBillDetailOfWangYCApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfWangYCApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getNotGenBillDetailOfWangYCApiReply); err != nil {
				return nil, err
			}
		}
		getNotGenBillDetailOfWangYCApiResp.GetNotGenBillDetailOfWangYCApiReply = &getNotGenBillDetailOfWangYCApiReply
	}
	return &getNotGenBillDetailOfWangYCApiResp, nil
}

// GetTransactionBillDetail 网约车、出租车交易明细 ~ 网约车
func (bill *bill) GetTransactionBillDetail(ctx context.Context, req *GetTransactionBillDetailApiReq, reqOption *core.ReqOption) (*GetTransactionBillDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/transactionDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getTransactionBillDetailApiResp := GetTransactionBillDetailApiResp{
		ApiResp: apiResp,
	}
	getTransactionBillDetailApiReply := GetTransactionBillDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getTransactionBillDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getTransactionBillDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getTransactionBillDetailApiReply); err != nil {
				return nil, err
			}
		}
		getTransactionBillDetailApiResp.GetTransactionBillDetailApiReply = &getTransactionBillDetailApiReply
	}
	return &getTransactionBillDetailApiResp, nil
}

// GetTransactionBillDetailOfTaxi 网约车、出租车交易明细 ~ 出租车
func (bill *bill) GetTransactionBillDetailOfTaxi(ctx context.Context, req *GetTransactionBillDetailOfTaxiApiReq, reqOption *core.ReqOption) (*GetTransactionBillDetailOfTaxiApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/transactionDetail"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	getTransactionBillDetailOfTaxiApiResp := GetTransactionBillDetailOfTaxiApiResp{
		ApiResp: apiResp,
	}
	getTransactionBillDetailOfTaxiApiReply := GetTransactionBillDetailOfTaxiApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getTransactionBillDetailOfTaxiApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getTransactionBillDetailOfTaxiApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getTransactionBillDetailOfTaxiApiReply); err != nil {
				return nil, err
			}
		}
		getTransactionBillDetailOfTaxiApiResp.GetTransactionBillDetailOfTaxiApiReply = &getTransactionBillDetailOfTaxiApiReply
	}
	return &getTransactionBillDetailOfTaxiApiResp, nil
}

// ListBill 账单列表接口
func (bill *bill) ListBill(ctx context.Context, req *ListBillApiReq, reqOption *core.ReqOption) (*ListBillApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Bill/get"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	listBillApiResp := ListBillApiResp{
		ApiResp: apiResp,
	}
	listBillApiReply := ListBillApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listBillApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listBillApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listBillApiReply); err != nil {
				return nil, err
			}
		}
		listBillApiResp.ListBillApiReply = &listBillApiReply
	}
	return &listBillApiResp, nil
}

// UpdateAdjustBillData 账单调整提交接口
func (bill *bill) UpdateAdjustBillData(ctx context.Context, req *UpdateAdjustBillDataApiReq, reqOption *core.ReqOption) (*UpdateAdjustBillDataApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/Bill/adjustBillData"

	apiResp, err := core.Request(ctx, apiReq, bill.option, reqOption)
	if err != nil {
		return nil, err
	}
	updateAdjustBillDataApiResp := UpdateAdjustBillDataApiResp{
		ApiResp: apiResp,
	}
	updateAdjustBillDataApiReply := UpdateAdjustBillDataApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := bill.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if bill.option.EnableEncryption && nil != bill.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch bill.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", bill.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(bill.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					bill.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &updateAdjustBillDataApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &updateAdjustBillDataApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &updateAdjustBillDataApiReply); err != nil {
				return nil, err
			}
		}
		updateAdjustBillDataApiResp.UpdateAdjustBillDataApiReply = &updateAdjustBillDataApiReply
	}
	return &updateAdjustBillDataApiResp, nil
}
