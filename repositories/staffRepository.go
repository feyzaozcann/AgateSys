package repositories

import (
	"agate-project/models"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type StaffRepository interface {
	GetAllStaff() ([]models.Staff, error)
	AddStaff(staff *models.Staff) error
	RemoveStaff(StaffID int) error
	UpdateStaff(StaffID int, updatedDetails *models.Staff) error
	GetStaffByID(staffID int) (models.Staff, error)
}

type staffRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewStaffRepository(ctx context.Context, db *sqlx.DB) StaffRepository {
	return &staffRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *staffRepository) GetAllStaff() ([]models.Staff, error) {
	var staff []models.Staff
	query := "SELECT staff_id, name, role, grade_id,starting_grade, start_date FROM staff"

	if err := r.db.SelectContext(r.ctx, &staff, query); err != nil {
		log.Printf("GetAllStaff: Failed to retrieve staff: %v", err)
		return nil, fmt.Errorf("failed to retrieve staff: %w", err)
	}
	return staff, nil
}

func (r *staffRepository) GetStaffByID(staffID int) (models.Staff, error) {
	var staff models.Staff
	query := `SELECT * FROM staff WHERE staff_id = $1`
	err := r.db.GetContext(r.ctx, &staff, query, staffID)
	if err != nil {
		log.Printf("GetStaffByID: Failed to get staff with ID %d: %v", staffID, err)
		return staff, fmt.Errorf("failed to get staff by ID: %w", err)
	}
	return staff, nil
}

func (r *staffRepository) AddStaff(staff *models.Staff) error {
	query := `
		INSERT INTO staff (name, role, grade_id, starting_grade) 
		VALUES (:name, :role, :grade_id, :starting_grade)
	`

	_, err := r.db.NamedExecContext(r.ctx, query, staff)
	if err != nil {
		log.Printf("AddStaff: Failed to add staff: %v", err)
		return fmt.Errorf("failed to add staff: %w", err)
	}
	return nil
}

func (r *staffRepository) RemoveStaff(staffID int) error {
	query := "DELETE FROM staff WHERE staff_id = $1"
	if _, err := r.db.ExecContext(r.ctx, query, staffID); err != nil {
		log.Printf("RemoveStaff: Failed to delete staff with ID %d: %v", staffID, err)
		return fmt.Errorf("failed to delete staff with ID %d: %w", staffID, err)
	}
	return nil
}

func (r *staffRepository) UpdateStaff(staffID int, updatedDetails *models.Staff) error {
	query := "UPDATE staff SET name = $1, role = $2, grade_id = $3, starting_grade =$4 WHERE staff_id = $5"

	_, err := r.db.ExecContext(r.ctx, query, updatedDetails.Name, updatedDetails.Role, updatedDetails.GradeID, staffID)
	if err != nil {
		log.Printf("UpdateStaff: Failed to update staff with ID %d: %v", staffID, err)
		return fmt.Errorf("failed to update staff with ID %d: %w", staffID, err)
	}
	return nil
}
