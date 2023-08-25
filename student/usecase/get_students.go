package usecase

import (
	"github.com/simanasiry/school/student/domain"
	"net/http"
)

func (uc *usecase) GetStudents() (error, int, *[]domain.Student) {
	result := make([]domain.Student, 0)
	err := uc.repo.Model(&domain.Student{}).Scan(&result).Error
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}

	return nil, http.StatusOK, &result
}
