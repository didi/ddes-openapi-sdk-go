package v1

// BillDetailOfDomesticFlightItem 已出账单 ~ 国内机票返回参数
type BillDetailOfDomesticFlightItem struct {
	AddedCutReason             *string  `json:"added_cut_reason,omitempty"`
	AddedEsCutFee              *float64 `json:"added_es_cut_fee,omitempty"`
	AddedGoodsName             *string  `json:"added_goods_name,omitempty"`
	Adjustment                 *string  `json:"adjustment,omitempty"`
	AirlineSimpleName          *string  `json:"airline_simple_name,omitempty"`
	ApplyChangeTime            *string  `json:"apply_change_time,omitempty"`
	ApplyRefundTime            *string  `json:"apply_refund_time,omitempty"`
	ApprovalChangeHistory      *string  `json:"approval_change_history,omitempty"`
	ApprovalFirst              *string  `json:"approval_first,omitempty"`
	ApprovalHistory            *string  `json:"approval_history,omitempty"`
	ApprovalId                 *string  `json:"approval_id,omitempty"`
	ApprovalNormalHistory      *string  `json:"approval_normal_history,omitempty"`
	ApprovalSecond             *string  `json:"approval_second,omitempty"`
	ApprovalThird              *string  `json:"approval_third,omitempty"`
	ArrivalAirportCode         *string  `json:"arrival_airport_code,omitempty"`
	ArrivalAirportName         *string  `json:"arrival_airport_name,omitempty"`
	ArrivalCityId              *string  `json:"arrival_city_id,omitempty"`
	ArrivalCityName            *string  `json:"arrival_city_name,omitempty"`
	ArrivalTime                *string  `json:"arrival_time,omitempty"`
	BaseFee                    *float64 `json:"base_fee,omitempty"`
	BeforeCutServiceFee        *float64 `json:"before_cut_service_fee,omitempty"`
	BillId                     *int64   `json:"bill_id,omitempty"`
	BookingDate                *string  `json:"booking_date,omitempty"`
	BookingDepCode             *string  `json:"booking_dep_code,omitempty"`
	BookingDepName             *string  `json:"booking_dep_name,omitempty"`
	BookingMemberId            *int64   `json:"booking_member_id,omitempty"`
	BookingMemberName          *string  `json:"booking_member_name,omitempty"`
	BookingMemberNumber        *string  `json:"booking_member_number,omitempty"`
	BookingParentPathDepName   *string  `json:"booking_parent_path_dep_name,omitempty"`
	BudgetCenterCode           *string  `json:"budget_center_code,omitempty"`
	BudgetCenterId             *string  `json:"budget_center_id,omitempty"`
	BudgetCenterName           *string  `json:"budget_center_name,omitempty"`
	BudgetCenterParentPathName *string  `json:"budget_center_parent_path_name,omitempty"`
	CabinCode                  *string  `json:"cabin_code,omitempty"`
	CabinName                  *string  `json:"cabin_name,omitempty"`
	CabinType                  *int32   `json:"cabin_type,omitempty"`
	ChangeHandleCost           *string  `json:"change_handle_cost,omitempty"`
	CompanyChangeServiceCost   *string  `json:"company_change_service_cost,omitempty"`
	CompanyRealPay             *string  `json:"company_real_pay,omitempty"`
	CompanyTicketCost          *string  `json:"company_ticket_cost,omitempty"`
	CompanyUpgradeCost         *string  `json:"company_upgrade_cost,omitempty"`
	ConstructionCost           *string  `json:"construction_cost,omitempty"`
	CreateTime                 *string  `json:"create_time,omitempty"`
	CurrentCarbonEmission      *string  `json:"current_carbon_emission,omitempty"`
	DaysInAdvance              *int32   `json:"days_in_advance,omitempty"`
	Department                 *string  `json:"department,omitempty"`
	DepartmentId               *string  `json:"department_id,omitempty"`
	DepartureAirportCode       *string  `json:"departure_airport_code,omitempty"`
	DepartureAirportName       *string  `json:"departure_airport_name,omitempty"`
	DepartureCityId            *string  `json:"departure_city_id,omitempty"`
	DepartureCityName          *string  `json:"departure_city_name,omitempty"`
	DepartureTime              *string  `json:"departure_time,omitempty"`
	Discount                   *string  `json:"discount,omitempty"`
	EconomyOriginalCost        *string  `json:"economy_original_cost,omitempty"`
	EmployeeNumber             *string  `json:"employee_number,omitempty"`
	ExInfo01                   *string  `json:"ex_info_01,omitempty"`
	ExInfo01Code               *string  `json:"ex_info_01_code,omitempty"`
	ExInfo02                   *string  `json:"ex_info_02,omitempty"`
	ExInfo02Code               *string  `json:"ex_info_02_code,omitempty"`
	ExInfo03                   *string  `json:"ex_info_03,omitempty"`
	ExInfo03Code               *string  `json:"ex_info_03_code,omitempty"`
	ExInfo04                   *string  `json:"ex_info_04,omitempty"`
	ExInfo04Code               *string  `json:"ex_info_04_code,omitempty"`
	ExInfo05                   *string  `json:"ex_info_05,omitempty"`
	ExInfo05Code               *string  `json:"ex_info_05_code,omitempty"`
	ExInfo06                   *string  `json:"ex_info_06,omitempty"`
	ExInfo06Code               *string  `json:"ex_info_06_code,omitempty"`
	ExInfo07                   *string  `json:"ex_info_07,omitempty"`
	ExInfo07Code               *string  `json:"ex_info_07_code,omitempty"`
	ExInfo08                   *string  `json:"ex_info_08,omitempty"`
	ExInfo08Code               *string  `json:"ex_info_08_code,omitempty"`
	ExceptionInfo              *string  `json:"exception_info,omitempty"`
	ExcludingServiceFee        *float64 `json:"excluding_service_fee,omitempty"`
	ExemptReason               *string  `json:"exempt_reason,omitempty"`
	ExemptType                 *string  `json:"exempt_type,omitempty"`
	ExtendInfo                 *string  `json:"extend_info,omitempty"`
	ExtraInfo                  *string  `json:"extra_info,omitempty"`
	FinalDepartureTime         *string  `json:"final_departure_time,omitempty"`
	FinalOfflineStatus         *string  `json:"final_offline_status,omitempty"`
	FlightDistance             *string  `json:"flight_distance,omitempty"`
	FlightNumber               *string  `json:"flight_number,omitempty"`
	FlightSegmentNumber        *int32   `json:"flight_segment_number,omitempty"`
	FlightSegmentTravel        *string  `json:"flight_segment_travel,omitempty"`
	FlightTravelType           *int32   `json:"flight_travel_type,omitempty"`
	FuelCost                   *string  `json:"fuelCost,omitempty"`
	Group1Code                 *string  `json:"group_1_code,omitempty"`
	Group1Name                 *string  `json:"group_1_name,omitempty"`
	Group2Code                 *string  `json:"group_2_code,omitempty"`
	Group2Name                 *string  `json:"group_2_name,omitempty"`
	Group3Code                 *string  `json:"group_3_code,omitempty"`
	Group3Name                 *string  `json:"group_3_name,omitempty"`
	Group4Code                 *string  `json:"group_4_code,omitempty"`
	Group4Name                 *string  `json:"group_4_name,omitempty"`
	Group5Code                 *string  `json:"group_5_code,omitempty"`
	Group5Name                 *string  `json:"group_5_name,omitempty"`
	Group6Code                 *string  `json:"group_6_code,omitempty"`
	Group6Name                 *string  `json:"group_6_name,omitempty"`
	Group7Code                 *string  `json:"group_7_code,omitempty"`
	Group7Name                 *string  `json:"group_7_name,omitempty"`
	Group8Code                 *string  `json:"group_8_code,omitempty"`
	Group8Name                 *string  `json:"group_8_name,omitempty"`
	Group9Code                 *string  `json:"group_9_code,omitempty"`
	Group9Name                 *string  `json:"group_9_name,omitempty"`
	InstitutionId              *int64   `json:"institution_id,omitempty"`
	InstitutionName            *string  `json:"institution_name,omitempty"`
	InvoiceEntityInfo          *string  `json:"invoice_entity_info,omitempty"`
	InvoiceFee                 *string  `json:"invoice_fee,omitempty"`
	IsAdded                    *int32   `json:"is_added,omitempty"`
	IsAgreement                *string  `json:"is_agreement,omitempty"`
	IsDiscountTicket           *string  `json:"is_discount_ticket,omitempty"`
	IsSensitive                *string  `json:"is_sensitive,omitempty"`
	IsTraveler                 *string  `json:"is_traveler,omitempty"`
	LastupdateTime             *string  `json:"lastupdate_time,omitempty"`
	LegalEntityId              *int64   `json:"legal_entity_id,omitempty"`
	LegalEntityName            *string  `json:"legal_entity_name,omitempty"`
	MainFlightNumber           *string  `json:"main_flight_number,omitempty"`
	MemberId                   *string  `json:"member_id,omitempty"`
	OrderFixSettleTime         *string  `json:"order_fix_settle_time,omitempty"`
	OrderFixType               *int32   `json:"order_fix_type,omitempty"`
	OrderId                    *string  `json:"order_id,omitempty"`
	OriginOrderId              *string  `json:"origin_order_id,omitempty"`
	OriginTicketId             *string  `json:"origin_ticket_id,omitempty"`
	OutLegalEntityId           *string  `json:"out_legal_entity_id,omitempty"`
	OutRequisitionId           *string  `json:"out_requisition_id,omitempty"`
	ParentInstitutionId        *int64   `json:"parent_institution_id,omitempty"`
	ParentInstitutionName      *string  `json:"parent_institution_name,omitempty"`
	PassengerDepCode           *string  `json:"passenger_dep_code,omitempty"`
	PassengerDepName           *string  `json:"passenger_dep_name,omitempty"`
	PassengerGroup1Code        *string  `json:"passenger_group_1_code,omitempty"`
	PassengerGroup1Name        *string  `json:"passenger_group_1_name,omitempty"`
	PassengerGroup2Code        *string  `json:"passenger_group_2_code,omitempty"`
	PassengerGroup2Name        *string  `json:"passenger_group_2_name,omitempty"`
	PassengerGroup3Code        *string  `json:"passenger_group_3_code,omitempty"`
	PassengerGroup3Name        *string  `json:"passenger_group_3_name,omitempty"`
	PassengerGroup4Code        *string  `json:"passenger_group_4_code,omitempty"`
	PassengerGroup4Name        *string  `json:"passenger_group_4_name,omitempty"`
	PassengerGroup5Code        *string  `json:"passenger_group_5_code,omitempty"`
	PassengerGroup5Name        *string  `json:"passenger_group_5_name,omitempty"`
	PassengerGroup6Code        *string  `json:"passenger_group_6_code,omitempty"`
	PassengerGroup6Name        *string  `json:"passenger_group_6_name,omitempty"`
	PassengerGroup7Code        *string  `json:"passenger_group_7_code,omitempty"`
	PassengerGroup7Name        *string  `json:"passenger_group_7_name,omitempty"`
	PassengerGroup8Code        *string  `json:"passenger_group_8_code,omitempty"`
	PassengerGroup8Name        *string  `json:"passenger_group_8_name,omitempty"`
	PassengerGroup9Code        *string  `json:"passenger_group_9_code,omitempty"`
	PassengerGroup9Name        *string  `json:"passenger_group_9_name,omitempty"`
	PassengerLegalEntityId     *int64   `json:"passenger_legal_entity_id,omitempty"`
	PassengerLegalEntityName   *string  `json:"passenger_legal_entity_name,omitempty"`
	PassengerMemberId          *int64   `json:"passenger_member_id,omitempty"`
	PassengerMemberName        *string  `json:"passenger_member_name,omitempty"`
	PassengerMemberNumber      *string  `json:"passenger_member_number,omitempty"`
	PassengerName              *string  `json:"passenger_name,omitempty"`
	PassengerParentPathDepName *string  `json:"passenger_parent_path_dep_name,omitempty"`
	PayChannel                 *string  `json:"pay_channel,omitempty"`
	PayType                    *string  `json:"pay_type,omitempty"`
	PersonalChangeServiceCost  *string  `json:"personal_change_service_cost,omitempty"`
	PersonalRealPay            *string  `json:"personal_real_pay,omitempty"`
	PersonalTicketCost         *string  `json:"personal_ticket_cost,omitempty"`
	PersonalUpgradeCost        *string  `json:"personal_upgrade_cost,omitempty"`
	PnrNumber                  *string  `json:"pnr_number,omitempty"`
	PreOrderId                 *string  `json:"pre_order_id,omitempty"`
	ProjectExtInfo             *string  `json:"project_ext_info,omitempty"`
	RankName                   *string  `json:"rank_name,omitempty"`
	RcBook                     *string  `json:"rc_book,omitempty"`
	RcDistancePeriod           *string  `json:"rc_distance_period,omitempty"`
	RcLevel                    *string  `json:"rc_level,omitempty"`
	RcLowPrice                 *string  `json:"rc_lowPrice,omitempty"`
	RcTimePeriod               *string  `json:"rc_time_period,omitempty"`
	Reason                     *string  `json:"reason,omitempty"`
	Rebook                     *string  `json:"rebook,omitempty"`
	ReductionCarbonEmission    *string  `json:"reduction_carbon_emission,omitempty"`
	ReductionFee               *float64 `json:"reduction_fee,omitempty"`
	RefundHandleCost           *string  `json:"refund_handle_cost,omitempty"`
	Remark                     *string  `json:"remark,omitempty"`
	RequisitionId              *string  `json:"requisition_id,omitempty"`
	ServiceFee                 *string  `json:"service_fee,omitempty"`
	SettleTime                 *string  `json:"settle_time,omitempty"`
	SettlementTime             *string  `json:"settlement_time,omitempty"`
	SubAccountCompanyName      *string  `json:"sub_account_company_name,omitempty"`
	SubCompanyNo               *string  `json:"sub_company_no,omitempty"`
	SupplierOrderId            *string  `json:"supplier_order_id,omitempty"`
	SupplierTicketNumber       *string  `json:"supplier_ticket_number,omitempty"`
	TicketFee                  *string  `json:"ticket_fee,omitempty"`
	TicketIdText               *string  `json:"ticket_id_text,omitempty"`
	TotalFee                   *string  `json:"total_fee,omitempty"`
	TrainComparePrice          *float64 `json:"train_compare_price,omitempty"`
	TransactionType            *int32   `json:"transaction_type,omitempty"`
	TravelItineraryFee         *string  `json:"travel_itinerary_fee,omitempty"`
	TravelPurpose              *string  `json:"travel_purpose,omitempty"`
	TripDescription            *string  `json:"trip_description,omitempty"`
	TripReason                 *string  `json:"trip_reason,omitempty"`
	UniqueKey                  *string  `json:"unique_key,omitempty"`
	UnusualContent             *string  `json:"unusual_content,omitempty"`
	UnusualReason              *string  `json:"unusual_reason,omitempty"`
	UnusualRemark              *string  `json:"unusual_remark,omitempty"`
	UnusualType                *string  `json:"unusual_type,omitempty"`
	UpgradeCost                *string  `json:"upgrade_cost,omitempty"`
	UpgradeFinalStatus         *int32   `json:"upgrade_final_status,omitempty"`
}

