package mcu

import (
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	mqttclient "greenhouse-monitoring-iot/internal/pkg/mqttClient"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type returnedMicrocontrollerServiceData struct {
	code    int
	payload interface{}
}

func (m *Microcontroller) getAllMicrocontrollersService(userSession *model.UserSession) *returnedMicrocontrollerServiceData {
	microcontrollers, err := m.getMicrocontrollersByUserID(&userSession.ID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getAllMicrocontrollersService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusOK, payload: &microcontrollers}
}

func (m *Microcontroller) getFarmlandMicrocontrollersService(microcontrollersFarmlandGetRequest *microcontrollersFarmlandGetRequest) *returnedMicrocontrollerServiceData {
	farmlandID64, err := strconv.ParseUint(microcontrollersFarmlandGetRequest.ID, 10, 64)
	if err != nil {
		errorHandler.LogErrorThenContinue("getFarmlandMicrocontrollersService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	farmlandID := uint(farmlandID64)
	microcontrollers, err := m.getMicrocontrollersByFarmlandID(&farmlandID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getFarmlandMicrocontrollersService2", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusOK, payload: &microcontrollers}
}

func (m *Microcontroller) getMicrocontrollerService(microcontrollerGetRequest *microcontrollerGetRequest) *returnedMicrocontrollerServiceData {
	farmlandID64, err := strconv.ParseUint(microcontrollerGetRequest.ID, 10, 64)
	if err != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollerService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	farmlandID := uint(farmlandID64)
	microcontroller, err := m.getMicrocontrollerByID(&farmlandID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getMicrocontrollerService2", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}

	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusOK, payload: &microcontroller}
}

func (m *Microcontroller) postMicrocontrollerService(userSession *model.UserSession, microcontrollerAddRequest *microcontrollerAddRequest) *returnedMicrocontrollerServiceData {
	_, err := m.getFarmlandByID(&microcontrollerAddRequest.FarmlandID)
	if err != nil {
		errorHandler.LogErrorThenContinue("postMicrocontrollerService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	microcontrollerData, err := m.getMicrocontrollerByDeviceID(&microcontrollerAddRequest.DeviceID)
	if err == nil {
		if microcontrollerData.UserID == userSession.ID {
			err = errors.New("device has been registered")
		} else {
			err = errors.New("device has been registered by another user")
		}
		errorHandler.LogErrorThenContinue("postMicrocontrollerService2", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	microcontroller := model.Microcontroller{Name: microcontrollerAddRequest.Name, Description: microcontrollerAddRequest.Description, Location: microcontrollerAddRequest.Location, DeviceID: microcontrollerAddRequest.DeviceID, FarmlandID: microcontrollerAddRequest.FarmlandID, UserID: userSession.ID}
	if err := m.createMicrocontroller(&microcontroller); err != nil {
		errorHandler.LogErrorThenContinue("postMicrocontrollerService3", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusCreated, payload: nil}
}

func (m *Microcontroller) patchMicrocontrollerService(microcontrollerEditRequest *microcontrollerEditRequest) *returnedMicrocontrollerServiceData {
	toBeUpdatedMicrocontroller := model.Microcontroller{Name: microcontrollerEditRequest.Name, Description: microcontrollerEditRequest.Description, Location: microcontrollerEditRequest.Location, DeviceID: microcontrollerEditRequest.DeviceID}
	if err := m.updateMicrocontrollerByID(&microcontrollerEditRequest.ID, &toBeUpdatedMicrocontroller); err != nil {
		errorHandler.LogErrorThenContinue("patchMicrocontrollerService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusOK, payload: nil}
}

func (m *Microcontroller) deleteMicrocontrollerService(microcontrollerDeleteRequest *microcontrollerDeleteRequest) *returnedMicrocontrollerServiceData {
	if err := m.deleteMicrocontrollerByID(&microcontrollerDeleteRequest.ID); err != nil {
		errorHandler.LogErrorThenContinue("deleteMicrocontrollerService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusOK, payload: nil}
}

func (m *Microcontroller) authMicrocontrollerService(microcontrollerAuthRequest *microcontrollerAuthRequest) *returnedMicrocontrollerServiceData {
	_, err := m.getMicrocontrollerByDeviceID(&microcontrollerAuthRequest.DeviceID)
	if err != nil {
		errorHandler.LogErrorThenContinue("authMicrocontrollerService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusOK, payload: nil}
}

func (m *Microcontroller) sendDataToMcuService(userSession *model.UserSession, microcontrollerSendDataToMcuRequest *microcontrollerSendDataToMcuRequest) *returnedMicrocontrollerServiceData {
	microcontroller, err := m.getMicrocontrollerByDeviceID(&microcontrollerSendDataToMcuRequest.DeviceID)
	if err != nil {
		errorHandler.LogErrorThenContinue("forwardDataToMcuService1", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	if microcontroller.UserID != userSession.ID {
		err := errors.New("not authorized")
		errorHandler.LogErrorThenContinue("forwardDataToMcuService2", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	num := 1
	qos := 2
	topic := "srirejeki/server/" + microcontroller.DeviceID + microcontrollerSendDataToMcuRequest.Type
	if err := mqttclient.Pub(&num, &qos, &topic, &microcontrollerSendDataToMcuRequest.Data); err != nil {
		errorHandler.LogErrorThenContinue("forwardDataToMcuService3", err)
		return &returnedMicrocontrollerServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedMicrocontrollerServiceData{code: fiber.StatusOK, payload: nil}
}
