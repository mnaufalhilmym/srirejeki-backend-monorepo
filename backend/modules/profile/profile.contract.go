package profile

type profileEditRequest struct {
	Name        string `json:"name" xml:"name" form:"name"`
	PhoneNumber string `json:"phonenumber" xml:"phonenumber" form:"phonenumber"`
	Password    string `json:"password" xml:"password" form:"password"`
}

type profileResponse struct {
	Payload interface{} `json:"payload"`
}
