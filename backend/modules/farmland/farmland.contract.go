package farmland

type farmlandGetRequest struct {
	ID string `query:"id" validate:"required"`
}

type farmlandAddRequest struct {
	Name        string `json:"name" xml:"name" form:"name" validate:"required"`
	Description string `json:"description" xml:"description" form:"description"`
	Location    string `json:"location" xml:"location" form:"location"`
}

type farmlandEditRequest struct {
	ID          uint   `json:"id" xml:"id" form:"id" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name"`
	Description string `json:"description" xml:"description" form:"description"`
	Location    string `json:"location" xml:"location" form:"location"`
}

type farmlandDeleteRequest struct {
	ID uint `json:"id" xml:"id" form:"id" validate:"required"`
}

type farmlandResponse struct {
	Payload interface{} `json:"payload"`
}
