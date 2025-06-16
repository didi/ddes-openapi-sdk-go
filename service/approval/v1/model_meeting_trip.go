package v1

// MeetingTrip 会议行程信息(TravelDetail中meeting_trip信息)
type MeetingTrip struct {
	MeetingCity []BusinessCity `json:"meeting_city,omitempty"` // 会议城市，会议管控必传
}

type MeetingTripBuilder struct {
	meetingCity    []BusinessCity // 会议城市，会议管控必传
	meetingCitySet bool
}

func NewMeetingTripBuilder() *MeetingTripBuilder {
	return &MeetingTripBuilder{}
}
func (builder *MeetingTripBuilder) MeetingCity(meetingCity []BusinessCity) *MeetingTripBuilder {
	builder.meetingCity = meetingCity
	builder.meetingCitySet = true
	return builder
}

func (builder *MeetingTripBuilder) Build() *MeetingTrip {
	data := &MeetingTrip{}
	if builder.meetingCitySet {
		data.MeetingCity = builder.meetingCity
	}
	return data
}
