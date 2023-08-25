package app

func (a *App) RegisterRoutes() {
	router := a.gin
	v1 := router.Group("/v1")
	v1.POST("/school/add", a.schoolModule.AddSchool.Handle())
	v1.GET("/school/items", a.schoolModule.GetSchools.Handle())

	v1.POST("/teacher/add", a.teacherModule.AddTeacher.Handle())
	v1.GET("/teacher/items", a.teacherModule.GetTeachers.Handle())
	v1.GET("/teacher/students/:teacherId", a.teacherModule.GetStudents.Handle())

	v1.POST("/student/add", a.studentModule.Addstudent.Handle())
	v1.GET("/student/items", a.studentModule.GetStudents.Handle())

	v1.POST("/registration/add", a.registrationModule.AddRegistration.Handle())
	v1.GET("/registration/items", a.registrationModule.GetRegisters.Handle())

}
