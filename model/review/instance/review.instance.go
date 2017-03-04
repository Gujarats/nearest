package reviewInstance

import (
	"github.com/Gujarats/API-Golang/model/review"
	"github.com/Gujarats/API-Golang/model/review/interface"
	"github.com/Gujarats/API-Golang/model/review/mock"
	"github.com/Gujarats/API-Golang/util/logger"
)

var ReviewInstance reviewInterface.ReviewInterface
var InputReviewStruct interface{}

func init() {
	// setup logger location and prin the bugs and specific location
	logger.InitLogger("Review Instance :: ", "../../logs/", "reviewInstace.txt")
}

func SetInstance(inputReviewStruct interface{}) {
	InputReviewStruct = inputReviewStruct
	switch inputReviewStruct.(type) {
	case review.ReviewData:
		reviewData := inputReviewStruct.(review.ReviewData)
		ReviewInstance = &reviewData
		break
	case reviewMock.ReviewDataMock:
		reviewData := inputReviewStruct.(reviewMock.ReviewDataMock)
		ReviewInstance = &reviewData
		break
	default:
		logger.FatalLog("Error Pass struct")
		break
	}
}

func PassParams(shopID int64) {
	// pass the params if Review Instance is null.
	// because as unit test define the ReviewInstance first
	if ReviewInstance == nil {
		reviewData := review.ReviewData{ShopID: shopID}
		ReviewInstance = &reviewData
	}
}

func GetReviewInstance() reviewInterface.ReviewInterface {
	if ReviewInstance == nil {
		logger.FatalLog("Error Review Instance Null")
	}
	return ReviewInstance
}
