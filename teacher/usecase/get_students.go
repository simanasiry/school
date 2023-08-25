package usecase

import (
	std "github.com/simanasiry/school/student/domain"
	"net/http"
)

func (uc *usecase) GetStudents(teacherId uint64) (error, int, *[]std.Student) {

	uc.Repo.GetStudents(teacherId)
	return nil, http.StatusOK, nil
}
