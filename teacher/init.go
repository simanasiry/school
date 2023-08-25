package teacher

import (
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/teacher/delivery"
	"github.com/simanasiry/school/teacher/domain"
	"github.com/simanasiry/school/teacher/repository"
	"github.com/simanasiry/school/teacher/usecase"
	"github.com/simanasiry/school/utils"
	"gorm.io/gorm"
)

type Module struct {
	Repo        domain.Reposiory
	useCase     domain.Usecase
	AddTeacher  utils.Handler
	GetStudents utils.Handler
	GetTeachers utils.Handler
}

func New(db *gorm.DB, schoolModule *school.Module) *Module {

	m := new(Module)
	m.Repo = repository.NewRepository(db)
	m.useCase = usecase.NewUc(m.Repo, schoolModule)
	m.AddTeacher = delivery.NewTeacherHandler(m.useCase)
	m.GetStudents = delivery.NewStudentHandler(m.useCase)
	m.GetTeachers = delivery.NewGetTeachersHandler(m.useCase)
	return m
}
