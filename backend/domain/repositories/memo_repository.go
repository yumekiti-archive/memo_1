package repositories

import "backend/domain"

type MemoRepository interface {
	Create(memo *domain.Memo) error
	Read(hash string) (*domain.Memo, error)
	ReadAll(page int, per int) (domain.Memos, error)
	Update(memo *domain.Memo) error
	Delete(memo *domain.Memo) error
}
