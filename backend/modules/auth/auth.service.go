package auth

import (
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/cryptograph"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"greenhouse-monitoring-iot/internal/pkg/random"
	"greenhouse-monitoring-iot/internal/pkg/waba"

	"github.com/gofiber/fiber/v2"
)

type returnedAuthServiceData struct {
	code    int
	payload interface{}
}

func (a *Auth) postSignUpService(authSignUpRequest *authSignUpRequest) *returnedAuthServiceData {
	_, err := a.getUserDataByPhoneNumber(&authSignUpRequest.PhoneNumber)
	if err == nil {
		err := errors.New("user is exist")
		errorHandler.LogErrorThenContinue("postSignUpService1", err)
		return &returnedAuthServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	encryptedPassword, err := cryptograph.Encrypt(authSignUpRequest.Password)
	if err != nil {
		errorHandler.LogErrorThenContinue("postSignUpService2", err)
		return &returnedAuthServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	verificationCode := random.RandomRangeNumber(100000, 999999)
	registrant := model.Registrant{VerificationCode: verificationCode, Name: authSignUpRequest.Name, PhoneNumber: authSignUpRequest.PhoneNumber, Password: *encryptedPassword}
	if err := a.createRegistrantSession(&registrant); err != nil {
		errorHandler.LogErrorThenContinue("postSignUpService3", err)
		return &returnedAuthServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	go func() {
		if err := waba.SendMessage(registrant.PhoneNumber, verificationCode+" adalah kode verifikasi SriRejeki. Mohon tidak membagikan kode verifikasi ini kepada pihak lain."); err != nil {
			errorHandler.LogErrorThenContinue("postSignUpService4", err)
			// Pass
		}
	}()
	return &returnedAuthServiceData{code: fiber.StatusCreated, payload: &verificationCode}
}

func (a *Auth) postVerifySignUpService(authVerificationRequest *authVerificationRequest) *returnedAuthServiceData {
	registrant, err := a.getRegistrantDataByVerificationCodeSession(&authVerificationRequest.Code)
	if err != nil {
		errorHandler.LogErrorThenContinue("postVerifySignUpService1", err)
		return &returnedAuthServiceData{code: fiber.StatusForbidden, payload: err.Error()}
	}
	if registrant.PhoneNumber != authVerificationRequest.PhoneNumber {
		errorHandler.LogErrorThenContinue("postVerifySignUpService2", err)
		return &returnedAuthServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	user := model.User{Name: registrant.Name, PhoneNumber: registrant.PhoneNumber, Password: registrant.Password}
	if err := a.createUser(&user); err != nil {
		errorHandler.LogErrorThenContinue("postVerifySignUpService3", err)
		return &returnedAuthServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	go func() {
		if err := a.deleteRegistrantDataByVerificationCodeSession(&authVerificationRequest.Code); err != nil {
			errorHandler.LogErrorThenContinue("postVerifySignUpService4", err)
			// Pass
		}
	}()
	return &returnedAuthServiceData{code: fiber.StatusCreated, payload: nil}
}

func (a *Auth) postSignInService(authSignInRequest *authSignInRequest) *returnedAuthServiceData {
	user, err := a.getUserDataByPhoneNumber(&authSignInRequest.PhoneNumber)
	if err != nil {
		errorHandler.LogErrorThenContinue("postSignInService1", err)
		return &returnedAuthServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	decryptedPass, err := cryptograph.Decrypt(user.Password)
	if err != nil {
		errorHandler.LogErrorThenContinue("postSignInService2", err)
		return &returnedAuthServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	if user.PhoneNumber != authSignInRequest.PhoneNumber || *decryptedPass != authSignInRequest.Password {
		err := errors.New("wrong credential")
		errorHandler.LogErrorThenContinue("postSignInService3", err)
		return &returnedAuthServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	authBearerToken, err := a.createAuthBearerTokenSession(&model.UserSession{ID: user.ID, Name: user.Name, PhoneNumber: user.PhoneNumber})
	if err != nil {
		errorHandler.LogErrorThenContinue("postSignInService4", err)
		return &returnedAuthServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedAuthServiceData{code: fiber.StatusOK, payload: &authBearerToken}
}

func (a *Auth) postRequestResetPasswordService(authRequestResetPasswordRequest *authRequestResetPasswordRequest) *returnedAuthServiceData {
	user, err := a.getUserDataByPhoneNumber(&authRequestResetPasswordRequest.PhoneNumber)
	if err != nil {
		errorHandler.LogErrorThenContinue("postRequestResetPasswordService1", err)
		return &returnedAuthServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	verificationCode := random.RandomRangeNumber(100000, 999999)
	if err := a.createUserResetPasswordSession(&model.UserResetPassword{VerificationCode: verificationCode, PhoneNumber: user.PhoneNumber}); err != nil {
		errorHandler.LogErrorThenContinue("postRequestResetPasswordService2", err)
		return &returnedAuthServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	go func() {
		if err := waba.SendMessage(user.PhoneNumber, verificationCode); err != nil {
			errorHandler.LogErrorThenContinue("postRequestResetPasswordService3", err)
			// Pass
		}
	}()
	return &returnedAuthServiceData{code: fiber.StatusOK, payload: &verificationCode}
}

func (a *Auth) postVerifyRequestResetPasswordService(authVerificationRequest *authVerificationRequest) *returnedAuthServiceData {
	userResetPassword, err := a.getUserResetPasswordByVerificationCodeSession(&authVerificationRequest.Code)
	if err != nil {
		errorHandler.LogErrorThenContinue("postVerifyRequestResetPasswordService1", err)
		return &returnedAuthServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	if userResetPassword.PhoneNumber != authVerificationRequest.PhoneNumber {
		err := errors.New("request not valid")
		errorHandler.LogErrorThenContinue("postVerifyRequestResetPasswordService2", err)
		return &returnedAuthServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	return &returnedAuthServiceData{code: fiber.StatusOK, payload: nil}
}

func (a *Auth) patchResetPasswordService(authResetPasswordRequest *authResetPasswordRequest) *returnedAuthServiceData {
	result := a.postVerifyRequestResetPasswordService(&authVerificationRequest{Code: authResetPasswordRequest.Code, PhoneNumber: authResetPasswordRequest.PhoneNumber})
	if result.code != fiber.StatusOK {
		err := errors.New(result.payload.(string))
		errorHandler.LogErrorThenContinue("patchResetPasswordService1", err)
		return &returnedAuthServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	encryptedPassword, err := cryptograph.Encrypt(authResetPasswordRequest.Password)
	if err != nil {
		errorHandler.LogErrorThenContinue("patchResetPasswordService2", err)
		return &returnedAuthServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	if err := a.updateUserPhoneNumberByPhoneNumber(&authResetPasswordRequest.PhoneNumber, &model.User{Password: *encryptedPassword}); err != nil {
		errorHandler.LogErrorThenContinue("patchResetPasswordService3", err)
		return &returnedAuthServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	go func() {
		if err := a.deleteUserResetPasswordByVerificationCodeSession(&authResetPasswordRequest.Code); err != nil {
			errorHandler.LogErrorThenContinue("patchResetPasswordService3", err)
			// Pass
		}
	}()
	return &returnedAuthServiceData{code: fiber.StatusOK, payload: nil}
}

func (a *Auth) getSignOutService(authBearerToken string) *returnedAuthServiceData {
	go func() {
		if err := a.deleteAuthBearerTokenSession(authBearerToken); err != nil {
			errorHandler.LogErrorThenContinue("getSignOutService1", err)
			// Pass
		}
	}()
	return &returnedAuthServiceData{code: fiber.StatusOK, payload: nil}
}
