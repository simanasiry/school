package usecase

import (
	"errors"
	sd "github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/teacher/domain"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func (uc *usecase) AddTeacher(requestBody *domain.AddTeacherRequest) (error, int, *domain.Teacher) {

	// userRequested must be master (check school table) to add teachers
	school := sd.School{
		HeadMaster: requestBody.UserRequested,
		ID:         requestBody.SchoolId,
	}
	err := uc.schoolModule.Repo.First(&school).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no school has been found with this id and headmaster"),
				http.StatusNotFound, nil
		}
		return err, http.StatusInternalServerError, nil
	}

	role := domain.Teacher{
		Teacher:  requestBody.Teacher,
		SchoolId: requestBody.SchoolId,
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
