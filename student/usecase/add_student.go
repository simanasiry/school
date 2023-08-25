package usecase

import (
	"errors"
	sd "github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/student/domain"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func (uc *usecase) AddStudent(requestBody *domain.AddStudentRequest) (error, int, *domain.Student) {

	// userRequested must be master (check school table) to add students
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

	role := domain.Student{
		Student:  requestBody.Student,
		SchoolId: requestBody.SchoolId,
	}
	err = uc.repo.Create(&role).Error
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
