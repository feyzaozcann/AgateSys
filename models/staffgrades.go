package models

type StaffGrade struct {
	GradeID   int    `db:"grade_id" json:"grade_id"`
	GradeName string `db:"grade_name" json:"grade_name"`
	PayRate   int    `db:"pay_rate" json:"pay_rate"`
}
