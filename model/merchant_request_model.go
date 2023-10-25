package model

type MerchantRequestModel struct {
	ID       string
	Username string
	Password string
	User_id  string
	Role     string
	Active   bool
	FullName string
	NoPhone  string
	Email    string
	Address  string
}
