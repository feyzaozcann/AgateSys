package repositories

import (
	"agate-project/models"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type StaffGradeRepository interface {
	AddStaffGrade(staffGrade *models.StaffGrade) error
	GetAllStaffGrades() ([]models.StaffGrade, error)
	GetStaffGradeById(gradeID int) (models.StaffGrade, error)
	DeleteStaffGrade(gradeID int) error
	UpdateStaffGrade(gradeID int, gradeName *string, payRate *int) error
	AssignStaffToGrade(staffID int, gradeID int) error
}

type staffGradeRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewStaffGradeRepository(ctx context.Context, db *sqlx.DB) StaffGradeRepository {
	return &staffGradeRepository{
		db:  db,
		ctx: ctx,
	}
}

func (s *staffGradeRepository) AddStaffGrade(staffGrade *models.StaffGrade) error {
	query := `INSERT INTO staff_grades (grade_name, pay_rate) 
              VALUES ($1, $2) RETURNING grade_id`
	if err := s.db.GetContext(s.ctx, &staffGrade.GradeID, query, staffGrade.GradeName, staffGrade.PayRate); err != nil {
		log.Printf("AddStaffGrade: Failed to add staff grade: %v", err)
		return fmt.Errorf("failed to add staff grade: %w", err)
	}
	return nil
}

func (s *staffGradeRepository) GetStaffGradeById(gradeID int) (models.StaffGrade, error) {
	query := `SELECT grade_id, grade_name, pay_rate
			  FROM staff_grades
			  WHERE grade_id = $1`
	var grade models.StaffGrade
	err := s.db.GetContext(s.ctx, &grade, query, gradeID)
	if err != nil {
		log.Printf("GetStaffGradeById: Failed to get staff grade with ID %d: %v", gradeID, err)
		return grade, fmt.Errorf("failed to get staff grade with id %d: %w", gradeID, err)
	}
	return grade, nil
}

func (s *staffGradeRepository) GetAllStaffGrades() ([]models.StaffGrade, error) {
	query := `SELECT grade_id, grade_name, pay_rate FROM staff_grades`
	var grades []models.StaffGrade
	err := s.db.SelectContext(s.ctx, &grades, query)
	if err != nil {
		log.Printf("GetAllStaffGrades: Failed to get all staff grades: %v", err)
		return nil, fmt.Errorf("failed to get all staff grades: %w", err)
	}
	return grades, nil
}

func (s *staffGradeRepository) DeleteStaffGrade(gradeID int) error {
	query := `DELETE FROM staff_grades WHERE grade_id = $1`
	_, err := s.db.ExecContext(s.ctx, query, gradeID)
	if err != nil {
		log.Printf("DeleteStaffGrade: Failed to delete staff grade with ID %d: %v", gradeID, err)
		return fmt.Errorf("failed to delete staff grade: %w", err)
	}
	return nil
}

func (s *staffGradeRepository) UpdateStaffGrade(gradeID int, gradeName *string, payRate *int) error {
	existingGrade, err := s.GetStaffGradeById(gradeID)
	if err != nil {
		log.Printf("UpdateStaffGrade: Failed to fetch existing staff grade with ID %d: %v", gradeID, err)
		return err
	}

	newGradeName := existingGrade.GradeName
	if gradeName != nil {
		newGradeName = *gradeName
	}

	newPayRate := existingGrade.PayRate
	if payRate != nil {
		newPayRate = *payRate
	}

	query := `
		UPDATE staff_grades 
		   SET grade_name = $1, pay_rate = $2
		 WHERE grade_id = $3
	`
	_, err = s.db.ExecContext(s.ctx, query, newGradeName, newPayRate, gradeID)
	if err != nil {
		log.Printf("UpdateStaffGrade: Failed to update staff grade with ID %d: %v", gradeID, err)
		return fmt.Errorf("failed to update staff grade: %w", err)
	}
	return nil
}

func (s *staffGradeRepository) AssignStaffToGrade(staffID int, gradeID int) error {
	query := `UPDATE staff SET grade_id = $1 WHERE staff_id = $2`
	_, err := s.db.ExecContext(s.ctx, query, gradeID, staffID)
	if err != nil {
		log.Printf("AssignStaffToGrade: Failed to assign staff with ID %d to grade with ID %d: %v", staffID, gradeID, err)
		return fmt.Errorf("failed to assign staff to grade: %w", err)
	}
	return nil
}
