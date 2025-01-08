package services

import (
	"agate-project/models"
	"agate-project/repositories"
	"fmt"
	"log"
)

type StaffGradeService interface {
	FetchAllGrades() ([]models.StaffGrade, error)
	AddGrade(staffGrade *models.StaffGrade) error
	RemoveGrade(gradeID int) error
	UpdateGrade(gradeID int, gradeName *string, payRate *int) error
}

type staffGradeService struct {
	repo repositories.StaffGradeRepository
}

func NewStaffGradeService(repo repositories.StaffGradeRepository) StaffGradeService {
	return &staffGradeService{repo: repo}
}

func (s *staffGradeService) FetchAllGrades() ([]models.StaffGrade, error) {
	grades, err := s.repo.GetAllStaffGrades()
	if err != nil {
		log.Printf("FetchAllGrades: Error fetching grades: %v", err)
		return nil, fmt.Errorf("fetching grades failed: %w", err)
	}
	return grades, nil
}

func (s *staffGradeService) AddGrade(staffGrade *models.StaffGrade) error {
	if staffGrade.GradeName == "" || staffGrade.PayRate <= 0 {
		log.Println("AddGrade: Invalid grade data: name or pay rate is missing.")
		return fmt.Errorf("invalid grade data: name or pay rate is missing")
	}

	if err := s.repo.AddStaffGrade(staffGrade); err != nil {
		log.Printf("AddGrade: Error adding grade: %v", err)
		return fmt.Errorf("adding grade failed: %w", err)
	}

	return nil
}

func (s *staffGradeService) RemoveGrade(gradeID int) error {
	if gradeID <= 0 {
		log.Printf("RemoveGrade: Invalid grade ID: %d", gradeID)
		return fmt.Errorf("invalid grade ID")
	}

	if err := s.repo.DeleteStaffGrade(gradeID); err != nil {
		log.Printf("RemoveGrade: Error removing grade with ID %d: %v", gradeID, err)
		return fmt.Errorf("failed to remove grade with ID %d: %w", gradeID, err)
	}

	return nil
}

// *int çünkü güncellemek istemezsek nil olarak gönderilir
func (s *staffGradeService) UpdateGrade(gradeID int, gradeName *string, payRate *int) error {
	if gradeID <= 0 {
		log.Printf("UpdateGrade: Invalid grade ID: %d", gradeID)
		return fmt.Errorf("invalid grade ID")
	}

	if err := s.repo.UpdateStaffGrade(gradeID, gradeName, payRate); err != nil {
		log.Printf("UpdateGrade: Error updating grade with ID %d: %v", gradeID, err)
		return fmt.Errorf("failed to update grade with ID %d: %w", gradeID, err)
	}

	return nil
}
