package reviewStruct

import (
	"github.com/training_project/model/review"
	"github.com/training_project/model/review/interface"
)

var ReviewObject reviewMethod.ReviewInterface
var InputReviewStruct interface{}

func SetStruct(inputReviewStruct interface{}) {
	InputReviewStruct = inputReviewStruct
	switch inputReviewStruct.(type) {
	case review.ReviewData:
		var reviewData review.ReviewData
		reviewData = inputReviewStruct.(review.ReviewData)
		ReviewObject = &reviewData
		break
	case reviewMethod.ReviewDataMock:
		var reviewData reviewMethod.ReviewDataMock
		reviewData = inputReviewStruct.(reviewMethod.ReviewDataMock)
		ReviewObject = &reviewData
		break
	default:
		var reviewData review.ReviewData
		reviewData = inputReviewStruct.(review.ReviewData)
		ReviewObject = &reviewData
		break
	}
}

func PassParams(shopID int64) {
	//checking if we setup InputReviewStruct
	if InputReviewStruct == nil {
		reviewData := review.ReviewData{}
		reviewData.ShopID = shopID
		ReviewObject = &reviewData
	} else {
		reviewDataMock := reviewMethod.ReviewDataMock{}
		reviewDataMock.ShopID = shopID
		ReviewObject = &reviewDataMock

	}

}

func GetStruct() reviewMethod.ReviewInterface {
	return ReviewObject
}
