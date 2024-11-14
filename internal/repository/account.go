package repository

import (
	"PandoraHelper/internal/model"
	"context"
	"github.com/gin-gonic/gin"
)

type AccountRepository interface {
	GetAccount(ctx context.Context, id int64) (*model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Create(ctx context.Context, account *model.Account) error
	SearchAccount(ctx context.Context, accountType string, keyword string) ([]*model.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetShareAccountList(ctx *gin.Context) ([]*model.Account, error)
}

func NewAccountRepository(
	repository *Repository,

) AccountRepository {
	return &accountRepository{
		Repository: repository,
	}
}

type accountRepository struct {
	*Repository
}

func (r *accountRepository) GetShareAccountList(ctx *gin.Context) ([]*model.Account, error) {
	var accounts []*model.Account
	if err := r.DB(ctx).
		Select("id,account_type, shared").
		Where("shared = ?", 1).
		Find(&accounts).
		Error; err != nil {
		return nil, err
	}
	r.logger.Info("accounts")
	return accounts, nil
}

func (r *accountRepository) SearchAccount(ctx context.Context, accountType string, keyword string) ([]*model.Account, error) {
	var accounts []*model.Account
	if err := r.DB(ctx).
		Where("account_type = ?", accountType).
		Where("email like ?", "%"+keyword+"%").
		Find(&accounts).
		Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *accountRepository) Update(ctx context.Context, account *model.Account) error {
	if err := r.DB(ctx).Save(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) Create(ctx context.Context, account *model.Account) error {
	if err := r.DB(ctx).Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) DeleteAccount(ctx context.Context, id int64) error {
	r.DB(ctx).Delete(&model.Account{}, id)
	return nil
}

func (r *accountRepository) GetAccount(ctx context.Context, id int64) (*model.Account, error) {
	var account model.Account
	if err := r.DB(ctx).Preload("Shares").Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
