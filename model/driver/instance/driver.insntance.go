package driverInstance

import (
	"github.com/training_project/model/driver"
	"github.com/training_project/model/driver/interface"
	"github.com/training_project/model/driver/mock"
	"github.com/training_project/util/logger"
)

var DriverInstance driverInterface.DriverInterfacce
var InputReviewStruct interface{}

func init() {
	// setup logger location and prin the bugs and specific location
	logger.InitLogger("Driver Instance :: ", "../../logs/", "reviewInstace.txt")
}

func SetInstance(inputReviewStruct interface{}) {
	InputReviewStruct = inputReviewStruct
	switch inputReviewStruct.(type) {
	case driver.DriverData:
		driverData := inputReviewStruct.(driver.DriverData)
		DriverInstance = &driverData
		break
	case driverMock.DriverDataMock:
		driverData := inputReviewStruct.(driverMock.DriverDataMock)
		DriverInstance = &driverData
		break
	default:
		logger.FatalLog("Error Pass struct")
		break
	}
}

func GetInstance() driverInterface.DriverInterfacce {
	// if DriverInstance is nil then pass the real driver
	if DriverInstance == nil {
		DriverInstance = &driver.DriverData{}
		return DriverInstance
	}

	return DriverInstance
}
