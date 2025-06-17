package v1

// TrainOrderTicketListItem struct for TrainOrderTicketListItem
type TrainOrderTicketListItem struct {
	TicketInfo *TrainOrderTicketInfo `json:"ticket_info,omitempty"`
}

type TrainOrderTicketListItemBuilder struct {
	ticketInfo    TrainOrderTicketInfo
	ticketInfoSet bool
}

func NewTrainOrderTicketListItemBuilder() *TrainOrderTicketListItemBuilder {
	return &TrainOrderTicketListItemBuilder{}
}
func (builder *TrainOrderTicketListItemBuilder) TicketInfo(ticketInfo TrainOrderTicketInfo) *TrainOrderTicketListItemBuilder {
	builder.ticketInfo = ticketInfo
	builder.ticketInfoSet = true
	return builder
}

func (builder *TrainOrderTicketListItemBuilder) Build() *TrainOrderTicketListItem {
	data := &TrainOrderTicketListItem{}
	if builder.ticketInfoSet {
		data.TicketInfo = &builder.ticketInfo
	}
	return data
}
