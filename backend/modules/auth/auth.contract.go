package auth

type authSignUpRequest struct {
	Name        string `json:"name" xml:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phonenumber" xml:"phonenumber" form:"phonenumber" validate:"required"`
	Password    string `json:"password" xml:"password" form:"password" validate:"required"`
}

type authVerificationRequest struct {
	Code        string `json:"code" xml:"code" form:"code" validate:"required"`
	PhoneNumber string `json:"phonenumber" xml:"phonenumber" form:"phonenumber" validate:"required"`
}

type authSignInRequest struct {
	PhoneNumber string `json:"phonenumber" xml:"phonenumber" form:"phonenumber" validate:"required"`
	Password    string `json:"password" xml:"password" form:"password" validate:"required"`
}

type authRequestResetPasswordRequest struct {
	PhoneNumber string `json:"phonenumber" xml:"phonenumber" form:"phonenumber" validate:"required"`
}

type authResetPasswordRequest struct {
	Code        string `json:"code" xml:"code" form:"code" validate:"required"`
	PhoneNumber string `json:"phonenumber" xml:"phonenumber" form:"code" validate:"required"`
	Password    string `json:"password" xml:"password" form:"password" validate:"password"`
}

type authResponse struct {
	Payload interface{} `json:"payload"`
}
