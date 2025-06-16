package v1

// CountyRecord 区县信息
type CountyRecord struct {
	CountyId   *int32  `json:"county_id,omitempty"`   // 区县ID
	CountyName *string `json:"county_name,omitempty"` // 区县名称
}

type CountyRecordBuilder struct {
	countyId      int32 // 区县ID
	countyIdSet   bool
	countyName    string // 区县名称
	countyNameSet bool
}

func NewCountyRecordBuilder() *CountyRecordBuilder {
	return &CountyRecordBuilder{}
}
func (builder *CountyRecordBuilder) CountyId(countyId int32) *CountyRecordBuilder {
	builder.countyId = countyId
	builder.countyIdSet = true
	return builder
}
func (builder *CountyRecordBuilder) CountyName(countyName string) *CountyRecordBuilder {
	builder.countyName = countyName
	builder.countyNameSet = true
	return builder
}

func (builder *CountyRecordBuilder) Build() *CountyRecord {
	data := &CountyRecord{}
	if builder.countyIdSet {
		data.CountyId = &builder.countyId
	}
	if builder.countyNameSet {
		data.CountyName = &builder.countyName
	}
	return data
}
