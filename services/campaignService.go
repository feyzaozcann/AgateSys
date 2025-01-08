package services

import (
	"agate-project/models"
	"agate-project/repositories"
	"fmt"
	"log"
)

type CampaignService interface {
	CreateCampaign(campaign models.Campaign) error
	GetCampaignByID(campaignID int) (models.Campaign, error)
	UpdateCampaign(campaign models.Campaign) error
	RemoveCampaign(campaignID int) error
	AssignManager(campaignID, managerID int) error
	FetchAllCampaigns() ([]models.Campaign, error)
	//CheckBudget(campaignID int) (float64, error)
	GetCampaignsByClientID(clientID int) ([]models.Campaign, error)
}

type campaignService struct {
	repo repositories.CampaignRepository
}

func NewCampaignService(repo repositories.CampaignRepository) CampaignService {
	return &campaignService{repo: repo}
}

func (s *campaignService) CreateCampaign(campaign models.Campaign) error {
	log.Println("CreateCampaign: Attempting to create a new campaign.")
	if err := s.repo.CreateCampaign(campaign); err != nil {
		log.Printf("CreateCampaign: Error creating campaign: %v", err)
		return fmt.Errorf("failed to create campaign: %w", err)
	}
	return nil
}

func (s *campaignService) FetchAllCampaigns() ([]models.Campaign, error) {
	log.Println("FetchAllCampaigns: Fetching all campaigns.")
	campaigns, err := s.repo.GetAllCampaigns()
	if err != nil {
		log.Printf("FetchAllCampaigns: Error fetching campaigns: %v", err)
		return nil, fmt.Errorf("fetching campaigns failed: %w", err)
	}
	return campaigns, nil
}

func (s *campaignService) GetCampaignByID(campaignID int) (models.Campaign, error) {
	log.Printf("GetCampaignByID: Fetching campaign with ID %d.", campaignID)
	campaign, err := s.repo.GetCampaignByID(campaignID)
	if err != nil {
		log.Printf("GetCampaignByID: Error fetching campaign by ID: %v", err)
		return campaign, fmt.Errorf("failed to fetch campaign by ID: %w", err)
	}
	return campaign, nil
}

func (s *campaignService) UpdateCampaign(campaign models.Campaign) error {
	log.Printf("UpdateCampaign: Attempting to update campaign with ID %d.", campaign.CampaignID)
	if err := s.repo.UpdateCampaign(campaign); err != nil {
		log.Printf("UpdateCampaign: Error updating campaign with ID %d: %v", campaign.CampaignID, err)
		return fmt.Errorf("failed to update campaign: %w", err)
	}
	return nil
}

func (s *campaignService) RemoveCampaign(campaignID int) error {
	log.Printf("RemoveCampaign: Removing campaign with ID %d.", campaignID)
	if err := s.repo.DeleteCampaign(campaignID); err != nil {
		log.Printf("RemoveCampaign: Failed to remove campaign with ID %d: %v", campaignID, err)
		return fmt.Errorf("failed to remove campaign with id %d: %w", campaignID, err)
	}
	return nil
}

func (s *campaignService) AssignManager(campaignID, managerID int) error {
	log.Printf("AssignManager: Assigning manager with ID %d to campaign with ID %d.", managerID, campaignID)
	if err := s.repo.AssignManager(campaignID, managerID); err != nil {
		log.Printf("AssignManager: Error assigning manager with ID %d to campaign with ID %d: %v", managerID, campaignID, err)
		return fmt.Errorf("failed to assign manager to campaign: %w", err)
	}
	return nil
}

// CheckBudget checks the budget of a campaign
/*func (s *campaignService) CheckBudget(campaignID int) (float64, error) {
	budgetDifference, err := s.repo.CheckBudget(campaignID)
	if err != nil {
		log.Printf("Error checking budget for campaign: %v", err)
		return 0, fmt.Errorf("failed to check budget for campaign: %w", err)
	}
	return budgetDifference, nil
}*/

// GetCampaignsByClientID fetches all campaigns for a specific client
func (s *campaignService) GetCampaignsByClientID(clientID int) ([]models.Campaign, error) {
	log.Printf("GetCampaignsByClientID: Fetching campaigns for client ID %d.", clientID)
	campaigns, err := s.repo.GetCampaignsByClientID(clientID)
	if err != nil {
		log.Printf("GetCampaignsByClientID: Error fetching campaigns for client ID %d: %v", clientID, err)
		return nil, fmt.Errorf("failed to fetch campaigns for client ID %d: %w", clientID, err)
	}
	return campaigns, nil
}
