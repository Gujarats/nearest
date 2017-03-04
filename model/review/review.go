package review

import (
	"fmt"

	"github.com/Gujarats/API-Golang/model/global"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	ReviewResponse struct {
		global.Response // use global.Response for consistence response
	}

	ReviewRequest struct {
		UserID string `json:"user_id"`
		ShopID string `json:"shop_id"`
	}

	ReviewData struct {
		Message string
		UserID  int64
		ShopID  int64
	}
)

var postgres *sqlx.DB

func (r *ReviewData) GetConn(connection *sqlx.DB) {
	postgres = connection
}

func (r *ReviewData) Exist() bool {
	// TODO :  checking the shopID value is nil or not.
	// it is a good way to check the value before go to the execution
	query := fmt.Sprintf(
		`
		SELECT
		'x'
		FROM
		ws_active_seller
		WHERE
		shop_id = %d
		`,
		r.ShopID,
	)

	result := postgres.QueryRow(query)
	var data []uint8
	result.Scan(&data)

	if len(data) < 1 {
		return false
	}

	if data[0] != 'x' {
		return false
	}

	return true
}
