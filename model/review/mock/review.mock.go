package reviewMock

import "github.com/jmoiron/sqlx"

type ReviewDataMock struct {
	Message     string
	UserID      int64
	ShopID      int64
	IsDataExist bool
}

// below is bunch of the mehod that mocking the real result

// return true for mocking data exist
func (r *ReviewDataMock) Exist() bool {
	return r.IsDataExist
}

// We create this method to satisfy the interface method
func (r *ReviewDataMock) GetConn(connection *sqlx.DB) {

}
