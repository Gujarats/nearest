package reviewInterface

import "github.com/jmoiron/sqlx"

type ReviewInterface interface {
	GetConn(*sqlx.DB)
	Exist() bool
}
