package student

import (
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/student/delivery"
	"github.com/simanasiry/school/student/domain"
	"github.com/simanasiry/school/student/repository"
	"github.com/simanasiry/school/student/usecase"
	"github.com/simanasiry/school/utils"
	"gorm.io/gorm"
)

type Module struct {
	Repo        domain.Reposiory
	useCase     domain.Usecase
	Addstudent  utils.Handler
	GetStudents utils.Handler
}

func New(db *gorm.DB, schoolModule *school.Module) *Module {

	m := new(Module)
	m.Repo = repository.NewRepository(db)
	m.useCase = usecase.NewUc(m.Repo, schoolModule)
	m.Addstudent = delivery.NewStudentHandler(m.useCase)
	m.GetStudents = delivery.NewGetStudentsHandler(m.useCase)
	return m
}
