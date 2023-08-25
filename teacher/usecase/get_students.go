package usecase

import (
	std "github.com/simanasiry/school/student/domain"
	"net/http"
)

func (uc *usecase) GetStudents(teacherId uint64) (error, int, *[]std.Student) {

	result, err := uc.Repo.GetStudents(teacherId)
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}
	return nil, http.StatusOK, result
}
