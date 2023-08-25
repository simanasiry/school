package usecase

import (
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/teacher/domain"
)

type usecase struct {
	Repo         domain.Reposiory
	schoolModule *school.Module
}

func NewUc(repo domain.Reposiory, schoolModule *school.Module) domain.Usecase {

	uc := &usecase{
		Repo:         repo,
		schoolModule: schoolModule,
	}
	return uc
}
