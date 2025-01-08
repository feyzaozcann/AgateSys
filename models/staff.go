package models

import "time"

type Staff struct {
	StaffID   int       `db:"staff_id" json:"staff_id"`
	Name      string    `db:"name" json:"name"`
	Role      string    `db:"role" json:"role"`
	GradeID   int       `db:"grade_id" json:"grade_id"`
	StartDate time.Time `db:"start_date" json:"start_date"`
}
