package model

import "database/sql"

type Business struct {
	ABN                 string         `db:"abn"`
	Name                string         `db:"name"`
	Status              sql.NullString `db:"status"`
	RegisteredAt        sql.NullString `db:"registered_at"` // eg: 18/10/2024
	CancelAt            sql.NullString `db:"cancel_at"`
	RenewAt             sql.NullString `db:"renew_at"`
	StateNumber         sql.NullString `db:"state_number"`
	StateOfRegistration sql.NullString `db:"state_of_registration"`
}

type BusinessSearch struct {
	Name  string         `db:"name"`
	ABN   string         `db:"abn"`
	State sql.NullString `db:"state"`
}
