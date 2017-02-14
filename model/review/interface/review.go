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

// below is bunch of the mehod that mocking the real result

// return true for mocking data exist
func (r *ReviewDataMock) Exist() bool {
	return true
}

// Get something cool
func (r *ReviewDataMock) GetConn(connection *sqlx.DB) {

}
