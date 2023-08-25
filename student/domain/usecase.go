package domain

type Usecase interface {
	AddStudent(request *AddStudentRequest) (error, int, *Student)
}
