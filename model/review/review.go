package review

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	reviewResponse struct {
		Status  string
		Message string
		Data    interface{}
	}

	reviewRequest struct {
		UserID string `json:"user_id"`
		ShopID string `json:"shop_id"`
	}
)

var postgres *sqlx.DB

func GetConn(connection *sqlx.DB) {
	postgres = connection
}

//checking if the data exist in the table ws_product_feedback
func IsDataExist(date string, shopID int64) bool {
	// get slave db connection

	query := fmt.Sprintf(
		`
		SELECT
		'x'
		FROM
		ws_active_seller
		WHERE
		shop_id = %d
		`,
		shopID,
	)

	result := postgres.QueryRow(query)
	fmt.Printf("Postgre Query Row = %+v\n", result)

	var data []uint8
	result.Scan(&data)

	fmt.Printf("Postgre Scan = %+v\n", data)

	if data[0] != 'x' {
		return false
	}

	return true
}
