package data

type dataSubscribeRequest struct {
	Topic string `query:"topic" validate:"required"`
	Limit string `query:"limit"`
}

type dataSnapshotGetRequest struct {
	Type     string `query:"type" validate:"required"`
	DeviceID string `query:"deviceId" validate:"required"`
	Duration string `query:"duration"`
	Limit    string `query:"limit"`
}

type dataSnapshotPostRequest struct {
	Type      string   `json:"type" xml:"type" form:"type" validate:"required"`
	Data      string   `json:"data" xml:"data" form:"data" validate:"required"`
	DeviceID  string   `json:"deviceId" xml:"deviceId" form:"deviceId" validate:"required"`
	Durations []string `json:"durations" xml:"durations" form:"durations" validate:"required"`
}

type dataResponse struct {
	Payload interface{} `json:"payload"`
}
