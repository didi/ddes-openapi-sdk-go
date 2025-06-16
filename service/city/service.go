package city

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/city/v1"
)

type Service struct {
	V1 *v1.V1
}

func NewService(option *core.Option) *Service {
	return &Service{
		V1: v1.NewV1(option),
	}
}
