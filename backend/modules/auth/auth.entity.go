package auth

import (
	"encoding/json"
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/database"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"greenhouse-monitoring-iot/internal/pkg/random"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Auth struct {
	FiberRouter fiber.Router
	PgDB        *gorm.DB
	RDB         *redis.Client
}

func (a *Auth) createRegistrantSession(registrant *model.Registrant) error {
	val, err := json.Marshal(registrant)
	if err != nil {
		errorHandler.LogErrorThenContinue("createRegistrant1", err)
		return err
	}
	if err := a.RDB.Set("registration-"+registrant.VerificationCode, val, 5*time.Minute).Err(); err != nil {
		errorHandler.LogErrorThenContinue("createRegistrant2", err)
		return err
	}
	return nil
}

func (a *Auth) createUser(user *model.User) error {
	result := a.PgDB.Create(&user)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("createUser1", result.Error)
		return result.Error
	}
	return nil
}

func (a *Auth) createUserResetPasswordSession(userResetPassword *model.UserResetPassword) error {
	val, err := json.Marshal(userResetPassword)
	if err != nil {
		errorHandler.LogErrorThenContinue("createUserResetPassword1", err)
		return err
	}
	if err := a.RDB.Set("resetpassword-"+userResetPassword.VerificationCode, val, 5*time.Minute).Err(); err != nil {
		errorHandler.LogErrorThenContinue("createUserResetPassword2", err)
		return err
	}
	return nil
}

func (a *Auth) createAuthBearerTokenSession(userSession *model.UserSession) (*string, error) {
	authBearerToken := ""
	for {
		authBearerToken = random.RandStringBytesMaskImprSrcUnsafe(172)
		val, err := a.RDB.Exists("user-" + authBearerToken).Result()
		if err != nil {
			errorHandler.LogErrorThenContinue("authBearer/Generate1", err)
			return nil, err
		}
		if val == 0 {
			break
		}
	}
	val, err := json.Marshal(userSession)
	if err != nil {
		errorHandler.LogErrorThenContinue("createUserResetPassword1", err)
		return nil, err
	}
	if err := a.RDB.Set("user-"+authBearerToken, val, 3*24*time.Hour).Err(); err != nil {
		errorHandler.LogErrorThenContinue("createUserResetPassword2", err)
		return nil, err
	}
	return &authBearerToken, nil
}

func (a *Auth) updateUserPhoneNumberByPhoneNumber(phonenumber *string, toBeUpdatedUser *model.User) error {
	result := a.PgDB.Model(&model.User{}).Where("phone_number", &phonenumber).Updates(&toBeUpdatedUser)
	if result.Error != nil {
		errorHandler.LogErrorThenContinue("updateUserByPhoneNumber1", result.Error)
		return result.Error
	}
	return nil
}

func (a *Auth) getRegistrantDataByVerificationCodeSession(verificationCode *string) (*model.Registrant, error) {
	var registrant *model.Registrant
	val, err := a.RDB.Get("registration-" + *verificationCode).Bytes()
	if err != nil {
		errorHandler.LogErrorThenContinue("getRegistrantDataByVerificationCode1", err)
		return nil, err
	}
	if err := json.Unmarshal(val, &registrant); err != nil {
		errorHandler.LogErrorThenContinue("getRegistrantDataByVerificationCode2", err)
		return nil, err
	}
	return registrant, nil
}

func (a *Auth) getUserDataByPhoneNumber(phoneNumber *string) (*model.User, error) {
	user, err := database.GetUserDataByPhoneNumber(a.PgDB, phoneNumber)
	if err != nil {
		errorHandler.LogErrorThenContinue("getUserDataByPhoneNumber1", err)
		return nil, err
	}
	return user, nil
}

func (a *Auth) getUserResetPasswordByVerificationCodeSession(verificationCode *string) (*model.UserResetPassword, error) {
	var userResetPassword *model.UserResetPassword
	val, err := a.RDB.Get("resetpassword-" + *verificationCode).Bytes()
	if err != nil {
		errorHandler.LogErrorThenContinue("getUserResetPasswordByVerificationCode1", err)
		return nil, err
	}
	if err := json.Unmarshal(val, &userResetPassword); err != nil {
		errorHandler.LogErrorThenContinue("getUserResetPasswordByVerificationCode2", err)
		return nil, err
	}
	return userResetPassword, nil
}

func (a *Auth) deleteRegistrantDataByVerificationCodeSession(verificationCode *string) error {
	val, err := a.RDB.Del("registration-" + *verificationCode).Result()
	if err != nil {
		errorHandler.LogErrorThenContinue("deleteRegistrantDataByVerificationCode1", err)
		return err
	}
	if val == 0 {
		err := errors.New("failed delete registrant data")
		errorHandler.LogErrorThenContinue("deleteRegistrantDataByVerificationCode2", err)
		return err
	}
	return nil
}

func (a *Auth) deleteUserResetPasswordByVerificationCodeSession(verificationCode *string) error {
	val, err := a.RDB.Del("resetpassword-" + *verificationCode).Result()
	if err != nil {
		errorHandler.LogErrorThenContinue("deleteUserResetPasswordByVerificationCode1", err)
		return err
	}
	if val == 0 {
		err := errors.New("failed delete user reset password data")
		errorHandler.LogErrorThenContinue("deleteUserResetPasswordByVerificationCode2", err)
		return err
	}
	return nil
}

func (a *Auth) deleteAuthBearerTokenSession(authBearerToken string) error {
	val, err := a.RDB.Del("user-" + authBearerToken).Result()
	if err != nil {
		errorHandler.LogErrorThenContinue("deleteAuthBearerToken1", err)
		return err
	}
	if val == 0 {
		err := errors.New("authorization Bearer code not found")
		errorHandler.LogErrorThenContinue("deleteAuthBearerToken2", err)
		return err
	}
	return nil
}
