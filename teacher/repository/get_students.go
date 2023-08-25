package repository

import (
	"fmt"
	rg "github.com/simanasiry/school/registration/domain"
	std "github.com/simanasiry/school/student/domain"
)

func (repo *repository) GetStudents(teacherId uint64) (*[]std.Student, error) {
	type StudentStr struct {
		Student string
	}
	rows, err := repo.Model(&rg.Register{}).Select("students.student").Where("teachers.id=?", teacherId).
		Joins("join teachers on teachers.id = registrations.teacher_id").
		Joins("join students on students.id = registrations.student_id").Rows()
	defer rows.Close()
	for rows.Next() {
		vv := StudentStr{}
		repo.ScanRows(rows, &vv)
		fmt.Println(vv)
	}
	fmt.Println(rows, err)
	return nil, nil

}
