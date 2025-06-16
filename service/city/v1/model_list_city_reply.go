package v1

// ListCityReply struct for ListCityReply
type ListCityReply struct {
	Total   *int32       `json:"total,omitempty"`   // 查询结果条数
	Records []CityRecord `json:"records,omitempty"` // 返回具体的结果信息
}
