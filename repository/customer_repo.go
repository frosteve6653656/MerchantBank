package repository

import (
	"database/sql"
	"final-project/model"
	"final-project/utils"
	"fmt"
)

type CustomerRepo interface {
	InsertCustomer(cust *model.CustomerRequestModel) error
	GetCustomerById(id string) (*model.CustomerModel, error)
	GetCustomerByName(name string) (*model.CustomerModel, error)
	GetAllCustomer() ([]*model.CustomerModel, error)
	EditCustomerById(cust model.CustomerModel) error
	GetCustomerByUserId(id string) (*model.CustomerModel, error)
}

type customerRepoImpl struct {
	db *sql.DB
}

func (customerRepo *customerRepoImpl) InsertCustomer(cust *model.CustomerRequestModel) error {
	tx, err := customerRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.InsertCustomer() 1 : %w", err)
	}

	qryUser := utils.INSERT_CUST
	err = tx.QueryRow(qryUser, utils.UuidGenerate(), cust.Username, cust.Password, cust.Role, cust.Active).Scan(&cust.User_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on customerRepoImpl.InsertCustomer() 2  : %w", err)
	}

	qry := utils.INSERT_CUST_USR
	_, err = tx.Exec(qry, utils.UuidGenerate(), cust.User_id, cust.FullName, cust.NoPhone, cust.Email, cust.Address)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on customerRepoImpl.InsertCustomer() 3 : %w", err)
	}

	tx.Commit()
	return nil
}

func (customerRepo *customerRepoImpl) GetCustomerById(id string) (*model.CustomerModel, error) {
	qry := utils.GET_CUST_ID
	cust := &model.CustomerModel{}
	err := customerRepo.db.QueryRow(qry, id).Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NoPhone, &cust.Email, &cust.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetCustomerById() : %w", err)
	}
	return cust, nil
}

func (customerRepo *customerRepoImpl) GetCustomerByUserId(id string) (*model.CustomerModel, error) {
	qry := utils.GET_CUST_USRID
	cust := &model.CustomerModel{}
	err := customerRepo.db.QueryRow(qry, id).Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NoPhone, &cust.Email, &cust.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetCustomerById() : %w", err)
	}
	return cust, nil
}

func (customerRepo *customerRepoImpl) GetAllCustomer() ([]*model.CustomerModel, error) {
	qry := utils.GET_ALL_CUSTOMER
	rows, err := customerRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on customerRepoImpl.GetAllCustomer() : %w", err)
	}
	defer rows.Close()
	var arrCustomer []*model.CustomerModel
	for rows.Next() {
		cust := &model.CustomerModel{}
		rows.Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NoPhone, &cust.Email, &cust.Address)
		arrCustomer = append(arrCustomer, cust)
	}
	return arrCustomer, nil
}

func (customerRepo *customerRepoImpl) GetCustomerByName(name string) (*model.CustomerModel, error) {
	qry := utils.GET_CUST_NAME
	cust := &model.CustomerModel{}
	err := customerRepo.db.QueryRow(qry, name).Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NoPhone, &cust.Email, &cust.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetCustomerByName() : %w", err)
	}
	return cust, nil
}

func (customerRepo *customerRepoImpl) EditCustomerById(cust model.CustomerModel) error {
	qry := utils.EDIT_CUST_ID
	_, err := customerRepo.db.Exec(qry, cust.FullName, cust.NoPhone, cust.Email, cust.Address, cust.ID)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.EditCustomerById() 3 : %w", err)
	}

	return nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepoImpl{
		db: db,
	}
}
