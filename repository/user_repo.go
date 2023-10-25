package repository

import (
	"database/sql"
	"final-project/model"
	"final-project/utils"
	"fmt"
)

type UserRepo interface {
	InsertUser(usr *model.UserModel) error
	GetUserById(id string) (*model.UserModel, error)
	GetUserByName(name string) (*model.UserModel, error)
	GetAllUser() ([]*model.UserModel, error)
	EditUserById(usr model.UserModel) error
}

type userRepoImpl struct {
	db *sql.DB
}

func (userRepo *userRepoImpl) InsertUser(usr *model.UserModel) error {
	qry := utils.INSERT_USER
	fmt.Println(utils.UuidGenerate())
	_, err := userRepo.db.Exec(qry, utils.UuidGenerate(), usr.UserName, usr.Password, usr.Role, usr.Active)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.InsertUser() : %w", err)
	}
	return nil
}

func (userRepo *userRepoImpl) GetUserById(id string) (*model.UserModel, error) {
	qry := utils.GET_USER_BY_ID
	usr := &model.UserModel{}
	err := userRepo.db.QueryRow(qry, id).Scan(&usr.Id, &usr.UserName, &usr.Role, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.GetUserById() : %w", err)
	}
	return usr, nil
}

func (userRepo *userRepoImpl) GetAllUser() ([]*model.UserModel, error) {
	qry := utils.GET_ALL_USER
	rows, err := userRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on userRepoImpl.GetAllUser() : %w", err)
	}
	defer rows.Close()
	var arrUser []*model.UserModel
	for rows.Next() {
		usr := &model.UserModel{}
		rows.Scan(&usr.Id, &usr.UserName, &usr.Role, &usr.Active)
		arrUser = append(arrUser, usr)
	}
	return arrUser, nil
}

func (userRepo *userRepoImpl) GetUserByName(name string) (*model.UserModel, error) {
	qry := utils.GET_USER_BY_NAME
	usr := &model.UserModel{}
	err := userRepo.db.QueryRow(qry, name).Scan(&usr.Id, &usr.UserName, &usr.Password, &usr.Role, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.GetUserByName() : %w", err)
	}
	return usr, nil
}

func (userRepo *userRepoImpl) EditUserById(usr model.UserModel) error {
	qry := utils.EDIT_USER_ID
	_, err := userRepo.db.Exec(qry, usr.UserName, usr.Password, usr.Active, usr.Id)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.EditUserById() : %w", err)
	}
	return nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}
