package domain

type Usecase interface {
	AddSchool(request *AddSchoolRequest) (error, int, *School)
	GetSchools() (error, int, *[]School)
}
