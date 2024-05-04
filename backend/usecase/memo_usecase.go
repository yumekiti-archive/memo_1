package usecase

import (
	"backend/domain"
	"backend/domain/repositories"
)

type MemoUsecase struct {
	memoRepository repositories.MemoRepository
}

func NewMemoUsecase(memoRepository repositories.MemoRepository) MemoUsecase {
	return MemoUsecase{memoRepository}
}

func (u *MemoUsecase) Create(memo *domain.Memo) error {
	memo.GenerateShortHash()
	if err := memo.Validate(); err != nil {
		return err
	}
	return u.memoRepository.Create(memo)
}

func (u *MemoUsecase) Read(hash string) (*domain.Memo, error) {
	return u.memoRepository.Read(hash)
}

func (u *MemoUsecase) ReadAll(page int, per int) (domain.Memos, error) {
	return u.memoRepository.ReadAll(page, per)
}

func (u *MemoUsecase) Update(memo *domain.Memo) error {
	return u.memoRepository.Update(memo)
}

func (u *MemoUsecase) Delete(memo *domain.Memo) error {
	return u.memoRepository.Delete(memo)
}
