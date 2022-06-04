package farmland

import (
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type returnedFarmlandServiceData struct {
	code    int
	payload interface{}
}

func (f *Farmland) getFarmlandsService(userSession *model.UserSession) *returnedFarmlandServiceData {
	farmlands, err := f.getFarmlandsByUserID(&userSession.ID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getFarmlandsService1", err)
		return &returnedFarmlandServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedFarmlandServiceData{code: fiber.StatusOK, payload: &farmlands}
}

func (f *Farmland) getFarmlandService(farmlandGetRequest *farmlandGetRequest) *returnedFarmlandServiceData {
	farmlandID64, err := strconv.ParseUint(farmlandGetRequest.ID, 10, 64)
	if err != nil {
		errorHandler.LogErrorThenContinue("getFarmlandService1", err)
		return &returnedFarmlandServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	farmlandID := uint(farmlandID64)
	farmland, err := f.getFarmlandByID(&farmlandID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getFarmlandService2", err)
		return &returnedFarmlandServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedFarmlandServiceData{code: fiber.StatusOK, payload: &farmland}
}

func (f *Farmland) postFarmlandService(userSession *model.UserSession, farmlandAddRequest *farmlandAddRequest) *returnedFarmlandServiceData {
	farmland := model.Farmland{Name: farmlandAddRequest.Name, Description: farmlandAddRequest.Description, Location: farmlandAddRequest.Location, UserID: userSession.ID}
	if err := f.createFarmland(&farmland); err != nil {
		errorHandler.LogErrorThenContinue("postFarmlandService1", err)
		return &returnedFarmlandServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedFarmlandServiceData{code: fiber.StatusCreated, payload: nil}
}

func (f *Farmland) patchFarmlandService(userSession *model.UserSession, farmlandEditRequest *farmlandEditRequest) *returnedFarmlandServiceData {
	toBeUpdatedFarmland := model.Farmland{Name: farmlandEditRequest.Name, Description: farmlandEditRequest.Description, Location: farmlandEditRequest.Location, UserID: userSession.ID}
	if err := f.updateFarmlandByID(&farmlandEditRequest.ID, &toBeUpdatedFarmland, &userSession.ID); err != nil {
		errorHandler.LogErrorThenContinue("patchFarmlandService1", err)
		return &returnedFarmlandServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedFarmlandServiceData{code: fiber.StatusOK, payload: nil}
}

func (f *Farmland) deleteFarmlandService(farmlandDeleteRequest *farmlandDeleteRequest) *returnedFarmlandServiceData {
	if err := f.deleteMicrocontrollerByFarmlandID(&farmlandDeleteRequest.ID); err != nil {
		errorHandler.LogErrorThenContinue("deleteFarmlandService1", err)
		return &returnedFarmlandServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	if err := f.deleteFarmlandByID(&farmlandDeleteRequest.ID); err != nil {
		errorHandler.LogErrorThenContinue("deleteFarmlandService2", err)
		return &returnedFarmlandServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedFarmlandServiceData{code: fiber.StatusOK, payload: nil}
}
