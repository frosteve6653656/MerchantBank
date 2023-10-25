package repository

import (
	"database/sql"
	"final-project/model"
	"final-project/utils"
	"fmt"
)

type BankRepo interface {
	InsertBank(bank *model.BankModel) error
	GetBankById(id string) (*model.BankModel, error)
	GetAllBank() ([]*model.BankModel, error)
	EditBankById(bank model.BankModel) error
	GetBankByUserId(id string) (*model.BankModel, error)
}

type bankRepoImpl struct {
	db *sql.DB
}

func (bankRepo *bankRepoImpl) InsertBank(bank *model.BankModel) error {

	qryUser := utils.INSERT_BANK
	err := bankRepo.db.QueryRow(qryUser, utils.UuidGenerate(), bank.User_id, bank.Wallet)
	if err != nil {
		return fmt.Errorf("error on bankRepoImpl.InsertBank() 1  : %v", err)
	}

	return nil
}

func (bankRepo *bankRepoImpl) GetBankById(id string) (*model.BankModel, error) {
	qry := utils.GET_BANK_ID
	bank := &model.BankModel{}
	err := bankRepo.db.QueryRow(qry, id).Scan(&bank.ID, &bank.User_id, &bank.Wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on bankRepoImpl.GetBankById() : %w", err)
	}
	return bank, nil
}

func (bankRepo *bankRepoImpl) GetBankByUserId(id string) (*model.BankModel, error) {
	qry := utils.GET_BANK_USRID
	bank := &model.BankModel{}
	err := bankRepo.db.QueryRow(qry, id).Scan(&bank.ID, &bank.User_id, &bank.Wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on bankRepoImpl.GetBankById() : %w", err)
	}
	return bank, nil
}

func (bankRepo *bankRepoImpl) GetAllBank() ([]*model.BankModel, error) {
	qry := utils.GET_ALL_BANK
	rows, err := bankRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on bankRepoImpl.GetAllBank() : %w", err)
	}
	defer rows.Close()
	var arrBank []*model.BankModel
	for rows.Next() {
		bank := &model.BankModel{}
		rows.Scan(&bank.ID, &bank.User_id, &bank.Wallet)
		arrBank = append(arrBank, bank)
	}
	return arrBank, nil
}

func (bankRepo *bankRepoImpl) EditBankById(bank model.BankModel) error {
	qry := utils.EDIT_BANK_ID
	_, err := bankRepo.db.Exec(qry, bank.Wallet, bank.ID)
	if err != nil {
		return fmt.Errorf("error on bankRepoImpl.EditBankById() 3 : %w", err)
	}

	return nil
}

func NewBankRepo(db *sql.DB) BankRepo {
	return &bankRepoImpl{
		db: db,
	}
}
