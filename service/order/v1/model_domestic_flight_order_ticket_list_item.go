package v1

// DomesticFlightOrderTicketListItem 国内机票订单票张详情
type DomesticFlightOrderTicketListItem struct {
	TicketInfo *DomesticFlightOrderTicketInfo `json:"ticket_info,omitempty"`
}

type DomesticFlightOrderTicketListItemBuilder struct {
	ticketInfo    DomesticFlightOrderTicketInfo
	ticketInfoSet bool
}

func NewDomesticFlightOrderTicketListItemBuilder() *DomesticFlightOrderTicketListItemBuilder {
	return &DomesticFlightOrderTicketListItemBuilder{}
}
func (builder *DomesticFlightOrderTicketListItemBuilder) TicketInfo(ticketInfo DomesticFlightOrderTicketInfo) *DomesticFlightOrderTicketListItemBuilder {
	builder.ticketInfo = ticketInfo
	builder.ticketInfoSet = true
	return builder
}

func (builder *DomesticFlightOrderTicketListItemBuilder) Build() *DomesticFlightOrderTicketListItem {
	data := &DomesticFlightOrderTicketListItem{}
	if builder.ticketInfoSet {
		data.TicketInfo = &builder.ticketInfo
	}
	return data
}
