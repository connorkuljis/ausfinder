package model

import "database/sql"

type Business struct {
	ABN          sql.NullString `db:"abn"`
	Name         sql.NullString `db:"name"`
	Status       sql.NullString `db:"status"`
	RegisteredAt sql.NullString `db:"registered_at"` // eg: 18/10/2024
	CancelAt     sql.NullString `db:"cancel_at"`
	RenewAt      sql.NullString `db:"renew_at"`
	StateNumber  sql.NullString `db:"state_num"`
	StateOfReg   sql.NullString `db:"state_of_reg"`
}

type BusinessSearch struct {
	Name  sql.NullString `db:"name"`
	ABN   sql.NullString `db:"abn"`
	State sql.NullString `db:"state"`
}
