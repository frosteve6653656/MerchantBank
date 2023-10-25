package manager

import (
	"final-project/repository"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repository.UserRepo
	GetLoginRepo() repository.LoginRepo
	GetCustomerRepo() repository.CustomerRepo
	GetMerchantRepo() repository.MerchantRepo
	GetBankRepo() repository.BankRepo
	GetTransferRepo() repository.TransferRepo
}

type repoManager struct {
	infraManager InfraManager
	usrRepo      repository.UserRepo
	lgRepo       repository.LoginRepo
	custRepo     repository.CustomerRepo
	mercRepo     repository.MerchantRepo
	bankRepo     repository.BankRepo
	trxRepo      repository.TransferRepo
}

var onceLoadUserRepo sync.Once
var onceLoadLoginRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadMerchantRepo sync.Once
var onceLoadBankRepo sync.Once
var onceLoadTransferRepo sync.Once

func (rm *repoManager) GetUserRepo() repository.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func (rm *repoManager) GetLoginRepo() repository.LoginRepo {
	onceLoadLoginRepo.Do(func() {
		rm.lgRepo = repository.NewLoginRepo(rm.infraManager.GetDB())
	})
	return rm.lgRepo
}

func (rm *repoManager) GetCustomerRepo() repository.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.custRepo = repository.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.custRepo
}

func (rm *repoManager) GetMerchantRepo() repository.MerchantRepo {
	onceLoadMerchantRepo.Do(func() {
		rm.mercRepo = repository.NewMerchantRepo(rm.infraManager.GetDB())
	})
	return rm.mercRepo
}

func (rm *repoManager) GetBankRepo() repository.BankRepo {
	onceLoadBankRepo.Do(func() {
		rm.bankRepo = repository.NewBankRepo(rm.infraManager.GetDB())
	})
	return rm.bankRepo
}

func (rm *repoManager) GetTransferRepo() repository.TransferRepo {
	onceLoadTransferRepo.Do(func() {
		rm.trxRepo = repository.NewTransferRepo(rm.infraManager.GetDB())
	})
	return rm.trxRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
