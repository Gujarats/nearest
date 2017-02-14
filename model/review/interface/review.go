package reviewMethod

import "github.com/jmoiron/sqlx"

type ReviewInterface interface {
	GetConn(*sqlx.DB)
	Exist() bool
}

type ReviewDataMock struct {
	Message string
	UserID  int64
	ShopID  int64
}

func (r *ReviewDataMock) GetConn(connection *sqlx.DB) {

}

// return true for mocking data exist
func (r *ReviewDataMock) Exist() bool {
	return true
}
