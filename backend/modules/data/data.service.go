package data

import (
	"errors"
	"greenhouse-monitoring-iot/internal/pkg/database/model"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	mqttclient "greenhouse-monitoring-iot/internal/pkg/mqttClient"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type returnedDataServiceData struct {
	code    int
	payload interface{}
}

func (d *Data) getSubscribeService(userSession *model.UserSession, dataSubscribeRequest *dataSubscribeRequest) *returnedDataServiceData {
	topicSplitted := strings.Split(dataSubscribeRequest.Topic, "/")
	if len(topicSplitted) != 4 {
		err := errors.New("invalid topic format")
		errorHandler.LogErrorThenContinue("getSubscribeService1", err)
		return &returnedDataServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
	}
	deviceID := topicSplitted[2]
	microcontroller, err := d.getMicrocontrollerByDeviceID(&deviceID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getSubscribeService2", err)
		return &returnedDataServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	if microcontroller.UserID != userSession.ID {
		err := errors.New("user unauthorized")
		errorHandler.LogErrorThenContinue("getSubscribeService4", err)
		return &returnedDataServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	num := 1
	qos := 0
	msgArr, err := mqttclient.Sub(&num, &qos, &dataSubscribeRequest.Topic)
	if err != nil {
		errorHandler.LogErrorThenContinue("getSubscribeService5", err)
		return &returnedDataServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedDataServiceData{code: fiber.StatusOK, payload: &msgArr}
}

func (d *Data) getSnapshotSevice(userSession *model.UserSession, dataSnapshotGetRequest *dataSnapshotGetRequest) *returnedDataServiceData {
	microcontroller, err := d.getMicrocontrollerByDeviceID(&dataSnapshotGetRequest.DeviceID)
	if err != nil {
		errorHandler.LogErrorThenContinue("getSnapshotSevice1", err)
		return &returnedDataServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	if microcontroller.UserID != userSession.ID {
		err := errors.New("user unauthorized")
		errorHandler.LogErrorThenContinue("getSnapshotSevice2", err)
		return &returnedDataServiceData{code: fiber.StatusUnauthorized, payload: err.Error()}
	}
	duration := dataSnapshotGetRequest.Duration
	switch duration {
	case "hour", "day", "month":
		// Pass
	default:
		duration = "month"
	}
	limit, err := strconv.Atoi(dataSnapshotGetRequest.Limit)
	if err != nil {
		errorHandler.LogErrorThenContinue("getSnapshotSevice3", err)
		limit = 30
	}
	snapshots, err := d.getSnapshotsByDeviceIDAndDurationAndTypeWithLimit(&dataSnapshotGetRequest.DeviceID, &duration, &dataSnapshotGetRequest.Type, &limit)
	if err != nil {
		errorHandler.LogErrorThenContinue("getSnapshotSevice4", err)
		return &returnedDataServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedDataServiceData{code: fiber.StatusCreated, payload: &snapshots}
}

func (d *Data) postSnapshotService(dataSnapshotPostRequest *dataSnapshotPostRequest) *returnedDataServiceData {
	microcontroller, err := d.getMicrocontrollerByDeviceID(&dataSnapshotPostRequest.DeviceID)
	if err != nil {
		errorHandler.LogErrorThenContinue("postSnapshotService1", err)
		return &returnedDataServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	var durations []string
	for _, duration := range dataSnapshotPostRequest.Durations {
		switch duration {
		case "hour", "day", "month":
			durations = append(durations, duration)
		default:
			err := errors.New("wrong duration format")
			errorHandler.LogErrorThenContinue("postSnapshotService2", err)
			return &returnedDataServiceData{code: fiber.StatusBadRequest, payload: err.Error()}
		}
	}
	snapshot := model.Snapshot{Type: dataSnapshotPostRequest.Type, Data: dataSnapshotPostRequest.Data, DeviceID: dataSnapshotPostRequest.DeviceID, Durations: durations, MicrocontrollerID: microcontroller.ID}
	if err := d.createSnapshot(&snapshot); err != nil {
		errorHandler.LogErrorThenContinue("postSnapshotService3", err)
		return &returnedDataServiceData{code: fiber.StatusInternalServerError, payload: err.Error()}
	}
	return &returnedDataServiceData{code: fiber.StatusCreated, payload: nil}
}
