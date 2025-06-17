package v1

// GetTrainOrderDetailReply struct for GetTrainOrderDetailReply
type GetTrainOrderDetailReply struct {
	Total             int32              `json:"total"` // 符合查询条件的数据总数
	DomestictrainData *DomesticTrainData `json:"domestictrain_data,omitempty"`
}
