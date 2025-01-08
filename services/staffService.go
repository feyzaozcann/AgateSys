package services

import (
	"agate-project/models"
	"agate-project/repositories"
	"fmt"
	"log"
)

type StaffService interface {
	FetchAllStaff() ([]models.Staff, error)
	AddStaff(staff *models.Staff) error
	RemoveStaff(staffID int) error
	UpdateStaff(staffID int, updatedDetails *models.Staff) error
	GetStaffByID(staffID int) (models.Staff, error)
}

type staffService struct {
	repo repositories.StaffRepository
}

func NewStaffService(repo repositories.StaffRepository) StaffService {
	return &staffService{repo: repo}
}

func (s *staffService) FetchAllStaff() ([]models.Staff, error) {
	staff, err := s.repo.GetAllStaff()
	if err != nil {
		log.Printf("FetchAllStaff: Error fetching staff: %v", err)
		return nil, fmt.Errorf("fetching staff failed: %w", err)
	}
	return staff, nil
}

func (s *staffService) GetStaffByID(staffID int) (models.Staff, error) {
	staff, err := s.repo.GetStaffByID(staffID)
	if err != nil {
		log.Printf("GetStaffByID: Failed to retrieve staff with ID %d: %v", staffID, err)
		return staff, fmt.Errorf("failed to retrieve staff by ID: %w", err)
	}
	return staff, nil
}

func (s *staffService) AddStaff(staff *models.Staff) error {
	if err := s.repo.AddStaff(staff); err != nil {
		log.Printf("AddStaff: Error adding staff: %v", err)
		return fmt.Errorf("adding staff failed: %w", err)
	}
	return nil
}

func (s *staffService) RemoveStaff(staffID int) error {
	if staffID <= 0 {
		log.Printf("RemoveStaff: Invalid staff ID: %d", staffID)
		return fmt.Errorf("invalid staff ID")
	}

	if err := s.repo.RemoveStaff(staffID); err != nil {
		log.Printf("RemoveStaff: Error removing staff with ID %d: %v", staffID, err)
		return fmt.Errorf("failed to remove staff with ID %d: %w", staffID, err)
	}
	return nil
}

func (s *staffService) UpdateStaff(staffID int, updatedDetails *models.Staff) error {
	if staffID <= 0 {
		log.Printf("UpdateStaff: Invalid staff ID: %d", staffID)
		return fmt.Errorf("invalid staff ID")
	}

	if err := s.repo.UpdateStaff(staffID, updatedDetails); err != nil {
		log.Printf("UpdateStaff: Error updating staff with ID %d: %v", staffID, err)
		return fmt.Errorf("failed to update staff with ID %d: %w", staffID, err)
	}
	return nil
}
