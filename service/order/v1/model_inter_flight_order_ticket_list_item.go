package v1

// InterFlightOrderTicketListItem 国际机票订单票张详情
type InterFlightOrderTicketListItem struct {
	TicketInfo *InterFlightOrderTicketInfo `json:"ticket_info,omitempty"`
}

type InterFlightOrderTicketListItemBuilder struct {
	ticketInfo    InterFlightOrderTicketInfo
	ticketInfoSet bool
}

func NewInterFlightOrderTicketListItemBuilder() *InterFlightOrderTicketListItemBuilder {
	return &InterFlightOrderTicketListItemBuilder{}
}
func (builder *InterFlightOrderTicketListItemBuilder) TicketInfo(ticketInfo InterFlightOrderTicketInfo) *InterFlightOrderTicketListItemBuilder {
	builder.ticketInfo = ticketInfo
	builder.ticketInfoSet = true
	return builder
}

func (builder *InterFlightOrderTicketListItemBuilder) Build() *InterFlightOrderTicketListItem {
	data := &InterFlightOrderTicketListItem{}
	if builder.ticketInfoSet {
		data.TicketInfo = &builder.ticketInfo
	}
	return data
}
