package reviewInstance

import (
	"log"

	"github.com/training_project/model/review"
	"github.com/training_project/model/review/interface"
	"github.com/training_project/model/review/mock"
)

var ReviewInstance reviewInterface.ReviewInterface
var InputReviewStruct interface{}

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
		// TODO : use an error print here.
		// because we don't need others case.
		// if non the case above meet then execute fatal error.
		log.Fatalln("Error Pass struct")
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
		log.Fatalln("Error Review Instance Null")
	}
	return ReviewInstance
}