type BillDetailOfDomesticFlightItemBuilder struct {
	addedCutReason                string
	addedCutReasonSet             bool
	addedEsCutFee                 float64
	addedEsCutFeeSet              bool
	addedGoodsName                string
	addedGoodsNameSet             bool
	adjustment                    string
	adjustmentSet                 bool
	airlineSimpleName             string
	airlineSimpleNameSet          bool
	applyChangeTime               string
	applyChangeTimeSet            bool
	applyRefundTime               string
	applyRefundTimeSet            bool
	approvalChangeHistory         string
	approvalChangeHistorySet      bool
	approvalFirst                 string
	approvalFirstSet              bool
	approvalHistory               string
	approvalHistorySet            bool
	approvalId                    string
	approvalIdSet                 bool
	approvalNormalHistory         string
	approvalNormalHistorySet      bool
	approvalSecond                string
	approvalSecondSet             bool
	approvalThird                 string
	approvalThirdSet              bool
	arrivalAirportCode            string
	arrivalAirportCodeSet         bool
	arrivalAirportName            string
	arrivalAirportNameSet         bool
	arrivalCityId                 string
	arrivalCityIdSet              bool
	arrivalCityName               string
	arrivalCityNameSet            bool
	arrivalTime                   string
	arrivalTimeSet                bool
	baseFee                       float64
	baseFeeSet                    bool
	beforeCutServiceFee           float64
	beforeCutServiceFeeSet        bool
	billId                        int64
	billIdSet                     bool
	bookingDate                   string
	bookingDateSet                bool
	bookingDepCode                string
	bookingDepCodeSet             bool
	bookingDepName                string
	bookingDepNameSet             bool
	bookingMemberId               int64
	bookingMemberIdSet            bool
	bookingMemberName             string
	bookingMemberNameSet          bool
	bookingMemberNumber           string
	bookingMemberNumberSet        bool
	bookingParentPathDepName      string
	bookingParentPathDepNameSet   bool
	budgetCenterCode              string
	budgetCenterCodeSet           bool
	budgetCenterId                string
	budgetCenterIdSet             bool
	budgetCenterName              string
	budgetCenterNameSet           bool
	budgetCenterParentPathName    string
	budgetCenterParentPathNameSet bool
	cabinCode                     string
	cabinCodeSet                  bool
	cabinName                     string
	cabinNameSet                  bool
	cabinType                     int32
	cabinTypeSet                  bool
	changeHandleCost              string
	changeHandleCostSet           bool
	companyChangeServiceCost      string
	companyChangeServiceCostSet   bool
	companyRealPay                string
	companyRealPaySet             bool
	companyTicketCost             string
	companyTicketCostSet          bool
	companyUpgradeCost            string
	companyUpgradeCostSet         bool
	constructionCost              string
	constructionCostSet           bool
	createTime                    string
	createTimeSet                 bool
	currentCarbonEmission         string
	currentCarbonEmissionSet      bool
	daysInAdvance                 int32
	daysInAdvanceSet              bool
	department                    string
	departmentSet                 bool
	departmentId                  string
	departmentIdSet               bool
	departureAirportCode          string
	departureAirportCodeSet       bool
	departureAirportName          string
	departureAirportNameSet       bool
	departureCityId               string
	departureCityIdSet            bool
	departureCityName             string
	departureCityNameSet          bool
	departureTime                 string
	departureTimeSet              bool
	discount                      string
	discountSet                   bool
	economyOriginalCost           string
	economyOriginalCostSet        bool
	employeeNumber                string
	employeeNumberSet             bool
	exInfo01                      string
	exInfo01Set                   bool
	exInfo01Code                  string
	exInfo01CodeSet               bool
	exInfo02                      string
	exInfo02Set                   bool
	exInfo02Code                  string
	exInfo02CodeSet               bool
	exInfo03                      string
	exInfo03Set                   bool
	exInfo03Code                  string
	exInfo03CodeSet               bool
	exInfo04                      string
	exInfo04Set                   bool
	exInfo04Code                  string
	exInfo04CodeSet               bool
	exInfo05                      string
	exInfo05Set                   bool
	exInfo05Code                  string
	exInfo05CodeSet               bool
	exInfo06                      string
	exInfo06Set                   bool
	exInfo06Code                  string
	exInfo06CodeSet               bool
	exInfo07                      string
	exInfo07Set                   bool
	exInfo07Code                  string
	exInfo07CodeSet               bool
	exInfo08                      string
	exInfo08Set                   bool
	exInfo08Code                  string
	exInfo08CodeSet               bool
	exceptionInfo                 string
	exceptionInfoSet              bool
	excludingServiceFee           float64
	excludingServiceFeeSet        bool
	exemptReason                  string
	exemptReasonSet               bool
	exemptType                    string
	exemptTypeSet                 bool
	extendInfo                    string
	extendInfoSet                 bool
	extraInfo                     string
	extraInfoSet                  bool
	finalDepartureTime            string
	finalDepartureTimeSet         bool
	finalOfflineStatus            string
	finalOfflineStatusSet         bool
	flightDistance                string
	flightDistanceSet             bool
	flightNumber                  string
	flightNumberSet               bool
	flightSegmentNumber           int32
	flightSegmentNumberSet        bool
	flightSegmentTravel           string
	flightSegmentTravelSet        bool
	flightTravelType              int32
	flightTravelTypeSet           bool
	fuelCost                      string
	fuelCostSet                   bool
	group1Code                    string
	group1CodeSet                 bool
	group1Name                    string
	group1NameSet                 bool
	group2Code                    string
	group2CodeSet                 bool
	group2Name                    string
	group2NameSet                 bool
	group3Code                    string
	group3CodeSet                 bool
	group3Name                    string
	group3NameSet                 bool
	group4Code                    string
	group4CodeSet                 bool
	group4Name                    string
	group4NameSet                 bool
	group5Code                    string
	group5CodeSet                 bool
	group5Name                    string
	group5NameSet                 bool
	group6Code                    string
	group6CodeSet                 bool
	group6Name                    string
	group6NameSet                 bool
	group7Code                    string
	group7CodeSet                 bool
	group7Name                    string
	group7NameSet                 bool
	group8Code                    string
	group8CodeSet                 bool
	group8Name                    string
	group8NameSet                 bool
	group9Code                    string
	group9CodeSet                 bool
	group9Name                    string
	group9NameSet                 bool
	institutionId                 int64
	institutionIdSet              bool
	institutionName               string
	institutionNameSet            bool
	invoiceEntityInfo             string
	invoiceEntityInfoSet          bool
	invoiceFee                    string
	invoiceFeeSet                 bool
	isAdded                       int32
	isAddedSet                    bool
	isAgreement                   string
	isAgreementSet                bool
	isDiscountTicket              string
	isDiscountTicketSet           bool
	isSensitive                   string
	isSensitiveSet                bool
	isTraveler                    string
	isTravelerSet                 bool
	lastupdateTime                string
	lastupdateTimeSet             bool
	legalEntityId                 int64
	legalEntityIdSet              bool
	legalEntityName               string
	legalEntityNameSet            bool
	mainFlightNumber              string
	mainFlightNumberSet           bool
	memberId                      string
	memberIdSet                   bool
	orderFixSettleTime            string
	orderFixSettleTimeSet         bool
	orderFixType                  int32
	orderFixTypeSet               bool
	orderId                       string
	orderIdSet                    bool
	originOrderId                 string
	originOrderIdSet              bool
	originTicketId                string
	originTicketIdSet             bool
	outLegalEntityId              string
	outLegalEntityIdSet           bool
	outRequisitionId              string
	outRequisitionIdSet           bool
	parentInstitutionId           int64
	parentInstitutionIdSet        bool
	parentInstitutionName         string
	parentInstitutionNameSet      bool
	passengerDepCode              string
	passengerDepCodeSet           bool
	passengerDepName              string
	passengerDepNameSet           bool
	passengerGroup1Code           string
	passengerGroup1CodeSet        bool
	passengerGroup1Name           string
	passengerGroup1NameSet        bool
	passengerGroup2Code           string
	passengerGroup2CodeSet        bool
	passengerGroup2Name           string
	passengerGroup2NameSet        bool
	passengerGroup3Code           string
	passengerGroup3CodeSet        bool
	passengerGroup3Name           string
	passengerGroup3NameSet        bool
	passengerGroup4Code           string
	passengerGroup4CodeSet        bool
	passengerGroup4Name           string
	passengerGroup4NameSet        bool
	passengerGroup5Code           string
	passengerGroup5CodeSet        bool
	passengerGroup5Name           string
	passengerGroup5NameSet        bool
	passengerGroup6Code           string
	passengerGroup6CodeSet        bool
	passengerGroup6Name           string
	passengerGroup6NameSet        bool
	passengerGroup7Code           string
	passengerGroup7CodeSet        bool
	passengerGroup7Name           string
	passengerGroup7NameSet        bool
	passengerGroup8Code           string
	passengerGroup8CodeSet        bool
	passengerGroup8Name           string
	passengerGroup8NameSet        bool
	passengerGroup9Code           string
	passengerGroup9CodeSet        bool
	passengerGroup9Name           string
	passengerGroup9NameSet        bool
	passengerLegalEntityId        int64
	passengerLegalEntityIdSet     bool
	passengerLegalEntityName      string
	passengerLegalEntityNameSet   bool
	passengerMemberId             int64
	passengerMemberIdSet          bool
	passengerMemberName           string
	passengerMemberNameSet        bool
	passengerMemberNumber         string
	passengerMemberNumberSet      bool
	passengerName                 string
	passengerNameSet              bool
	passengerParentPathDepName    string
	passengerParentPathDepNameSet bool
	payChannel                    string
	payChannelSet                 bool
	payType                       string
	payTypeSet                    bool
	personalChangeServiceCost     string
	personalChangeServiceCostSet  bool
	personalRealPay               string
	personalRealPaySet            bool
	personalTicketCost            string
	personalTicketCostSet         bool
	personalUpgradeCost           string
	personalUpgradeCostSet        bool
	pnrNumber                     string
	pnrNumberSet                  bool
	preOrderId                    string
	preOrderIdSet                 bool
	projectExtInfo                string
	projectExtInfoSet             bool
	rankName                      string
	rankNameSet                   bool
	rcBook                        string
	rcBookSet                     bool
	rcDistancePeriod              string
	rcDistancePeriodSet           bool
	rcLevel                       string
	rcLevelSet                    bool
	rcLowPrice                    string
	rcLowPriceSet                 bool
	rcTimePeriod                  string
	rcTimePeriodSet               bool
	reason                        string
	reasonSet                     bool
	rebook                        string
	rebookSet                     bool
	reductionCarbonEmission       string
	reductionCarbonEmissionSet    bool
	reductionFee                  float64
	reductionFeeSet               bool
	refundHandleCost              string
	refundHandleCostSet           bool
	remark                        string
	remarkSet                     bool
	requisitionId                 string
	requisitionIdSet              bool
	serviceFee                    string
	serviceFeeSet                 bool
	settleTime                    string
	settleTimeSet                 bool
	settlementTime                string
	settlementTimeSet             bool
	subAccountCompanyName         string
	subAccountCompanyNameSet      bool
	subCompanyNo                  string
	subCompanyNoSet               bool
	supplierOrderId               string
	supplierOrderIdSet            bool
	supplierTicketNumber          string
	supplierTicketNumberSet       bool
	ticketFee                     string
	ticketFeeSet                  bool
	ticketIdText                  string
	ticketIdTextSet               bool
	totalFee                      string
	totalFeeSet                   bool
	trainComparePrice             float64
	trainComparePriceSet          bool
	transactionType               int32
	transactionTypeSet            bool
	travelItineraryFee            string
	travelItineraryFeeSet         bool
	travelPurpose                 string
	travelPurposeSet              bool
	tripDescription               string
	tripDescriptionSet            bool
	tripReason                    string
	tripReasonSet                 bool
	uniqueKey                     string
	uniqueKeySet                  bool
	unusualContent                string
	unusualContentSet             bool
	unusualReason                 string
	unusualReasonSet              bool
	unusualRemark                 string
	unusualRemarkSet              bool
	unusualType                   string
	unusualTypeSet                bool
	upgradeCost                   string
	upgradeCostSet                bool
	upgradeFinalStatus            int32
	upgradeFinalStatusSet         bool
}

