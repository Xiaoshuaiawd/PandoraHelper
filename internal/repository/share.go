package repository

import (
	"PandoraHelper/internal/model"
	"context"
	"fmt"
)

type ShareRepository interface {
	GetShare(ctx context.Context, id int64) (*model.Share, error)
	Update(ctx context.Context, share *model.Share) error
	Create(ctx context.Context, share *model.Share) error
	SearchShare(ctx context.Context, accountType string, email string, uniqueName string) ([]*model.Share, error)
	DeleteShare(ctx context.Context, id int64) error
	GetShareByUniqueName(ctx context.Context, uniqueName string) (*model.Share, error)
	GetSharesByAccountId(ctx context.Context, accountId int) ([]*model.Share, error)
}

func NewShareRepository(
	repository *Repository,
) ShareRepository {
	return &shareRepository{
		Repository: repository,
	}
}

type shareRepository struct {
	*Repository
}

func (r *shareRepository) GetSharesByAccountId(ctx context.Context, accountId int) ([]*model.Share, error) {
	var shares []*model.Share
	if err := r.DB(ctx).Where("account_id = ?", accountId).Find(&shares).Error; err != nil {
		return nil, err
	}
	return shares, nil
}

func (r *shareRepository) GetShareByUniqueName(ctx context.Context, uniqueName string) (*model.Share, error) {
	var shares []*model.Share
	if err := r.DB(ctx).Where("unique_name = ?", uniqueName).Find(&shares).Error; err != nil {
		return nil, err
	}
	if len(shares) == 0 {
		return nil, nil
	}
	if len(shares) > 1 {
		return nil, fmt.Errorf("unique_name %s has more than one share", uniqueName)
	}
	return shares[0], nil
}

func (r *shareRepository) Update(ctx context.Context, share *model.Share) error {
	if err := r.DB(ctx).Save(share).Error; err != nil {
		return err
	}
	return nil
}

func (r *shareRepository) Create(ctx context.Context, share *model.Share) error {
	if err := r.DB(ctx).Create(share).Error; err != nil {
		return err
	}
	return nil
}

func (r *shareRepository) SearchShare(ctx context.Context, accountType string, email string, uniqueName string) ([]*model.Share, error) {
	// 关联Account like 模糊查询
	var shares []*model.Share
	if err := r.DB(ctx).Joins(
		"join account on share.account_id = account.id",
	).Select(
		"share.*, account.email",
	).Where("account.email like ?", "%"+email+"%").
		Where("unique_name like ?", "%"+uniqueName+"%").
		Where("share_type = ?", accountType).
		Find(&shares).Error; err != nil {
		return nil, err
	}
	return shares, nil
}

func (r *shareRepository) DeleteShare(ctx context.Context, id int64) error {
	r.DB(ctx).Delete(&model.Share{}, id)
	return nil
}

func (r *shareRepository) GetShare(ctx context.Context, id int64) (*model.Share, error) {
	var share model.Share
	if err := r.DB(ctx).Where("id = ?", id).First(&share).Error; err != nil {
		return nil, err
	}
	return &share, nil
}
