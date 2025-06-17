package v1

// ListTrainLeftTicketReply struct for ListTrainLeftTicketReply
type ListTrainLeftTicketReply struct {
	List []TrainLeftTicketListItem `json:"list,omitempty"` // 车次详细信息
}
