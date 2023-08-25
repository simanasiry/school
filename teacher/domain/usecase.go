package domain

import "github.com/simanasiry/school/student/domain"

type Usecase interface {
	AddTeacher(request *AddTeacherRequest) (error, int, *Teacher)
	GetTeachers() (error, int, *[]Teacher)
	GetStudents(teacherId uint64) (error, int, *[]domain.Student)
}
