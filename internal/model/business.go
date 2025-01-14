package model

type Business struct {
	RegisterName string `db:"REGISTER_NAME"`
	Name         string `db:"BN_NAME"`
	Status       string `db:"BN_STATUS"`
	RegisterDate string `db:"BN_REG_DT"` // eg: 18/10/2024
	CancelDate   string `db:"BN_CANCEL_DT"`
	RenewDate    string `db:"BN_RENEW_DT"`
	StateNumber  string `db:"BN_STATE_NUM"`
	StateOfReg   string `db:"BN_STATE_OF_REG"`
	ABN          string `db:"BN_ABN"`
}

type BusinessSearch struct {
	Name  string `db:"name"`
	ABN   string `db:"abn"`
	State string `db:"state"`
}
