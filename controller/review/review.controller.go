package review

import (
	"github.com/training_project/model/review/interface"
)

//checking if the data exist in the table ws_product_feedback
func IsDataExist(review reviewMethod.ReviewInterface) bool {
	return review.Exist()
}
