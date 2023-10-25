package repository

import (
	"database/sql"
	"final-project/model"
	"final-project/utils"
	"fmt"
)

type MerchantRepo interface {
	InsertMerchant(merc *model.MerchantRequestModel) error
	GetMerchantById(id string) (*model.MerchantModel, error)
	GetMerchantByName(name string) (*model.MerchantModel, error)
	GetAllMerchant() ([]*model.MerchantModel, error)
	EditMerchantById(merc model.MerchantModel) error
	GetMerchantByUserId(id string) (*model.MerchantModel, error)
}

type merchantRepoImpl struct {
	db *sql.DB
}

func (merchantRepo *merchantRepoImpl) InsertMerchant(merc *model.MerchantRequestModel) error {
	tx, err := merchantRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("error on merchantRepoImpl.InsertMerchant() 1 : %w", err)
	}

	qryUser := utils.INSERT_MERC
	err = tx.QueryRow(qryUser, utils.UuidGenerate(), merc.Username, merc.Password, merc.Role, merc.Active).Scan(&merc.User_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on merchantRepoImpl.InsertMerchant() 2  : %w", err)
	}

	qry := utils.INSERT_MERC_USR
	_, err = tx.Exec(qry, utils.UuidGenerate(), merc.User_id, merc.FullName, merc.NoPhone, merc.Email, merc.Address)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on merchantRepoImpl.InsertMerchant() 3 : %w", err)
	}

	tx.Commit()
	return nil
}

func (merchantRepo *merchantRepoImpl) GetMerchantById(id string) (*model.MerchantModel, error) {
	qry := utils.GET_MERC_ID
	merc := &model.MerchantModel{}
	err := merchantRepo.db.QueryRow(qry, id).Scan(&merc.ID, &merc.User_id, &merc.FullName, &merc.NoPhone, &merc.Email, &merc.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on merchantRepoImpl.GetMerchantById() : %w", err)
	}
	return merc, nil
}

func (merchantRepo *merchantRepoImpl) GetMerchantByUserId(id string) (*model.MerchantModel, error) {
	qry := utils.GET_MERC_USRID
	merc := &model.MerchantModel{}
	err := merchantRepo.db.QueryRow(qry, id).Scan(&merc.ID, &merc.User_id, &merc.FullName, &merc.NoPhone, &merc.Email, &merc.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on merchantRepoImpl.GetMerchantById() : %w", err)
	}
	return merc, nil
}

func (merchantRepo *merchantRepoImpl) GetAllMerchant() ([]*model.MerchantModel, error) {
	qry := utils.GET_ALL_MERCHANT
	rows, err := merchantRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on merchantRepoImpl.GetAllMerchant() : %w", err)
	}
	defer rows.Close()
	var arrMerchant []*model.MerchantModel
	for rows.Next() {
		merc := &model.MerchantModel{}
		rows.Scan(&merc.ID, &merc.User_id, &merc.FullName, &merc.NoPhone, &merc.Email, &merc.Address)
		arrMerchant = append(arrMerchant, merc)
	}
	return arrMerchant, nil
}

func (merchantRepo *merchantRepoImpl) GetMerchantByName(name string) (*model.MerchantModel, error) {
	qry := utils.GET_MERC_NAME
	merc := &model.MerchantModel{}
	err := merchantRepo.db.QueryRow(qry, name).Scan(&merc.ID, &merc.User_id, &merc.FullName, &merc.NoPhone, &merc.Email, &merc.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on merchantRepoImpl.GetMerchantByName() : %w", err)
	}
	return merc, nil
}

func (merchantRepo *merchantRepoImpl) EditMerchantById(merc model.MerchantModel) error {
	qry := utils.EDIT_MERC_ID
	_, err := merchantRepo.db.Exec(qry, merc.FullName, merc.NoPhone, merc.Email, merc.Address, merc.ID)
	if err != nil {
		return fmt.Errorf("error on merchantRepoImpl.EditMerchantById() 3 : %w", err)
	}

	return nil
}

func NewMerchantRepo(db *sql.DB) MerchantRepo {
	return &merchantRepoImpl{
		db: db,
	}
}
