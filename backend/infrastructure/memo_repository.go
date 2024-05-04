package infrastructure

import (
	"gorm.io/gorm"

	"backend/domain"
	"backend/domain/repositories"
)

type MemoRepository struct {
	db *gorm.DB
}

func NewMemoRepository(db *gorm.DB) repositories.MemoRepository {
	return &MemoRepository{db}
}

func (r *MemoRepository) Create(memo *domain.Memo) error {
	return r.db.Create(memo).Error
}

func (r *MemoRepository) Read(hash string) (*domain.Memo, error) {
	var memo domain.Memo
	if err := r.db.Where("short_hash = ?", hash).Preload("Children").Preload("Parent").First(&memo).Error; err != nil {
		return nil, err
	}
	return &memo, nil
}

func (r *MemoRepository) ReadAll(page int, per int) (domain.Memos, error) {
	var memos domain.Memos
	if err := r.db.Offset(page * per).Limit(per).Preload("Children").Preload("Parent").Find(&memos).Error; err != nil {
		return nil, err
	}
	return memos, nil
}

func (r *MemoRepository) Update(memo *domain.Memo) error {
	return r.db.Save(memo).Error
}

func (r *MemoRepository) Delete(memo *domain.Memo) error {
	return r.db.Delete(memo).Error
}
