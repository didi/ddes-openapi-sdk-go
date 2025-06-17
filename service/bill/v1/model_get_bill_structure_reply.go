package v1

// GetBillStructureReply struct for GetBillStructureReply
type GetBillStructureReply struct {
	List []BillListItem `json:"list,omitempty"` // 账单集合
}
