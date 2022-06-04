package mcu

type microcontrollersFarmlandGetRequest struct {
	ID string `query:"id" validate:"required"`
}

type microcontrollerGetRequest struct {
	ID string `query:"id" validate:"required"`
}

type microcontrollerAddRequest struct {
	Name        string `json:"name" xml:"name" form:"name" validate:"required"`
	Description string `json:"description" xml:"description" form:"description"`
	Location    string `json:"location" xml:"location" form:"location"`
	FarmlandID  uint   `json:"farmlandId" xml:"farmlandId" form:"farmlandId" validate:"required"`
	DeviceID    string `json:"deviceId" xml:"deviceId" form:"deviceId" validate:"required"`
}

type microcontrollerEditRequest struct {
	ID          uint   `json:"id" xml:"id" form:"id" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name"`
	Description string `json:"description" xml:"description" form:"description"`
	Location    string `json:"location" xml:"location" form:"location"`
	DeviceID    string `json:"deviceId" xml:"deviceId" form:"deviceId"`
}

type microcontrollerDeleteRequest struct {
	ID uint `json:"id" xml:"id" form:"id" validate:"required"`
}

type microcontrollerAuthRequest struct {
	DeviceID string `json:"deviceId" xml:"deviceId" form:"deviceId" validate:"required"`
}

type microcontrollerSendDataToMcuRequest struct {
	DeviceID string `json:"deviceId" xml:"deviceId" form:"deviceId" validate:"required"`
	Type     string `json:"type" xml:"type" form:"type" validate:"required"`
	Data     string `json:"data" xml:"data" form:"data" validate:"required"`
}

type microcontrollerResponse struct {
	Payload interface{} `json:"payload"`
}
