package usecase

import (
	"final-project/model"
	"final-project/repository"
	"final-project/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	InsertUser(usr *model.UserModel) error
	GetUserById(id string) (*model.UserModel, error)
	GetUserByName(usr *model.LoginModel) (*model.UserModel, error)
	GetAllUser() ([]*model.UserModel, error)
	EditUserById(usr *model.UserModel) error
}

type userUseCaseImpl struct {
	userRepo repository.UserRepo
}

func (userRepoCase *userUseCaseImpl) GetUserById(id string) (*model.UserModel, error) {
	return userRepoCase.userRepo.GetUserById(id)
}

func (userRepoCase *userUseCaseImpl) GetAllUser() ([]*model.UserModel, error) {
	return userRepoCase.userRepo.GetAllUser()
}

func (userRepoCase *userUseCaseImpl) GetUserByName(usr *model.LoginModel) (*model.UserModel, error) {
	existData, err := userRepoCase.userRepo.GetUserByName(usr.Username)
	if err != nil {
		return nil, fmt.Errorf("userUsecaseImpl.GetUserByName(): %w", err)
	}

	existData.Password = ""
	return existData, nil
}

func (userRepoCase *userUseCaseImpl) InsertUser(usr *model.UserModel) error {
	if usr.UserName == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if usr.Password == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}

	existData, err := userRepoCase.userRepo.GetUserByName(usr.UserName)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.InsertUser(): %w", err)
	}
	if existData != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the name %v already exists", usr.UserName),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.GenerateFromPassword(): %w", err)
	}
	usr.Password = string(passHash)
	usr.Role = "Admin"
	usr.Active = true
	return userRepoCase.userRepo.InsertUser(usr)
}

func (userRepoCase *userUseCaseImpl) EditUserById(usr *model.UserModel) error {
	if usr.UserName == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if usr.Password == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}

	existData, err := userRepoCase.userRepo.GetUserById(usr.Id)
	if err != nil {
		return fmt.Errorf("userUseCaseImpl.EditUserById(): %w", err)
	}
	if existData == nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the ID %v does not exist", usr.Id),
		}
	}

	existDataUsr, err := userRepoCase.userRepo.GetUserByName(usr.UserName)
	if err != nil {
		return fmt.Errorf("userUseCaseImpl.GetUserByName(): %w", err)
	}
	if existDataUsr != nil && existDataUsr.Id != usr.Id {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the username %v already exists", usr.UserName),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.GenerateFromPassword(): %w", err)
	}
	usr.Password = string(passHash)
	usr.Role = "Customer"
	usr.Active = true
	return userRepoCase.userRepo.EditUserById(*usr)
}

func NewUserUseCase(userRepo repository.UserRepo) UserUseCase {
	return &userUseCaseImpl{
		userRepo: userRepo,
	}
}
