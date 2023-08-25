package usecase

import (
	"github.com/simanasiry/school/registration/domain"
	"net/http"
)

func (uc *usecase) GetRegisters() (error, int, *[]domain.Register) {

	result := make([]domain.Register, 0)
	err := uc.repo.Model(&domain.Register{}).Scan(&result).Error
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}
	return nil, http.StatusOK, &result
}
