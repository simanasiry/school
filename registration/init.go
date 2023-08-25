package registration

import (
	"github.com/simanasiry/school/registration/delivery"
	"github.com/simanasiry/school/registration/domain"
	"github.com/simanasiry/school/registration/repository"
	"github.com/simanasiry/school/registration/usecase"
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/student"
	"github.com/simanasiry/school/teacher"
	"github.com/simanasiry/school/utils"
	"gorm.io/gorm"
)

type Module struct {
	repo            domain.Reposiory
	useCase         domain.Usecase
	AddRegistration utils.Handler
}

func New(db *gorm.DB, schoolModule *school.Module,
	studentModule *student.Module, teacherModule *teacher.Module) *Module {

	m := new(Module)
	m.repo = repository.NewRepository(db)
	m.useCase = usecase.NewUc(m.repo, schoolModule, studentModule, teacherModule)
	m.AddRegistration = delivery.NewRegisterHandler(m.useCase)
	return m
}
