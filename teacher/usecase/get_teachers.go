package usecase

import (
	"github.com/simanasiry/school/teacher/domain"
	"net/http"
)

func (uc *usecase) GetTeachers() (error, int, *[]domain.Teacher) {

	var result []domain.Teacher
	err := uc.Repo.Model(&domain.Teacher{}).Scan(&result).Error
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}
	return nil, http.StatusOK, &result
}
