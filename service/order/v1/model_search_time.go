package v1

// SearchTime search_time消息,时间传参
type SearchTime struct {
	Type      *string `json:"type,omitempty"`       // 时间类型，枚举值英文code，order_type 为 all 时，该字段可以传all.此时机票，酒店，火车为更新时间，用车为 pay_time,refund_time 支付时间、退款时间、个人支付转个人垫付时间，用于客户每天轮询订单数据落地使用；order_type 不为 all 时生效：机票条件：book_time 预订时间，exchange_time 改签时间 （多张票时，有一张票的改签成功时间满足就返回。），refund_time 退票时间 （多张票时，有一张票的退票成功时间满足就返回。），change_time 更新时间；酒店条件：book_time 预订时间，change_time 更新时间；火车票条件：book_time 预订时间，refund_time 退票时间 (多张票时，有一张票的退票成功时间满足就返回。），change_time 更新时间 ；用车：pay_time 支付时间，refund_time 退款时间
	StartTime *string `json:"start_time,omitempty"` // 开始时间，格式：yyyy-MM-dd HH:mm:ss 时间范围最大30天 最早时间距今不超过两年时间跨度
	EndTime   *string `json:"end_time,omitempty"`   // 结束时间，格式：yyyy-MM-dd HH:mm:ss 可查询的最早时间范围截止为当前时间往前300秒
}

type SearchTimeBuilder struct {
	type_        string // 时间类型，枚举值英文code，order_type 为 all 时，该字段可以传all.此时机票，酒店，火车为更新时间，用车为 pay_time,refund_time 支付时间、退款时间、个人支付转个人垫付时间，用于客户每天轮询订单数据落地使用；order_type 不为 all 时生效：机票条件：book_time 预订时间，exchange_time 改签时间 （多张票时，有一张票的改签成功时间满足就返回。），refund_time 退票时间 （多张票时，有一张票的退票成功时间满足就返回。），change_time 更新时间；酒店条件：book_time 预订时间，change_time 更新时间；火车票条件：book_time 预订时间，refund_time 退票时间 (多张票时，有一张票的退票成功时间满足就返回。），change_time 更新时间 ；用车：pay_time 支付时间，refund_time 退款时间
	type_Set     bool
	startTime    string // 开始时间，格式：yyyy-MM-dd HH:mm:ss 时间范围最大30天 最早时间距今不超过两年时间跨度
	startTimeSet bool
	endTime      string // 结束时间，格式：yyyy-MM-dd HH:mm:ss 可查询的最早时间范围截止为当前时间往前300秒
	endTimeSet   bool
}

func NewSearchTimeBuilder() *SearchTimeBuilder {
	return &SearchTimeBuilder{}
}
func (builder *SearchTimeBuilder) Type(type_ string) *SearchTimeBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *SearchTimeBuilder) StartTime(startTime string) *SearchTimeBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *SearchTimeBuilder) EndTime(endTime string) *SearchTimeBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}

func (builder *SearchTimeBuilder) Build() *SearchTime {
	data := &SearchTime{}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	return data
}
