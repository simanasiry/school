package usecase

import (
	"github.com/simanasiry/school/school/domain"
	"net/http"
)

func (uc *usecase) GetSchools() (error, int, *[]domain.School) {

	var result []domain.School
	err := uc.repo.Model(&domain.School{}).Scan(&result).Error
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}
	return nil, http.StatusOK, &result
}
