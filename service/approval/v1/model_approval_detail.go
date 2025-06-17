package v1

// ApprovalDetail struct for ApprovalDetail
type ApprovalDetail struct {
	ApprovalId         *string                `json:"approval_id,omitempty"`
	OutApprovalId      *string                `json:"out_approval_id,omitempty"`
	ApprovalType       *int32                 `json:"approval_type,omitempty"`
	Type               *int32                 `json:"type,omitempty"`
	BudgetCenterId     *string                `json:"budget_center_id,omitempty"`
	Reason             *string                `json:"reason,omitempty"`
	ExtraInfo          *string                `json:"extra_info,omitempty"`
	StartDate          *string                `json:"start_date,omitempty"`
	EndDate            *string                `json:"end_date,omitempty"`
	ApprovalStatus     *int32                 `json:"approval_status,omitempty"`
	SceneId            *int32                 `json:"scene_id,omitempty"`
	CityType           *int32                 `json:"city_type,omitempty"`
	TravelBudget       *TravelBudget          `json:"travel_budget,omitempty"`
	PassengerList      []PassengerInfo        `json:"passenger_list,omitempty"`
	TravelerList       []Traveler             `json:"traveler_list,omitempty"`
	Applicant          *PassengerInfo         `json:"applicant,omitempty"`
	TravelDetail       *string                `json:"travel_detail,omitempty"`
	BusinessTripDetail *string                `json:"business_trip_detail,omitempty"`
	CarRule            []CarRule              `json:"car_rule,omitempty"`
	HotelRule          []HotelRule            `json:"hotel_rule,omitempty"`
	FlightRule         []FlightRule           `json:"flight_rule,omitempty"`
	TrainRule          []TrainRule            `json:"train_rule,omitempty"`
	BudgetCenterList   []BudgetCenterListItem `json:"budget_center_list,omitempty"`
}

type ApprovalDetailBuilder struct {
	approvalId            string
	approvalIdSet         bool
	outApprovalId         string
	outApprovalIdSet      bool
	approvalType          int32
	approvalTypeSet       bool
	type_                 int32
	type_Set              bool
	budgetCenterId        string
	budgetCenterIdSet     bool
	reason                string
	reasonSet             bool
	extraInfo             string
	extraInfoSet          bool
	startDate             string
	startDateSet          bool
	endDate               string
	endDateSet            bool
	approvalStatus        int32
	approvalStatusSet     bool
	sceneId               int32
	sceneIdSet            bool
	cityType              int32
	cityTypeSet           bool
	travelBudget          TravelBudget
	travelBudgetSet       bool
	passengerList         []PassengerInfo
	passengerListSet      bool
	travelerList          []Traveler
	travelerListSet       bool
	applicant             PassengerInfo
	applicantSet          bool
	travelDetail          string
	travelDetailSet       bool
	businessTripDetail    string
	businessTripDetailSet bool
	carRule               []CarRule
	carRuleSet            bool
	hotelRule             []HotelRule
	hotelRuleSet          bool
	flightRule            []FlightRule
	flightRuleSet         bool
	trainRule             []TrainRule
	trainRuleSet          bool
	budgetCenterList      []BudgetCenterListItem
	budgetCenterListSet   bool
}

func NewApprovalDetailBuilder() *ApprovalDetailBuilder {
	return &ApprovalDetailBuilder{}
}
func (builder *ApprovalDetailBuilder) ApprovalId(approvalId string) *ApprovalDetailBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) OutApprovalId(outApprovalId string) *ApprovalDetailBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) ApprovalType(approvalType int32) *ApprovalDetailBuilder {
	builder.approvalType = approvalType
	builder.approvalTypeSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) Type(type_ int32) *ApprovalDetailBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *ApprovalDetailBuilder) BudgetCenterId(budgetCenterId string) *ApprovalDetailBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) Reason(reason string) *ApprovalDetailBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) ExtraInfo(extraInfo string) *ApprovalDetailBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) StartDate(startDate string) *ApprovalDetailBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) EndDate(endDate string) *ApprovalDetailBuilder {
	builder.endDate = endDate
	builder.endDateSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) ApprovalStatus(approvalStatus int32) *ApprovalDetailBuilder {
	builder.approvalStatus = approvalStatus
	builder.approvalStatusSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) SceneId(sceneId int32) *ApprovalDetailBuilder {
	builder.sceneId = sceneId
	builder.sceneIdSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) CityType(cityType int32) *ApprovalDetailBuilder {
	builder.cityType = cityType
	builder.cityTypeSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) TravelBudget(travelBudget TravelBudget) *ApprovalDetailBuilder {
	builder.travelBudget = travelBudget
	builder.travelBudgetSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) PassengerList(passengerList []PassengerInfo) *ApprovalDetailBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) TravelerList(travelerList []Traveler) *ApprovalDetailBuilder {
	builder.travelerList = travelerList
	builder.travelerListSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) Applicant(applicant PassengerInfo) *ApprovalDetailBuilder {
	builder.applicant = applicant
	builder.applicantSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) TravelDetail(travelDetail string) *ApprovalDetailBuilder {
	builder.travelDetail = travelDetail
	builder.travelDetailSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) BusinessTripDetail(businessTripDetail string) *ApprovalDetailBuilder {
	builder.businessTripDetail = businessTripDetail
	builder.businessTripDetailSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) CarRule(carRule []CarRule) *ApprovalDetailBuilder {
	builder.carRule = carRule
	builder.carRuleSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) HotelRule(hotelRule []HotelRule) *ApprovalDetailBuilder {
	builder.hotelRule = hotelRule
	builder.hotelRuleSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) FlightRule(flightRule []FlightRule) *ApprovalDetailBuilder {
	builder.flightRule = flightRule
	builder.flightRuleSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) TrainRule(trainRule []TrainRule) *ApprovalDetailBuilder {
	builder.trainRule = trainRule
	builder.trainRuleSet = true
	return builder
}
func (builder *ApprovalDetailBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *ApprovalDetailBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}

func (builder *ApprovalDetailBuilder) Build() *ApprovalDetail {
	data := &ApprovalDetail{}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	if builder.approvalTypeSet {
		data.ApprovalType = &builder.approvalType
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.reasonSet {
		data.Reason = &builder.reason
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.startDateSet {
		data.StartDate = &builder.startDate
	}
	if builder.endDateSet {
		data.EndDate = &builder.endDate
	}
	if builder.approvalStatusSet {
		data.ApprovalStatus = &builder.approvalStatus
	}
	if builder.sceneIdSet {
		data.SceneId = &builder.sceneId
	}
	if builder.cityTypeSet {
		data.CityType = &builder.cityType
	}
	if builder.travelBudgetSet {
		data.TravelBudget = &builder.travelBudget
	}
	if builder.passengerListSet {
		data.PassengerList = builder.passengerList
	}
	if builder.travelerListSet {
		data.TravelerList = builder.travelerList
	}
	if builder.applicantSet {
		data.Applicant = &builder.applicant
	}
	if builder.travelDetailSet {
		data.TravelDetail = &builder.travelDetail
	}
	if builder.businessTripDetailSet {
		data.BusinessTripDetail = &builder.businessTripDetail
	}
	if builder.carRuleSet {
		data.CarRule = builder.carRule
	}
	if builder.hotelRuleSet {
		data.HotelRule = builder.hotelRule
	}
	if builder.flightRuleSet {
		data.FlightRule = builder.flightRule
	}
	if builder.trainRuleSet {
		data.TrainRule = builder.trainRule
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	return data
}
