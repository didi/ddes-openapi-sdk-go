package v1

// ServicePackageContent 服务包内容
type ServicePackageContent struct {
	Name            *int32  `json:"name,omitempty"`              // 服务包内容名称
	Num             *int32  `json:"num,omitempty"`               // 内容对应数量
	Desc            *string `json:"desc,omitempty"`              // 服务包描述
	TicketUniqueKey *int32  `json:"ticket_unique_key,omitempty"` // 机票唯一标记
}

type ServicePackageContentBuilder struct {
	name               int32 // 服务包内容名称
	nameSet            bool
	num                int32 // 内容对应数量
	numSet             bool
	desc               string // 服务包描述
	descSet            bool
	ticketUniqueKey    int32 // 机票唯一标记
	ticketUniqueKeySet bool
}

func NewServicePackageContentBuilder() *ServicePackageContentBuilder {
	return &ServicePackageContentBuilder{}
}
func (builder *ServicePackageContentBuilder) Name(name int32) *ServicePackageContentBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *ServicePackageContentBuilder) Num(num int32) *ServicePackageContentBuilder {
	builder.num = num
	builder.numSet = true
	return builder
}
func (builder *ServicePackageContentBuilder) Desc(desc string) *ServicePackageContentBuilder {
	builder.desc = desc
	builder.descSet = true
	return builder
}
func (builder *ServicePackageContentBuilder) TicketUniqueKey(ticketUniqueKey int32) *ServicePackageContentBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}

func (builder *ServicePackageContentBuilder) Build() *ServicePackageContent {
	data := &ServicePackageContent{}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.numSet {
		data.Num = &builder.num
	}
	if builder.descSet {
		data.Desc = &builder.desc
	}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	return data
}
