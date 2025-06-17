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
	City *city // 城市
}

func NewV1(option *core.Option) *V1 {
	return &V1{
		City: &city{option: option},
	}
}

type city struct {
	option *core.Option
}

// ListAirportCity 机票城市
// Deprecated
func (city *city) ListAirportCity(ctx context.Context, req *ListAirportCityApiReq, reqOption *core.ReqOption) (*ListAirportCityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/DemeterAres/AirportCity/index"

	apiResp, err := core.Request(ctx, apiReq, city.option, reqOption)
	if err != nil {
		return nil, err
	}
	listAirportCityApiResp := ListAirportCityApiResp{
		ApiResp: apiResp,
	}
	listAirportCityApiReply := ListAirportCityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := city.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if city.option.EnableEncryption && nil != city.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch city.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", city.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(city.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					city.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listAirportCityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listAirportCityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listAirportCityApiReply); err != nil {
				return nil, err
			}
		}
		listAirportCityApiResp.ListAirportCityApiReply = &listAirportCityApiReply
	}
	return &listAirportCityApiResp, nil
}

// ListCarCity 用车城市
// Deprecated
func (city *city) ListCarCity(ctx context.Context, req *ListCarCityApiReq, reqOption *core.ReqOption) (*ListCarCityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/river/City/get"

	apiResp, err := core.Request(ctx, apiReq, city.option, reqOption)
	if err != nil {
		return nil, err
	}
	listCarCityApiResp := ListCarCityApiResp{
		ApiResp: apiResp,
	}
	listCarCityApiReply := ListCarCityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := city.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if city.option.EnableEncryption && nil != city.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch city.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", city.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(city.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					city.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listCarCityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listCarCityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listCarCityApiReply); err != nil {
				return nil, err
			}
		}
		listCarCityApiResp.ListCarCityApiReply = &listCarCityApiReply
	}
	return &listCarCityApiResp, nil
}

// ListCity 全量开城城市列表查询
func (city *city) ListCity(ctx context.Context, req *ListCityApiReq, reqOption *core.ReqOption) (*ListCityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/open-apis/v1/city/list"

	apiResp, err := core.Request(ctx, apiReq, city.option, reqOption)
	if err != nil {
		return nil, err
	}
	listCityApiResp := ListCityApiResp{
		ApiResp: apiResp,
	}
	listCityApiReply := ListCityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := city.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if city.option.EnableEncryption && nil != city.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch city.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", city.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(city.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					city.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listCityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listCityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listCityApiReply); err != nil {
				return nil, err
			}
		}
		listCityApiResp.ListCityApiReply = &listCityApiReply
	}
	return &listCityApiResp, nil
}

// ListCountry 国家
// Deprecated
func (city *city) ListCountry(ctx context.Context, req *ListCountryApiReq, reqOption *core.ReqOption) (*ListCountryApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/DemeterAres/Country/index"

	apiResp, err := core.Request(ctx, apiReq, city.option, reqOption)
	if err != nil {
		return nil, err
	}
	listCountryApiResp := ListCountryApiResp{
		ApiResp: apiResp,
	}
	listCountryApiReply := ListCountryApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := city.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if city.option.EnableEncryption && nil != city.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch city.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", city.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(city.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					city.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listCountryApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listCountryApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listCountryApiReply); err != nil {
				return nil, err
			}
		}
		listCountryApiResp.ListCountryApiReply = &listCountryApiReply
	}
	return &listCountryApiResp, nil
}

// ListHotelCity 酒店城市
// Deprecated
func (city *city) ListHotelCity(ctx context.Context, req *ListHotelCityApiReq, reqOption *core.ReqOption) (*ListHotelCityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/DemeterAres/HotelCity/index"

	apiResp, err := core.Request(ctx, apiReq, city.option, reqOption)
	if err != nil {
		return nil, err
	}
	listHotelCityApiResp := ListHotelCityApiResp{
		ApiResp: apiResp,
	}
	listHotelCityApiReply := ListHotelCityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := city.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if city.option.EnableEncryption && nil != city.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch city.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", city.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(city.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					city.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listHotelCityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listHotelCityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listHotelCityApiReply); err != nil {
				return nil, err
			}
		}
		listHotelCityApiResp.ListHotelCityApiReply = &listHotelCityApiReply
	}
	return &listHotelCityApiResp, nil
}

// ListTrainCity 火车票城市
// Deprecated
func (city *city) ListTrainCity(ctx context.Context, req *ListTrainCityApiReq, reqOption *core.ReqOption) (*ListTrainCityApiResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/river/DemeterAres/TrainCity"

	apiResp, err := core.Request(ctx, apiReq, city.option, reqOption)
	if err != nil {
		return nil, err
	}
	listTrainCityApiResp := ListTrainCityApiResp{
		ApiResp: apiResp,
	}
	listTrainCityApiReply := ListTrainCityApiReply{}
	if http.StatusOK == apiResp.StatusCode && nil != apiResp.Body {
		serializer := city.option.Serializer
		if nil != reqOption && nil != reqOption.Serializer {
			serializer = reqOption.Serializer
		}
		if city.option.EnableEncryption && nil != city.option.EncryptionOption {
			var respBodyMap map[string]interface{}
			if err := json.Unmarshal(apiResp.Body, &respBodyMap); err != nil {
				return nil, err
			}
			if value, ok := respBodyMap["encrypt_data"]; ok {
				if encryptData, ok := value.(string); ok {
					var decodedCiphertext []byte
					switch city.option.EncryptionOption.Ent {
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
						return nil, fmt.Errorf("未支持的Ent：%d", city.option.EncryptionOption.Ent)
					}
					decryptECB, err := core.AESDecryptECB(decodedCiphertext, []byte(city.option.EncryptionOption.Key))
					if err != nil {
						return nil, err
					}
					city.option.Logger.Debug(ctx, "decrypt data：", string(decryptECB))
					if err := serializer.Deserialize(decryptECB, &listTrainCityApiReply); err != nil {
						return nil, err
					}
				}
			} else {
				if err := serializer.Deserialize(apiResp.Body, &listTrainCityApiReply); err != nil {
					return nil, err
				}
			}
		} else {
			if err := serializer.Deserialize(apiResp.Body, &listTrainCityApiReply); err != nil {
				return nil, err
			}
		}
		listTrainCityApiResp.ListTrainCityApiReply = &listTrainCityApiReply
	}
	return &listTrainCityApiResp, nil
}
