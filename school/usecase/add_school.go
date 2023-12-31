package usecase

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

func (uc *usecase) AddSchool(requestBody *domain.AddSchoolRequest) (error, int, *domain.School) {

	if requestBody.UserName != utils.DesiredAdmin || requestBody.Role != utils.Admin {
		return errors.New("only admin is allowed to add roles"), http.StatusUnauthorized, nil
	}

	role := domain.School{
		HeadMaster: requestBody.HeadMaster,
		School:     requestBody.School,
	}
	err := uc.repo.Create(&role).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				return err, http.StatusConflict, nil
			}
			return err, http.StatusInternalServerError, nil
		}

	}
	return nil, http.StatusOK, &role
}
