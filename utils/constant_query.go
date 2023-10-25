package utils

const (
	// USER
	INSERT_USER      = "INSERT INTO ms_user(id,username, password, role, is_active ) VALUES($1, $2, $3, $4, $5)"
	GET_USER_BY_ID   = "SELECT id, username,role,is_active FROM ms_user WHERE id = $1"
	GET_ALL_USER     = "SELECT id, username,role,is_active FROM ms_user"
	GET_USER_BY_NAME = "SELECT id, username,password,role,is_active FROM ms_user WHERE username = $1"
	EDIT_USER_ID     = "UPDATE ms_user SET username=$1, password=$2, is_active=$3 WHERE id = $4"

	//customer
	INSERT_CUST      = "INSERT INTO ms_user(id,username, password, role, is_active ) VALUES($1, $2, $3, $4, $5) RETURNING id"
	INSERT_CUST_USR  = "INSERT INTO ms_customer (id, id_user, full_name, noPhone, email, address) VALUES ($1, $2, $3, $4, $5, $6);"
	GET_CUST_ID      = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_customer WHERE id=$1"
	GET_CUST_USRID   = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_customer WHERE id_user = $1"
	GET_ALL_CUSTOMER = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_customer"
	GET_CUST_NAME    = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_customer WHERE full_name = $1"
	EDIT_CUST_ID     = "UPDATE ms_customer SET full_name=$1,noPhone=$2,email=$3,address=$4 WHERE id = $5"

	//merchant
	INSERT_MERC      = "INSERT INTO ms_user(id,username, password, role, is_active ) VALUES($1, $2, $3, $4, $5) RETURNING id"
	INSERT_MERC_USR  = "INSERT INTO ms_merchant (id, id_user, full_name, noPhone, email, address) VALUES ($1, $2, $3, $4, $5, $6);"
	GET_MERC_ID      = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_merchant WHERE id=$1"
	GET_MERC_USRID   = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_merchant WHERE id_user = $1"
	GET_ALL_MERCHANT = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_merchant"
	GET_MERC_NAME    = "SELECT id,id_user,full_name,noPhone,email,address FROM ms_merchant WHERE full_name = $1"
	EDIT_MERC_ID     = "UPDATE ms_merchant SET full_name=$1,noPhone=$2,email=$3,address=$4 WHERE id = $5"

	//bank
	INSERT_BANK    = "INSERT INTO ms_bank(id,id_user,wallet) VALUES($1, $2, $3)"
	GET_BANK_ID    = "SELECT id,id_user,wallet FROM ms_bank WHERE id=$1"
	GET_BANK_USRID = "SELECT id,id_user,wallet FROM ms_bank WHERE id_user = $1"
	GET_ALL_BANK   = "SELECT id,id_user,wallet FROM ms_bank"
	EDIT_BANK_ID   = "UPDATE ms_bank SET wallet=$1 WHERE id = $2"

	//transfer
	INSERT_TRANSFER    = "INSERT INTO tx_transfer(id,id_send,id_receive,wallet) VALUES($1, $2, $3)"
	GET_TRANSFER_ID    = "SELECT id,id_user,wallet FROM tx_transfer WHERE id=$1"
	GET_TRANSFER_USRID = "SELECT id,id_user,wallet FROM tx_transfer WHERE id_user = $1"
	GET_ALL_TRANSFER   = "SELECT id,id_user,wallet FROM tx_transfer"
)
