package usecase

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	sd "github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/teacher/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

func (uc *usecase) AddTeacher(requestBody *domain.AddTeacherRequest) (error, int, *domain.Teacher) {

	if requestBody.Role != utils.HeadMaster {
		return errors.New("only headmaster can add teachers"), http.StatusUnauthorized, nil
	}
	// userRequested must be master (check school table) to add teachers
	newSchool := sd.School{}
	err := uc.schoolModule.Repo.Model(&sd.School{}).
		Where("head_master = ? ", requestBody.UserName).Scan(&newSchool).Error
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}
	if newSchool.ID == 0 {
		return errors.New("no school has been found with this  headmaster"),
			http.StatusNotFound, nil
	}

	role := domain.Teacher{
		Teacher:  requestBody.Teacher,
		SchoolId: newSchool.ID,
	}
	err = uc.Repo.Create(&role).Error
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
