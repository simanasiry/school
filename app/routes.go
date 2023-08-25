package app

func (a *App) RegisterRoutes() {
	router := a.gin
	v1 := router.Group("/v1")
	v1.POST("/school/add", a.schoolModule.AddSchool.Handle())

	v1.POST("/teacher/add", a.teacherModule.AddTeacher.Handle())
	v1.GET("/teacher/:id/students", a.teacherModule.GetStudents.Handle())

	v1.POST("/student/add", a.studentModule.Addstudent.Handle())
	
	v1.POST("/registration/add", a.registrationModule.AddRegistration.Handle())

}
