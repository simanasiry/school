package usecase

import (
	"errors"
	"github.com/simanasiry/school/registration/domain"
	sd "github.com/simanasiry/school/school/domain"
	std "github.com/simanasiry/school/student/domain"
	td "github.com/simanasiry/school/teacher/domain"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func (uc *usecase) AddRegister(requestBody *domain.AddRegisterRequest) (error, int, *domain.Register) {

	// userRequested must be master (check school table) to add Registers
	school := sd.School{}
	err := uc.schoolModule.Repo.Model(&sd.School{}).
		Where("head_master =?", requestBody.UserRequested).Scan(&school).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no school has been found with this headmaster"),
				http.StatusNotFound, nil
		}
		return err, http.StatusInternalServerError, nil
	}

	//  check student belongs to this school
	err = uc.studentModule.Repo.First(&std.Student{}, "id = ? AND school_id = ?",
		requestBody.StudentId, school.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("this student doesn't belong to this school"),
				http.StatusNotFound, nil
		}
		return err, http.StatusInternalServerError, nil
	}

	//  check teacher belongs to this school
	err = uc.teacherModule.Repo.First(&td.Teacher{}, "id = ? AND school_id = ?",
		requestBody.TeacherId, school.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("this teacher doesn't belong to this school"),
				http.StatusNotFound, nil
		}
		return err, http.StatusInternalServerError, nil
	}

	role := domain.Register{
		TeacherId: requestBody.TeacherId,
		StudentId: requestBody.StudentId,
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
