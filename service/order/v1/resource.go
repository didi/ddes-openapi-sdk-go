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
	Order *order // 订单
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		Order: &order{option: option},
	}
}

type order struct {
	option *core.Option
}

// GetCarOrderDetail 用车订单详情查询
func (order *order) GetCarOrderDetail(ctx context.Context, req *GetCarOrderDetailApiReq, reqOption *core.ReqOption) (*GetCarOrderDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Order/detail"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	getCarOrderDetailApiResp := GetCarOrderDetailApiResp{
		ApiResp: apiResp,
	}
	getCarOrderDetailApiReply := GetCarOrderDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getCarOrderDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getCarOrderDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getCarOrderDetailApiReply); err != nil {
				return nil, err
			}
		}
		getCarOrderDetailApiResp.GetCarOrderDetailApiReply = &getCarOrderDetailApiReply
	}
	return &getCarOrderDetailApiResp, nil
}

// GetFlightEstimatePrice 机票预估价接口
func (order *order) GetFlightEstimatePrice(ctx context.Context, req *GetFlightEstimatePriceApiReq, reqOption *core.ReqOption) (*GetFlightEstimatePriceApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/api-gateway/g/flight/info/estimatePrice"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	getFlightEstimatePriceApiResp := GetFlightEstimatePriceApiResp{
		ApiResp: apiResp,
	}
	getFlightEstimatePriceApiReply := GetFlightEstimatePriceApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getFlightEstimatePriceApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getFlightEstimatePriceApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getFlightEstimatePriceApiReply); err != nil {
				return nil, err
			}
		}
		getFlightEstimatePriceApiResp.GetFlightEstimatePriceApiReply = &getFlightEstimatePriceApiReply
	}
	return &getFlightEstimatePriceApiResp, nil
}

// GetFlightOrderDetail 机票订单详情查询
func (order *order) GetFlightOrderDetail(ctx context.Context, req *GetFlightOrderDetailApiReq, reqOption *core.ReqOption) (*GetFlightOrderDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/api-gateway/g/flight/orderDetail"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	getFlightOrderDetailApiResp := GetFlightOrderDetailApiResp{
		ApiResp: apiResp,
	}
	getFlightOrderDetailApiReply := GetFlightOrderDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getFlightOrderDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getFlightOrderDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getFlightOrderDetailApiReply); err != nil {
				return nil, err
			}
		}
		getFlightOrderDetailApiResp.GetFlightOrderDetailApiReply = &getFlightOrderDetailApiReply
	}
	return &getFlightOrderDetailApiResp, nil
}

// GetHotelOrderDetail 酒店订单详情查询
func (order *order) GetHotelOrderDetail(ctx context.Context, req *GetHotelOrderDetailApiReq, reqOption *core.ReqOption) (*GetHotelOrderDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/api-gateway/g/hotel/orderDetail"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	getHotelOrderDetailApiResp := GetHotelOrderDetailApiResp{
		ApiResp: apiResp,
	}
	getHotelOrderDetailApiReply := GetHotelOrderDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getHotelOrderDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getHotelOrderDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getHotelOrderDetailApiReply); err != nil {
				return nil, err
			}
		}
		getHotelOrderDetailApiResp.GetHotelOrderDetailApiReply = &getHotelOrderDetailApiReply
	}
	return &getHotelOrderDetailApiResp, nil
}

// GetOrder 用车列表接口
func (order *order) GetOrder(ctx context.Context, req *GetOrderApiReq, reqOption *core.ReqOption) (*GetOrderApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/Order/get"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	getOrderApiResp := GetOrderApiResp{
		ApiResp: apiResp,
	}
	getOrderApiReply := GetOrderApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getOrderApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getOrderApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getOrderApiReply); err != nil {
				return nil, err
			}
		}
		getOrderApiResp.GetOrderApiReply = &getOrderApiReply
	}
	return &getOrderApiResp, nil
}

// GetTrainOrderDetail 火车票订单详情查询
func (order *order) GetTrainOrderDetail(ctx context.Context, req *GetTrainOrderDetailApiReq, reqOption *core.ReqOption) (*GetTrainOrderDetailApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/api-gateway/g/train/orderDetail"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	getTrainOrderDetailApiResp := GetTrainOrderDetailApiResp{
		ApiResp: apiResp,
	}
	getTrainOrderDetailApiReply := GetTrainOrderDetailApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &getTrainOrderDetailApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &getTrainOrderDetailApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &getTrainOrderDetailApiReply); err != nil {
				return nil, err
			}
		}
		getTrainOrderDetailApiResp.GetTrainOrderDetailApiReply = &getTrainOrderDetailApiReply
	}
	return &getTrainOrderDetailApiResp, nil
}

// ListOrder 订单号列表查询
func (order *order) ListOrder(ctx context.Context, req *ListOrderApiReq, reqOption *core.ReqOption) (*ListOrderApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/order/list"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	listOrderApiResp := ListOrderApiResp{
		ApiResp: apiResp,
	}
	listOrderApiReply := ListOrderApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listOrderApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listOrderApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listOrderApiReply); err != nil {
				return nil, err
			}
		}
		listOrderApiResp.ListOrderApiReply = &listOrderApiReply
	}
	return &listOrderApiResp, nil
}

// ListTrainLeftTicket 火车票直达列表接口
func (order *order) ListTrainLeftTicket(ctx context.Context, req *ListTrainLeftTicketApiReq, reqOption *core.ReqOption) (*ListTrainLeftTicketApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/api-gateway/train/queryLeftTicket"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	listTrainLeftTicketApiResp := ListTrainLeftTicketApiResp{
		ApiResp: apiResp,
	}
	listTrainLeftTicketApiReply := ListTrainLeftTicketApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listTrainLeftTicketApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listTrainLeftTicketApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listTrainLeftTicketApiReply); err != nil {
				return nil, err
			}
		}
		listTrainLeftTicketApiResp.ListTrainLeftTicketApiReply = &listTrainLeftTicketApiReply
	}
	return &listTrainLeftTicketApiResp, nil
}

// ListTransferTrainTicket 火车票中转车次列表
func (order *order) ListTransferTrainTicket(ctx context.Context, req *ListTransferTrainTicketApiReq, reqOption *core.ReqOption) (*ListTransferTrainTicketApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/api-gateway/g/train/transfer/queryLeftTicket"

	apiResp, err := core.Request(ctx, apiReq, order.option, reqOption)
	if err != nil {
		return nil, err
	}
	listTransferTrainTicketApiResp := ListTransferTrainTicketApiResp{
		ApiResp: apiResp,
	}
	listTransferTrainTicketApiReply := ListTransferTrainTicketApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := order.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if order.option.EnableEncryption && nil != order.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch order.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", order.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(order.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					order.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listTransferTrainTicketApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listTransferTrainTicketApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listTransferTrainTicketApiReply); err != nil {
				return nil, err
			}
		}
		listTransferTrainTicketApiResp.ListTransferTrainTicketApiReply = &listTransferTrainTicketApiReply
	}
	return &listTransferTrainTicketApiResp, nil
}
