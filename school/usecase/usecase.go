package usecase

import (
	"github.com/simanasiry/school/school/domain"
)

type usecase struct {
	repo domain.Reposiory
}

func NewUc(repo domain.Reposiory) domain.Usecase {

	uc := &usecase{
		repo: repo,
	}
	return uc
}
