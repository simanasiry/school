package school

import (
	"github.com/simanasiry/school/school/delivery"
	"github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/school/repository"
	"github.com/simanasiry/school/school/usecase"
	"github.com/simanasiry/school/utils"
	"gorm.io/gorm"
)

type Module struct {
	Repo      domain.Reposiory
	useCase   domain.Usecase
	AddSchool utils.Handler
}

func New(db *gorm.DB) *Module {

	m := new(Module)
	m.Repo = repository.NewRepository(db)
	m.useCase = usecase.NewUc(m.Repo)
	m.AddSchool = delivery.NewSchoolHandler(m.useCase)
	return m
}
