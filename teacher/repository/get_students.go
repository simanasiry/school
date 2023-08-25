package repository

import (
	rg "github.com/simanasiry/school/registration/domain"
	std "github.com/simanasiry/school/student/domain"
)

func (repo *repository) GetStudents(teacherId uint64) (*[]std.Student, error) {
	//type StudentStr struct {
	//	Student string
	//}
	rows, err := repo.Model(&rg.Register{}).Select("students.*").Where("teachers.id=?", teacherId).
		Joins("join teachers on teachers.id = registrations.teacher_id").
		Joins("join students on students.id = registrations.student_id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []std.Student
	for rows.Next() {
		newStudent := std.Student{}
		err = repo.ScanRows(rows, &newStudent)
		if err != nil {
			return nil, err
		}
		result = append(result, newStudent)
	}
	return &result, nil

}
