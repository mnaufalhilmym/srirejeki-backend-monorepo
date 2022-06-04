package profile

import (
	"greenhouse-monitoring-iot/internal/pkg/cryptograph"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"

	"github.com/gofiber/fiber/v2"
)

type returnedProfileServiceData struct {
	code    int
	payload interface{}
}

func (p *Profile) getProfileService(userSession *model.UserSession) *returnedProfileServiceData {
	user, err := p.getUserDataByPhoneNumber(&userSession.PhoneNumber)
	if err != nil {
		errorHandler.LogErrorThenContinue("getProfileService1", err)
		return &returnedProfileServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedProfileServiceData{code: fiber.StatusOK, payload: &user}
}

func (p *Profile) editProfileService(bearer *string, userSession *model.UserSession, profileEditRequest *profileEditRequest) *returnedProfileServiceData {
	toBeUpdatedUser := model.User{Name: profileEditRequest.Name, PhoneNumber: profileEditRequest.PhoneNumber}
	toBeUpdatedUserSession := model.UserSession{}
	if len(profileEditRequest.Name) > 0 {
		toBeUpdatedUserSession.Name = profileEditRequest.Name
	}
	if len(profileEditRequest.PhoneNumber) > 0 {
		toBeUpdatedUserSession.PhoneNumber = profileEditRequest.PhoneNumber
	}
	if len(profileEditRequest.Password) > 0 {
		encryptedPassword, err := cryptograph.Encrypt(profileEditRequest.Password)
		if err != nil {
			errorHandler.LogErrorThenContinue("editProfileService1", err)
			return &returnedProfileServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
		}
		toBeUpdatedUser.Password = *encryptedPassword
	}
	if err := p.updateUserDataByPhoneNumber(&userSession.PhoneNumber, &toBeUpdatedUser); err != nil {
		errorHandler.LogErrorThenContinue("editProfileService2", err)
		return &returnedProfileServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	if err := p.updateUserDataByPhoneNumberSession(bearer, userSession, &toBeUpdatedUserSession); err != nil {
		errorHandler.LogErrorThenContinue("profileServicePatch3", err)
		return &returnedProfileServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedProfileServiceData{code: fiber.StatusOK, payload: nil}
}
