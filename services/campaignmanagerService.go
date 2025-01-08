package services

import (
	"agate-project/models"
	"agate-project/repositories"
	"fmt"
	"log"
)

type CampaignManagerService interface {
	GetAllCampaignManager() ([]models.CampaignManager, error)
	AddCampaignManager(campaignManager *models.CampaignManager) error
	DeleteCampaignManager(managerID int) error
	//AssignStaffToManager(staffID int, managerID int) error
}

type campaignManagerService struct {
	repo repositories.CampaignManagerRepository
}

func NewCampaignManagerService(repo repositories.CampaignManagerRepository) CampaignManagerService {
	return &campaignManagerService{
		repo: repo,
	}
}

func (c *campaignManagerService) GetAllCampaignManager() ([]models.CampaignManager, error) {
	campaignManager, err := c.repo.GetAllCampaignManager()
	if err != nil {
		log.Printf("GetAllCampaignManager: Error fetching campaign managers: %v", err)
		return nil, fmt.Errorf("error fetching campaign managers: %w", err)
	}
	return campaignManager, nil
}

func (c *campaignManagerService) AddCampaignManager(campaignManager *models.CampaignManager) error {
	if err := c.repo.AddCampaignManager(campaignManager); err != nil {
		log.Printf("AddCampaignManager: Error adding campaign manager: %v", err)
		return fmt.Errorf("error adding campaign manager: %w", err)
	}
	return nil
}

func (c *campaignManagerService) DeleteCampaignManager(managerID int) error {
	if err := c.repo.DeleteCampaignManager(managerID); err != nil {
		log.Printf("DeleteCampaignManager: Error deleting campaign manager with ID %d: %v", managerID, err)
		return fmt.Errorf("error deleting campaign manager with id %d: %w", managerID, err)
	}
	return nil
}

/*func (c *campaignManagerService) AssignStaffToManager(staffID int, managerID int) error {
	if err := c.repo.AssignStaffToManager(staffID, managerID); err != nil {
		log.Printf("AssignStaffToManager: Error assigning staff with ID %d to manager with ID %d: %v", staffID, managerID, err)
		return fmt.Errorf("error assigning staff with id %d to manager with id %d: %w", staffID, managerID, err)
	}
	return nil
}*/