func NewBillDetailOfDomesticFlightItemBuilder() *BillDetailOfDomesticFlightItemBuilder {
	return &BillDetailOfDomesticFlightItemBuilder{}
}
func (builder *BillDetailOfDomesticFlightItemBuilder) AddedCutReason(addedCutReason string) *BillDetailOfDomesticFlightItemBuilder {
	builder.addedCutReason = addedCutReason
	builder.addedCutReasonSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) AddedEsCutFee(addedEsCutFee float64) *BillDetailOfDomesticFlightItemBuilder {
	builder.addedEsCutFee = addedEsCutFee
	builder.addedEsCutFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) AddedGoodsName(addedGoodsName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.addedGoodsName = addedGoodsName
	builder.addedGoodsNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Adjustment(adjustment string) *BillDetailOfDomesticFlightItemBuilder {
	builder.adjustment = adjustment
	builder.adjustmentSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) AirlineSimpleName(airlineSimpleName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.airlineSimpleName = airlineSimpleName
	builder.airlineSimpleNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApplyChangeTime(applyChangeTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.applyChangeTime = applyChangeTime
	builder.applyChangeTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApplyRefundTime(applyRefundTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.applyRefundTime = applyRefundTime
	builder.applyRefundTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApprovalChangeHistory(approvalChangeHistory string) *BillDetailOfDomesticFlightItemBuilder {
	builder.approvalChangeHistory = approvalChangeHistory
	builder.approvalChangeHistorySet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApprovalFirst(approvalFirst string) *BillDetailOfDomesticFlightItemBuilder {
	builder.approvalFirst = approvalFirst
	builder.approvalFirstSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApprovalHistory(approvalHistory string) *BillDetailOfDomesticFlightItemBuilder {
	builder.approvalHistory = approvalHistory
	builder.approvalHistorySet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApprovalId(approvalId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApprovalNormalHistory(approvalNormalHistory string) *BillDetailOfDomesticFlightItemBuilder {
	builder.approvalNormalHistory = approvalNormalHistory
	builder.approvalNormalHistorySet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApprovalSecond(approvalSecond string) *BillDetailOfDomesticFlightItemBuilder {
	builder.approvalSecond = approvalSecond
	builder.approvalSecondSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ApprovalThird(approvalThird string) *BillDetailOfDomesticFlightItemBuilder {
	builder.approvalThird = approvalThird
	builder.approvalThirdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ArrivalAirportCode(arrivalAirportCode string) *BillDetailOfDomesticFlightItemBuilder {
	builder.arrivalAirportCode = arrivalAirportCode
	builder.arrivalAirportCodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ArrivalAirportName(arrivalAirportName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.arrivalAirportName = arrivalAirportName
	builder.arrivalAirportNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ArrivalCityId(arrivalCityId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ArrivalCityName(arrivalCityName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.arrivalCityName = arrivalCityName
	builder.arrivalCityNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ArrivalTime(arrivalTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.arrivalTime = arrivalTime
	builder.arrivalTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BaseFee(baseFee float64) *BillDetailOfDomesticFlightItemBuilder {
	builder.baseFee = baseFee
	builder.baseFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BeforeCutServiceFee(beforeCutServiceFee float64) *BillDetailOfDomesticFlightItemBuilder {
	builder.beforeCutServiceFee = beforeCutServiceFee
	builder.beforeCutServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BillId(billId int64) *BillDetailOfDomesticFlightItemBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BookingDate(bookingDate string) *BillDetailOfDomesticFlightItemBuilder {
	builder.bookingDate = bookingDate
	builder.bookingDateSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BookingDepCode(bookingDepCode string) *BillDetailOfDomesticFlightItemBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BookingDepName(bookingDepName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.bookingDepName = bookingDepName
	builder.bookingDepNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BookingMemberId(bookingMemberId int64) *BillDetailOfDomesticFlightItemBuilder {
	builder.bookingMemberId = bookingMemberId
	builder.bookingMemberIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BookingMemberName(bookingMemberName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BookingMemberNumber(bookingMemberNumber string) *BillDetailOfDomesticFlightItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BookingParentPathDepName(bookingParentPathDepName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.bookingParentPathDepName = bookingParentPathDepName
	builder.bookingParentPathDepNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BudgetCenterCode(budgetCenterCode string) *BillDetailOfDomesticFlightItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BudgetCenterId(budgetCenterId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BudgetCenterName(budgetCenterName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) BudgetCenterParentPathName(budgetCenterParentPathName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.budgetCenterParentPathName = budgetCenterParentPathName
	builder.budgetCenterParentPathNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CabinCode(cabinCode string) *BillDetailOfDomesticFlightItemBuilder {
	builder.cabinCode = cabinCode
	builder.cabinCodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CabinName(cabinName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.cabinName = cabinName
	builder.cabinNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CabinType(cabinType int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.cabinType = cabinType
	builder.cabinTypeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ChangeHandleCost(changeHandleCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.changeHandleCost = changeHandleCost
	builder.changeHandleCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CompanyChangeServiceCost(companyChangeServiceCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.companyChangeServiceCost = companyChangeServiceCost
	builder.companyChangeServiceCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CompanyRealPay(companyRealPay string) *BillDetailOfDomesticFlightItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CompanyTicketCost(companyTicketCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.companyTicketCost = companyTicketCost
	builder.companyTicketCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CompanyUpgradeCost(companyUpgradeCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.companyUpgradeCost = companyUpgradeCost
	builder.companyUpgradeCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ConstructionCost(constructionCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.constructionCost = constructionCost
	builder.constructionCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CreateTime(createTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) CurrentCarbonEmission(currentCarbonEmission string) *BillDetailOfDomesticFlightItemBuilder {
	builder.currentCarbonEmission = currentCarbonEmission
	builder.currentCarbonEmissionSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) DaysInAdvance(daysInAdvance int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.daysInAdvance = daysInAdvance
	builder.daysInAdvanceSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Department(department string) *BillDetailOfDomesticFlightItemBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) DepartmentId(departmentId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.departmentId = departmentId
	builder.departmentIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) DepartureAirportCode(departureAirportCode string) *BillDetailOfDomesticFlightItemBuilder {
	builder.departureAirportCode = departureAirportCode
	builder.departureAirportCodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) DepartureAirportName(departureAirportName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.departureAirportName = departureAirportName
	builder.departureAirportNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) DepartureCityId(departureCityId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) DepartureCityName(departureCityName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.departureCityName = departureCityName
	builder.departureCityNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) DepartureTime(departureTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Discount(discount string) *BillDetailOfDomesticFlightItemBuilder {
	builder.discount = discount
	builder.discountSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) EconomyOriginalCost(economyOriginalCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.economyOriginalCost = economyOriginalCost
	builder.economyOriginalCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) EmployeeNumber(employeeNumber string) *BillDetailOfDomesticFlightItemBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo01(exInfo01 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo01Code(exInfo01Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo02(exInfo02 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo02Code(exInfo02Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo03(exInfo03 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo03Code(exInfo03Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo04(exInfo04 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo04Code(exInfo04Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo05(exInfo05 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo05Code(exInfo05Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo06(exInfo06 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo06Code(exInfo06Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo07(exInfo07 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo07Code(exInfo07Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo08(exInfo08 string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExInfo08Code(exInfo08Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExceptionInfo(exceptionInfo string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exceptionInfo = exceptionInfo
	builder.exceptionInfoSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExcludingServiceFee(excludingServiceFee float64) *BillDetailOfDomesticFlightItemBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExemptReason(exemptReason string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exemptReason = exemptReason
	builder.exemptReasonSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExemptType(exemptType string) *BillDetailOfDomesticFlightItemBuilder {
	builder.exemptType = exemptType
	builder.exemptTypeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExtendInfo(extendInfo string) *BillDetailOfDomesticFlightItemBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ExtraInfo(extraInfo string) *BillDetailOfDomesticFlightItemBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FinalDepartureTime(finalDepartureTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.finalDepartureTime = finalDepartureTime
	builder.finalDepartureTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FinalOfflineStatus(finalOfflineStatus string) *BillDetailOfDomesticFlightItemBuilder {
	builder.finalOfflineStatus = finalOfflineStatus
	builder.finalOfflineStatusSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FlightDistance(flightDistance string) *BillDetailOfDomesticFlightItemBuilder {
	builder.flightDistance = flightDistance
	builder.flightDistanceSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FlightNumber(flightNumber string) *BillDetailOfDomesticFlightItemBuilder {
	builder.flightNumber = flightNumber
	builder.flightNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FlightSegmentNumber(flightSegmentNumber int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.flightSegmentNumber = flightSegmentNumber
	builder.flightSegmentNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FlightSegmentTravel(flightSegmentTravel string) *BillDetailOfDomesticFlightItemBuilder {
	builder.flightSegmentTravel = flightSegmentTravel
	builder.flightSegmentTravelSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FlightTravelType(flightTravelType int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.flightTravelType = flightTravelType
	builder.flightTravelTypeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) FuelCost(fuelCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.fuelCost = fuelCost
	builder.fuelCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group1Code(group1Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group1Code = group1Code
	builder.group1CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group1Name(group1Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group1Name = group1Name
	builder.group1NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group2Code(group2Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group2Code = group2Code
	builder.group2CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group2Name(group2Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group2Name = group2Name
	builder.group2NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group3Code(group3Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group3Code = group3Code
	builder.group3CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group3Name(group3Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group3Name = group3Name
	builder.group3NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group4Code(group4Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group4Code = group4Code
	builder.group4CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group4Name(group4Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group4Name = group4Name
	builder.group4NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group5Code(group5Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group5Code = group5Code
	builder.group5CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group5Name(group5Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group5Name = group5Name
	builder.group5NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group6Code(group6Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group6Code = group6Code
	builder.group6CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group6Name(group6Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group6Name = group6Name
	builder.group6NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group7Code(group7Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group7Code = group7Code
	builder.group7CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group7Name(group7Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group7Name = group7Name
	builder.group7NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group8Code(group8Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group8Code = group8Code
	builder.group8CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group8Name(group8Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group8Name = group8Name
	builder.group8NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group9Code(group9Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group9Code = group9Code
	builder.group9CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Group9Name(group9Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.group9Name = group9Name
	builder.group9NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) InstitutionId(institutionId int64) *BillDetailOfDomesticFlightItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) InstitutionName(institutionName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) InvoiceEntityInfo(invoiceEntityInfo string) *BillDetailOfDomesticFlightItemBuilder {
	builder.invoiceEntityInfo = invoiceEntityInfo
	builder.invoiceEntityInfoSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) InvoiceFee(invoiceFee string) *BillDetailOfDomesticFlightItemBuilder {
	builder.invoiceFee = invoiceFee
	builder.invoiceFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) IsAdded(isAdded int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.isAdded = isAdded
	builder.isAddedSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) IsAgreement(isAgreement string) *BillDetailOfDomesticFlightItemBuilder {
	builder.isAgreement = isAgreement
	builder.isAgreementSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) IsDiscountTicket(isDiscountTicket string) *BillDetailOfDomesticFlightItemBuilder {
	builder.isDiscountTicket = isDiscountTicket
	builder.isDiscountTicketSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) IsSensitive(isSensitive string) *BillDetailOfDomesticFlightItemBuilder {
	builder.isSensitive = isSensitive
	builder.isSensitiveSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) IsTraveler(isTraveler string) *BillDetailOfDomesticFlightItemBuilder {
	builder.isTraveler = isTraveler
	builder.isTravelerSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) LastupdateTime(lastupdateTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.lastupdateTime = lastupdateTime
	builder.lastupdateTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) LegalEntityId(legalEntityId int64) *BillDetailOfDomesticFlightItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) LegalEntityName(legalEntityName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) MainFlightNumber(mainFlightNumber string) *BillDetailOfDomesticFlightItemBuilder {
	builder.mainFlightNumber = mainFlightNumber
	builder.mainFlightNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) MemberId(memberId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) OrderFixSettleTime(orderFixSettleTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.orderFixSettleTime = orderFixSettleTime
	builder.orderFixSettleTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) OrderFixType(orderFixType int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.orderFixType = orderFixType
	builder.orderFixTypeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) OrderId(orderId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) OriginOrderId(originOrderId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.originOrderId = originOrderId
	builder.originOrderIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) OriginTicketId(originTicketId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.originTicketId = originTicketId
	builder.originTicketIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) OutLegalEntityId(outLegalEntityId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) OutRequisitionId(outRequisitionId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.outRequisitionId = outRequisitionId
	builder.outRequisitionIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ParentInstitutionId(parentInstitutionId int64) *BillDetailOfDomesticFlightItemBuilder {
	builder.parentInstitutionId = parentInstitutionId
	builder.parentInstitutionIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ParentInstitutionName(parentInstitutionName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.parentInstitutionName = parentInstitutionName
	builder.parentInstitutionNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerDepCode(passengerDepCode string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerDepCode = passengerDepCode
	builder.passengerDepCodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerDepName(passengerDepName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerDepName = passengerDepName
	builder.passengerDepNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup1Code(passengerGroup1Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup1Code = passengerGroup1Code
	builder.passengerGroup1CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup1Name(passengerGroup1Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup1Name = passengerGroup1Name
	builder.passengerGroup1NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup2Code(passengerGroup2Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup2Code = passengerGroup2Code
	builder.passengerGroup2CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup2Name(passengerGroup2Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup2Name = passengerGroup2Name
	builder.passengerGroup2NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup3Code(passengerGroup3Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup3Code = passengerGroup3Code
	builder.passengerGroup3CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup3Name(passengerGroup3Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup3Name = passengerGroup3Name
	builder.passengerGroup3NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup4Code(passengerGroup4Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup4Code = passengerGroup4Code
	builder.passengerGroup4CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup4Name(passengerGroup4Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup4Name = passengerGroup4Name
	builder.passengerGroup4NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup5Code(passengerGroup5Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup5Code = passengerGroup5Code
	builder.passengerGroup5CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup5Name(passengerGroup5Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup5Name = passengerGroup5Name
	builder.passengerGroup5NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup6Code(passengerGroup6Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup6Code = passengerGroup6Code
	builder.passengerGroup6CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup6Name(passengerGroup6Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup6Name = passengerGroup6Name
	builder.passengerGroup6NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup7Code(passengerGroup7Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup7Code = passengerGroup7Code
	builder.passengerGroup7CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup7Name(passengerGroup7Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup7Name = passengerGroup7Name
	builder.passengerGroup7NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup8Code(passengerGroup8Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup8Code = passengerGroup8Code
	builder.passengerGroup8CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup8Name(passengerGroup8Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup8Name = passengerGroup8Name
	builder.passengerGroup8NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup9Code(passengerGroup9Code string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup9Code = passengerGroup9Code
	builder.passengerGroup9CodeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerGroup9Name(passengerGroup9Name string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerGroup9Name = passengerGroup9Name
	builder.passengerGroup9NameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerLegalEntityId(passengerLegalEntityId int64) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerLegalEntityId = passengerLegalEntityId
	builder.passengerLegalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerLegalEntityName(passengerLegalEntityName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerLegalEntityName = passengerLegalEntityName
	builder.passengerLegalEntityNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerMemberId(passengerMemberId int64) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerMemberName(passengerMemberName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerMemberName = passengerMemberName
	builder.passengerMemberNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerName(passengerName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PassengerParentPathDepName(passengerParentPathDepName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.passengerParentPathDepName = passengerParentPathDepName
	builder.passengerParentPathDepNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PayChannel(payChannel string) *BillDetailOfDomesticFlightItemBuilder {
	builder.payChannel = payChannel
	builder.payChannelSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PayType(payType string) *BillDetailOfDomesticFlightItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PersonalChangeServiceCost(personalChangeServiceCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.personalChangeServiceCost = personalChangeServiceCost
	builder.personalChangeServiceCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PersonalRealPay(personalRealPay string) *BillDetailOfDomesticFlightItemBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PersonalTicketCost(personalTicketCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.personalTicketCost = personalTicketCost
	builder.personalTicketCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PersonalUpgradeCost(personalUpgradeCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.personalUpgradeCost = personalUpgradeCost
	builder.personalUpgradeCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PnrNumber(pnrNumber string) *BillDetailOfDomesticFlightItemBuilder {
	builder.pnrNumber = pnrNumber
	builder.pnrNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) PreOrderId(preOrderId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.preOrderId = preOrderId
	builder.preOrderIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ProjectExtInfo(projectExtInfo string) *BillDetailOfDomesticFlightItemBuilder {
	builder.projectExtInfo = projectExtInfo
	builder.projectExtInfoSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RankName(rankName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.rankName = rankName
	builder.rankNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RcBook(rcBook string) *BillDetailOfDomesticFlightItemBuilder {
	builder.rcBook = rcBook
	builder.rcBookSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RcDistancePeriod(rcDistancePeriod string) *BillDetailOfDomesticFlightItemBuilder {
	builder.rcDistancePeriod = rcDistancePeriod
	builder.rcDistancePeriodSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RcLevel(rcLevel string) *BillDetailOfDomesticFlightItemBuilder {
	builder.rcLevel = rcLevel
	builder.rcLevelSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RcLowPrice(rcLowPrice string) *BillDetailOfDomesticFlightItemBuilder {
	builder.rcLowPrice = rcLowPrice
	builder.rcLowPriceSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RcTimePeriod(rcTimePeriod string) *BillDetailOfDomesticFlightItemBuilder {
	builder.rcTimePeriod = rcTimePeriod
	builder.rcTimePeriodSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Reason(reason string) *BillDetailOfDomesticFlightItemBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Rebook(rebook string) *BillDetailOfDomesticFlightItemBuilder {
	builder.rebook = rebook
	builder.rebookSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ReductionCarbonEmission(reductionCarbonEmission string) *BillDetailOfDomesticFlightItemBuilder {
	builder.reductionCarbonEmission = reductionCarbonEmission
	builder.reductionCarbonEmissionSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ReductionFee(reductionFee float64) *BillDetailOfDomesticFlightItemBuilder {
	builder.reductionFee = reductionFee
	builder.reductionFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RefundHandleCost(refundHandleCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.refundHandleCost = refundHandleCost
	builder.refundHandleCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) Remark(remark string) *BillDetailOfDomesticFlightItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) RequisitionId(requisitionId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.requisitionId = requisitionId
	builder.requisitionIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) ServiceFee(serviceFee string) *BillDetailOfDomesticFlightItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) SettleTime(settleTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.settleTime = settleTime
	builder.settleTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) SettlementTime(settlementTime string) *BillDetailOfDomesticFlightItemBuilder {
	builder.settlementTime = settlementTime
	builder.settlementTimeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) SubAccountCompanyName(subAccountCompanyName string) *BillDetailOfDomesticFlightItemBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) SubCompanyNo(subCompanyNo string) *BillDetailOfDomesticFlightItemBuilder {
	builder.subCompanyNo = subCompanyNo
	builder.subCompanyNoSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) SupplierOrderId(supplierOrderId string) *BillDetailOfDomesticFlightItemBuilder {
	builder.supplierOrderId = supplierOrderId
	builder.supplierOrderIdSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) SupplierTicketNumber(supplierTicketNumber string) *BillDetailOfDomesticFlightItemBuilder {
	builder.supplierTicketNumber = supplierTicketNumber
	builder.supplierTicketNumberSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TicketFee(ticketFee string) *BillDetailOfDomesticFlightItemBuilder {
	builder.ticketFee = ticketFee
	builder.ticketFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TicketIdText(ticketIdText string) *BillDetailOfDomesticFlightItemBuilder {
	builder.ticketIdText = ticketIdText
	builder.ticketIdTextSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TotalFee(totalFee string) *BillDetailOfDomesticFlightItemBuilder {
	builder.totalFee = totalFee
	builder.totalFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TrainComparePrice(trainComparePrice float64) *BillDetailOfDomesticFlightItemBuilder {
	builder.trainComparePrice = trainComparePrice
	builder.trainComparePriceSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TransactionType(transactionType int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.transactionType = transactionType
	builder.transactionTypeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TravelItineraryFee(travelItineraryFee string) *BillDetailOfDomesticFlightItemBuilder {
	builder.travelItineraryFee = travelItineraryFee
	builder.travelItineraryFeeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TravelPurpose(travelPurpose string) *BillDetailOfDomesticFlightItemBuilder {
	builder.travelPurpose = travelPurpose
	builder.travelPurposeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TripDescription(tripDescription string) *BillDetailOfDomesticFlightItemBuilder {
	builder.tripDescription = tripDescription
	builder.tripDescriptionSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) TripReason(tripReason string) *BillDetailOfDomesticFlightItemBuilder {
	builder.tripReason = tripReason
	builder.tripReasonSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) UniqueKey(uniqueKey string) *BillDetailOfDomesticFlightItemBuilder {
	builder.uniqueKey = uniqueKey
	builder.uniqueKeySet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) UnusualContent(unusualContent string) *BillDetailOfDomesticFlightItemBuilder {
	builder.unusualContent = unusualContent
	builder.unusualContentSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) UnusualReason(unusualReason string) *BillDetailOfDomesticFlightItemBuilder {
	builder.unusualReason = unusualReason
	builder.unusualReasonSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) UnusualRemark(unusualRemark string) *BillDetailOfDomesticFlightItemBuilder {
	builder.unusualRemark = unusualRemark
	builder.unusualRemarkSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) UnusualType(unusualType string) *BillDetailOfDomesticFlightItemBuilder {
	builder.unusualType = unusualType
	builder.unusualTypeSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) UpgradeCost(upgradeCost string) *BillDetailOfDomesticFlightItemBuilder {
	builder.upgradeCost = upgradeCost
	builder.upgradeCostSet = true
	return builder
}
func (builder *BillDetailOfDomesticFlightItemBuilder) UpgradeFinalStatus(upgradeFinalStatus int32) *BillDetailOfDomesticFlightItemBuilder {
	builder.upgradeFinalStatus = upgradeFinalStatus
	builder.upgradeFinalStatusSet = true
	return builder
}

func (builder *BillDetailOfDomesticFlightItemBuilder) Build() *BillDetailOfDomesticFlightItem {
	data := &BillDetailOfDomesticFlightItem{}
	if builder.addedCutReasonSet {
		data.AddedCutReason = &builder.addedCutReason
	}
	if builder.addedEsCutFeeSet {
		data.AddedEsCutFee = &builder.addedEsCutFee
	}
	if builder.addedGoodsNameSet {
		data.AddedGoodsName = &builder.addedGoodsName
	}
	if builder.adjustmentSet {
		data.Adjustment = &builder.adjustment
	}
	if builder.airlineSimpleNameSet {
		data.AirlineSimpleName = &builder.airlineSimpleName
	}
	if builder.applyChangeTimeSet {
		data.ApplyChangeTime = &builder.applyChangeTime
	}
	if builder.applyRefundTimeSet {
		data.ApplyRefundTime = &builder.applyRefundTime
	}
	if builder.approvalChangeHistorySet {
		data.ApprovalChangeHistory = &builder.approvalChangeHistory
	}
	if builder.approvalFirstSet {
		data.ApprovalFirst = &builder.approvalFirst
	}
	if builder.approvalHistorySet {
		data.ApprovalHistory = &builder.approvalHistory
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.approvalNormalHistorySet {
		data.ApprovalNormalHistory = &builder.approvalNormalHistory
	}
	if builder.approvalSecondSet {
		data.ApprovalSecond = &builder.approvalSecond
	}
	if builder.approvalThirdSet {
		data.ApprovalThird = &builder.approvalThird
	}
	if builder.arrivalAirportCodeSet {
		data.ArrivalAirportCode = &builder.arrivalAirportCode
	}
	if builder.arrivalAirportNameSet {
		data.ArrivalAirportName = &builder.arrivalAirportName
	}
	if builder.arrivalCityIdSet {
		data.ArrivalCityId = &builder.arrivalCityId
	}
	if builder.arrivalCityNameSet {
		data.ArrivalCityName = &builder.arrivalCityName
	}
	if builder.arrivalTimeSet {
		data.ArrivalTime = &builder.arrivalTime
	}
	if builder.baseFeeSet {
		data.BaseFee = &builder.baseFee
	}
	if builder.beforeCutServiceFeeSet {
		data.BeforeCutServiceFee = &builder.beforeCutServiceFee
	}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	if builder.bookingDateSet {
		data.BookingDate = &builder.bookingDate
	}
	if builder.bookingDepCodeSet {
		data.BookingDepCode = &builder.bookingDepCode
	}
	if builder.bookingDepNameSet {
		data.BookingDepName = &builder.bookingDepName
	}
	if builder.bookingMemberIdSet {
		data.BookingMemberId = &builder.bookingMemberId
	}
	if builder.bookingMemberNameSet {
		data.BookingMemberName = &builder.bookingMemberName
	}
	if builder.bookingMemberNumberSet {
		data.BookingMemberNumber = &builder.bookingMemberNumber
	}
	if builder.bookingParentPathDepNameSet {
		data.BookingParentPathDepName = &builder.bookingParentPathDepName
	}
	if builder.budgetCenterCodeSet {
		data.BudgetCenterCode = &builder.budgetCenterCode
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetCenterParentPathNameSet {
		data.BudgetCenterParentPathName = &builder.budgetCenterParentPathName
	}
	if builder.cabinCodeSet {
		data.CabinCode = &builder.cabinCode
	}
	if builder.cabinNameSet {
		data.CabinName = &builder.cabinName
	}
	if builder.cabinTypeSet {
		data.CabinType = &builder.cabinType
	}
	if builder.changeHandleCostSet {
		data.ChangeHandleCost = &builder.changeHandleCost
	}
	if builder.companyChangeServiceCostSet {
		data.CompanyChangeServiceCost = &builder.companyChangeServiceCost
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.companyTicketCostSet {
		data.CompanyTicketCost = &builder.companyTicketCost
	}
	if builder.companyUpgradeCostSet {
		data.CompanyUpgradeCost = &builder.companyUpgradeCost
	}
	if builder.constructionCostSet {
		data.ConstructionCost = &builder.constructionCost
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.currentCarbonEmissionSet {
		data.CurrentCarbonEmission = &builder.currentCarbonEmission
	}
	if builder.daysInAdvanceSet {
		data.DaysInAdvance = &builder.daysInAdvance
	}
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.departmentIdSet {
		data.DepartmentId = &builder.departmentId
	}
	if builder.departureAirportCodeSet {
		data.DepartureAirportCode = &builder.departureAirportCode
	}
	if builder.departureAirportNameSet {
		data.DepartureAirportName = &builder.departureAirportName
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.departureCityNameSet {
		data.DepartureCityName = &builder.departureCityName
	}
	if builder.departureTimeSet {
		data.DepartureTime = &builder.departureTime
	}
	if builder.discountSet {
		data.Discount = &builder.discount
	}
	if builder.economyOriginalCostSet {
		data.EconomyOriginalCost = &builder.economyOriginalCost
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.exInfo01Set {
		data.ExInfo01 = &builder.exInfo01
	}
	if builder.exInfo01CodeSet {
		data.ExInfo01Code = &builder.exInfo01Code
	}
	if builder.exInfo02Set {
		data.ExInfo02 = &builder.exInfo02
	}
	if builder.exInfo02CodeSet {
		data.ExInfo02Code = &builder.exInfo02Code
	}
	if builder.exInfo03Set {
		data.ExInfo03 = &builder.exInfo03
	}
	if builder.exInfo03CodeSet {
		data.ExInfo03Code = &builder.exInfo03Code
	}
	if builder.exInfo04Set {
		data.ExInfo04 = &builder.exInfo04
	}
	if builder.exInfo04CodeSet {
		data.ExInfo04Code = &builder.exInfo04Code
	}
	if builder.exInfo05Set {
		data.ExInfo05 = &builder.exInfo05
	}
	if builder.exInfo05CodeSet {
		data.ExInfo05Code = &builder.exInfo05Code
	}
	if builder.exInfo06Set {
		data.ExInfo06 = &builder.exInfo06
	}
	if builder.exInfo06CodeSet {
		data.ExInfo06Code = &builder.exInfo06Code
	}
	if builder.exInfo07Set {
		data.ExInfo07 = &builder.exInfo07
	}
	if builder.exInfo07CodeSet {
		data.ExInfo07Code = &builder.exInfo07Code
	}
	if builder.exInfo08Set {
		data.ExInfo08 = &builder.exInfo08
	}
	if builder.exInfo08CodeSet {
		data.ExInfo08Code = &builder.exInfo08Code
	}
	if builder.exceptionInfoSet {
		data.ExceptionInfo = &builder.exceptionInfo
	}
	if builder.excludingServiceFeeSet {
		data.ExcludingServiceFee = &builder.excludingServiceFee
	}
	if builder.exemptReasonSet {
		data.ExemptReason = &builder.exemptReason
	}
	if builder.exemptTypeSet {
		data.ExemptType = &builder.exemptType
	}
	if builder.extendInfoSet {
		data.ExtendInfo = &builder.extendInfo
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.finalDepartureTimeSet {
		data.FinalDepartureTime = &builder.finalDepartureTime
	}
	if builder.finalOfflineStatusSet {
		data.FinalOfflineStatus = &builder.finalOfflineStatus
	}
	if builder.flightDistanceSet {
		data.FlightDistance = &builder.flightDistance
	}
	if builder.flightNumberSet {
		data.FlightNumber = &builder.flightNumber
	}
	if builder.flightSegmentNumberSet {
		data.FlightSegmentNumber = &builder.flightSegmentNumber
	}
	if builder.flightSegmentTravelSet {
		data.FlightSegmentTravel = &builder.flightSegmentTravel
	}
	if builder.flightTravelTypeSet {
		data.FlightTravelType = &builder.flightTravelType
	}
	if builder.fuelCostSet {
		data.FuelCost = &builder.fuelCost
	}
	if builder.group1CodeSet {
		data.Group1Code = &builder.group1Code
	}
	if builder.group1NameSet {
		data.Group1Name = &builder.group1Name
	}
	if builder.group2CodeSet {
		data.Group2Code = &builder.group2Code
	}
	if builder.group2NameSet {
		data.Group2Name = &builder.group2Name
	}
	if builder.group3CodeSet {
		data.Group3Code = &builder.group3Code
	}
	if builder.group3NameSet {
		data.Group3Name = &builder.group3Name
	}
	if builder.group4CodeSet {
		data.Group4Code = &builder.group4Code
	}
	if builder.group4NameSet {
		data.Group4Name = &builder.group4Name
	}
	if builder.group5CodeSet {
		data.Group5Code = &builder.group5Code
	}
	if builder.group5NameSet {
		data.Group5Name = &builder.group5Name
	}
	if builder.group6CodeSet {
		data.Group6Code = &builder.group6Code
	}
	if builder.group6NameSet {
		data.Group6Name = &builder.group6Name
	}
	if builder.group7CodeSet {
		data.Group7Code = &builder.group7Code
	}
	if builder.group7NameSet {
		data.Group7Name = &builder.group7Name
	}
	if builder.group8CodeSet {
		data.Group8Code = &builder.group8Code
	}
	if builder.group8NameSet {
		data.Group8Name = &builder.group8Name
	}
	if builder.group9CodeSet {
		data.Group9Code = &builder.group9Code
	}
	if builder.group9NameSet {
		data.Group9Name = &builder.group9Name
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.institutionNameSet {
		data.InstitutionName = &builder.institutionName
	}
	if builder.invoiceEntityInfoSet {
		data.InvoiceEntityInfo = &builder.invoiceEntityInfo
	}
	if builder.invoiceFeeSet {
		data.InvoiceFee = &builder.invoiceFee
	}
	if builder.isAddedSet {
		data.IsAdded = &builder.isAdded
	}
	if builder.isAgreementSet {
		data.IsAgreement = &builder.isAgreement
	}
	if builder.isDiscountTicketSet {
		data.IsDiscountTicket = &builder.isDiscountTicket
	}
	if builder.isSensitiveSet {
		data.IsSensitive = &builder.isSensitive
	}
	if builder.isTravelerSet {
		data.IsTraveler = &builder.isTraveler
	}
	if builder.lastupdateTimeSet {
		data.LastupdateTime = &builder.lastupdateTime
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.legalEntityNameSet {
		data.LegalEntityName = &builder.legalEntityName
	}
	if builder.mainFlightNumberSet {
		data.MainFlightNumber = &builder.mainFlightNumber
	}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.orderFixSettleTimeSet {
		data.OrderFixSettleTime = &builder.orderFixSettleTime
	}
	if builder.orderFixTypeSet {
		data.OrderFixType = &builder.orderFixType
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.originOrderIdSet {
		data.OriginOrderId = &builder.originOrderId
	}
	if builder.originTicketIdSet {
		data.OriginTicketId = &builder.originTicketId
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.outRequisitionIdSet {
		data.OutRequisitionId = &builder.outRequisitionId
	}
	if builder.parentInstitutionIdSet {
		data.ParentInstitutionId = &builder.parentInstitutionId
	}
	if builder.parentInstitutionNameSet {
		data.ParentInstitutionName = &builder.parentInstitutionName
	}
	if builder.passengerDepCodeSet {
		data.PassengerDepCode = &builder.passengerDepCode
	}
	if builder.passengerDepNameSet {
		data.PassengerDepName = &builder.passengerDepName
	}
	if builder.passengerGroup1CodeSet {
		data.PassengerGroup1Code = &builder.passengerGroup1Code
	}
	if builder.passengerGroup1NameSet {
		data.PassengerGroup1Name = &builder.passengerGroup1Name
	}
	if builder.passengerGroup2CodeSet {
		data.PassengerGroup2Code = &builder.passengerGroup2Code
	}
	if builder.passengerGroup2NameSet {
		data.PassengerGroup2Name = &builder.passengerGroup2Name
	}
	if builder.passengerGroup3CodeSet {
		data.PassengerGroup3Code = &builder.passengerGroup3Code
	}
	if builder.passengerGroup3NameSet {
		data.PassengerGroup3Name = &builder.passengerGroup3Name
	}
	if builder.passengerGroup4CodeSet {
		data.PassengerGroup4Code = &builder.passengerGroup4Code
	}
	if builder.passengerGroup4NameSet {
		data.PassengerGroup4Name = &builder.passengerGroup4Name
	}
	if builder.passengerGroup5CodeSet {
		data.PassengerGroup5Code = &builder.passengerGroup5Code
	}
	if builder.passengerGroup5NameSet {
		data.PassengerGroup5Name = &builder.passengerGroup5Name
	}
	if builder.passengerGroup6CodeSet {
		data.PassengerGroup6Code = &builder.passengerGroup6Code
	}
	if builder.passengerGroup6NameSet {
		data.PassengerGroup6Name = &builder.passengerGroup6Name
	}
	if builder.passengerGroup7CodeSet {
		data.PassengerGroup7Code = &builder.passengerGroup7Code
	}
	if builder.passengerGroup7NameSet {
		data.PassengerGroup7Name = &builder.passengerGroup7Name
	}
	if builder.passengerGroup8CodeSet {
		data.PassengerGroup8Code = &builder.passengerGroup8Code
	}
	if builder.passengerGroup8NameSet {
		data.PassengerGroup8Name = &builder.passengerGroup8Name
	}
	if builder.passengerGroup9CodeSet {
		data.PassengerGroup9Code = &builder.passengerGroup9Code
	}
	if builder.passengerGroup9NameSet {
		data.PassengerGroup9Name = &builder.passengerGroup9Name
	}
	if builder.passengerLegalEntityIdSet {
		data.PassengerLegalEntityId = &builder.passengerLegalEntityId
	}
	if builder.passengerLegalEntityNameSet {
		data.PassengerLegalEntityName = &builder.passengerLegalEntityName
	}
	if builder.passengerMemberIdSet {
		data.PassengerMemberId = &builder.passengerMemberId
	}
	if builder.passengerMemberNameSet {
		data.PassengerMemberName = &builder.passengerMemberName
	}
	if builder.passengerMemberNumberSet {
		data.PassengerMemberNumber = &builder.passengerMemberNumber
	}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.passengerParentPathDepNameSet {
		data.PassengerParentPathDepName = &builder.passengerParentPathDepName
	}
	if builder.payChannelSet {
		data.PayChannel = &builder.payChannel
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.personalChangeServiceCostSet {
		data.PersonalChangeServiceCost = &builder.personalChangeServiceCost
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.personalTicketCostSet {
		data.PersonalTicketCost = &builder.personalTicketCost
	}
	if builder.personalUpgradeCostSet {
		data.PersonalUpgradeCost = &builder.personalUpgradeCost
	}
	if builder.pnrNumberSet {
		data.PnrNumber = &builder.pnrNumber
	}
	if builder.preOrderIdSet {
		data.PreOrderId = &builder.preOrderId
	}
	if builder.projectExtInfoSet {
		data.ProjectExtInfo = &builder.projectExtInfo
	}
	if builder.rankNameSet {
		data.RankName = &builder.rankName
	}
	if builder.rcBookSet {
		data.RcBook = &builder.rcBook
	}
	if builder.rcDistancePeriodSet {
		data.RcDistancePeriod = &builder.rcDistancePeriod
	}
	if builder.rcLevelSet {
		data.RcLevel = &builder.rcLevel
	}
	if builder.rcLowPriceSet {
		data.RcLowPrice = &builder.rcLowPrice
	}
	if builder.rcTimePeriodSet {
		data.RcTimePeriod = &builder.rcTimePeriod
	}
	if builder.reasonSet {
		data.Reason = &builder.reason
	}
	if builder.rebookSet {
		data.Rebook = &builder.rebook
	}
	if builder.reductionCarbonEmissionSet {
		data.ReductionCarbonEmission = &builder.reductionCarbonEmission
	}
	if builder.reductionFeeSet {
		data.ReductionFee = &builder.reductionFee
	}
	if builder.refundHandleCostSet {
		data.RefundHandleCost = &builder.refundHandleCost
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	if builder.requisitionIdSet {
		data.RequisitionId = &builder.requisitionId
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.settleTimeSet {
		data.SettleTime = &builder.settleTime
	}
	if builder.settlementTimeSet {
		data.SettlementTime = &builder.settlementTime
	}
	if builder.subAccountCompanyNameSet {
		data.SubAccountCompanyName = &builder.subAccountCompanyName
	}
	if builder.subCompanyNoSet {
		data.SubCompanyNo = &builder.subCompanyNo
	}
	if builder.supplierOrderIdSet {
		data.SupplierOrderId = &builder.supplierOrderId
	}
	if builder.supplierTicketNumberSet {
		data.SupplierTicketNumber = &builder.supplierTicketNumber
	}
	if builder.ticketFeeSet {
		data.TicketFee = &builder.ticketFee
	}
	if builder.ticketIdTextSet {
		data.TicketIdText = &builder.ticketIdText
	}
	if builder.totalFeeSet {
		data.TotalFee = &builder.totalFee
	}
	if builder.trainComparePriceSet {
		data.TrainComparePrice = &builder.trainComparePrice
	}
	if builder.transactionTypeSet {
		data.TransactionType = &builder.transactionType
	}
	if builder.travelItineraryFeeSet {
		data.TravelItineraryFee = &builder.travelItineraryFee
	}
	if builder.travelPurposeSet {
		data.TravelPurpose = &builder.travelPurpose
	}
	if builder.tripDescriptionSet {
		data.TripDescription = &builder.tripDescription
	}
	if builder.tripReasonSet {
		data.TripReason = &builder.tripReason
	}
	if builder.uniqueKeySet {
		data.UniqueKey = &builder.uniqueKey
	}
	if builder.unusualContentSet {
		data.UnusualContent = &builder.unusualContent
	}
	if builder.unusualReasonSet {
		data.UnusualReason = &builder.unusualReason
	}
	if builder.unusualRemarkSet {
		data.UnusualRemark = &builder.unusualRemark
	}
	if builder.unusualTypeSet {
		data.UnusualType = &builder.unusualType
	}
	if builder.upgradeCostSet {
		data.UpgradeCost = &builder.upgradeCost
	}
	if builder.upgradeFinalStatusSet {
		data.UpgradeFinalStatus = &builder.upgradeFinalStatus
	}
	return data
}
