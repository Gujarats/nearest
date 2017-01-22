package review

type Review struct {
	date    string `db:"date"`
	shop_id string `db:"shop_id"`
}
