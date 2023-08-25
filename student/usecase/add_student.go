package usecase

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	sd "github.com/simanasiry/school/school/domain"
	"github.com/simanasiry/school/student/domain"
	"github.com/simanasiry/school/utils"
	"net/http"
)

func (uc *usecase) AddStudent(requestBody *domain.AddStudentRequest) (error, int, *domain.Student) {

	// userRequested must be master (check school table) to add students
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

	role := domain.Student{
		Student:  requestBody.Student,
		SchoolId: newSchool.ID,
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
