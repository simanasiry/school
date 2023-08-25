package app

import (
	"github.com/simanasiry/school/registration"
	"github.com/simanasiry/school/school"
	"github.com/simanasiry/school/student"
	"github.com/simanasiry/school/teacher"
)

func (a *App) initSchools() {
	a.schoolModule = school.New(a.db)
}

func (a *App) initTeachers() {
	a.teacherModule = teacher.New(a.db, a.schoolModule)
}

func (a *App) initStudents() {
	a.studentModule = student.New(a.db, a.schoolModule)
}

func (a *App) initRegistrations() {
	a.registrationModule = registration.New(a.db, a.schoolModule, a.studentModule, a.teacherModule)
}
