package model

type ActiveSeller struct {
	Date              string `db:"date"`
	TotalActiveSeller string `db:"active_seller"`
}
