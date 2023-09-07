package base

type Model struct {
	ID         int64  `json:"id"`
	GmtCreate  string `json:"gmtCreate"`
	GmtModify  string `json:"gmtModify"`
	CreateUser string `json:"createUser"`
	ModifyUser string `json:"modifyUser"`
}
