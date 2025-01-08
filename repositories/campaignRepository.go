package repositories

import (
	"agate-project/models"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type CampaignRepository interface {
	CreateCampaign(campaign models.Campaign) error
	GetCampaignByID(campaignID int) (models.Campaign, error)
	UpdateCampaign(campaign models.Campaign) error
	DeleteCampaign(campaignID int) error
	AssignManager(campaignID, managerID int) error
	GetAllCampaigns() ([]models.Campaign, error)
	//CheckBudget(campaignID int) (float64, error)
	GetCampaignsByClientID(clientID int) ([]models.Campaign, error)
}

type campaignRepository struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewCampaignRepository(ctx context.Context, db *sqlx.DB) CampaignRepository {
	return &campaignRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *campaignRepository) CreateCampaign(campaign models.Campaign) error {
	log.Println("CreateCampaign: Starting to create a new campaign.")
	query := `
    INSERT INTO campaigns (
        client_id, title, start_date, end_date, estimated_cost, actual_cost, completion_status, current_state, manager_id, budget
    ) VALUES (
        :client_id, :title, :start_date, :end_date, :estimated_cost, :actual_cost, :completion_status, :current_state, :manager_id, :budget
    );`

	_, err := r.db.NamedExecContext(r.ctx, query, campaign)
	if err != nil {
		log.Printf("CreateCampaign: Failed to create campaign: %v\n", err)
		return fmt.Errorf("failed to create campaign: %w", err)
	}
	return nil
}

func (r *campaignRepository) GetAllCampaigns() ([]models.Campaign, error) {
	log.Println("GetAllCampaigns: Fetching all campaigns.")
	query := `SELECT campaign_id, title, start_date, end_date, estimated_cost, actual_cost, completion_status, current_state, manager_id, budget FROM campaigns`
	var campaigns []models.Campaign
	err := r.db.SelectContext(r.ctx, &campaigns, query)
	if err != nil {
		log.Printf("GetAllCampaigns: Failed to fetch campaigns: %v\n", err)
		return nil, fmt.Errorf("failed to get all campaigns: %w", err)
	}
	return campaigns, nil
}

func (r *campaignRepository) GetCampaignByID(campaignID int) (models.Campaign, error) {
	log.Printf("GetCampaignByID: Fetching campaign with ID %d.\n", campaignID)
	var campaign models.Campaign
	query := "SELECT * FROM campaigns WHERE campaign_id = $1"
	err := r.db.GetContext(r.ctx, &campaign, query, campaignID)
	if err != nil {
		log.Printf("GetCampaignByID: Failed to fetch campaign with ID %d: %v\n", campaignID, err)
		return campaign, fmt.Errorf("failed to get campaign by ID: %w", err)
	}
	return campaign, nil
}

func (r *campaignRepository) UpdateCampaign(campaign models.Campaign) error {
	log.Printf("UpdateCampaign: Updating campaign with ID %d.\n", campaign.CampaignID)
	query := `
		UPDATE campaigns 
		SET client_id = $1, title = $2, start_date = $3, end_date = $4, estimated_cost = $5, actual_cost = $6, completion_status = $7, current_state = $8, manager_id = $9, budget = $10 
		WHERE campaign_id = $11;
	`
	_, err := r.db.ExecContext(r.ctx, query, campaign.ClientID, campaign.Title, campaign.StartDate, campaign.EndDate, campaign.EstimatedCost, campaign.ActualCost, campaign.CompletionStatus, campaign.CurrentState, campaign.ManagerID, campaign.Budget, campaign.CampaignID)
	if err != nil {
		log.Printf("UpdateCampaign: Failed to update campaign with ID %d: %v\n", campaign.CampaignID, err)
		return fmt.Errorf("failed to update campaign: %w", err)
	}
	return nil
}

func (s *campaignRepository) DeleteCampaign(campaignID int) error {
	log.Printf("DeleteCampaign: Deleting campaign with ID %d.", campaignID)
	query := `DELETE FROM campaigns WHERE campaign_id = $1`
	_, err := s.db.ExecContext(s.ctx, query, campaignID)
	if err != nil {
		log.Printf("DeleteCampaign: Failed to delete campaign with ID %d: %v", campaignID, err)
		return fmt.Errorf("failed to delete campaign: %w", err)
	}
	return nil
}

func (r *campaignRepository) AssignManager(campaignID, managerID int) error {
	log.Printf("AssignManager: Assigning manager with ID %d to campaign with ID %d.\n", managerID, campaignID)
	query := `
		UPDATE campaigns 
		SET manager_id = $1 
		WHERE campaign_id = $2
	`
	_, err := r.db.ExecContext(r.ctx, query, managerID, campaignID)
	if err != nil {
		log.Printf("AssignManager: Failed to assign manager with ID %d to campaign with ID %d: %v\n", managerID, campaignID, err)
		return fmt.Errorf("failed to assign manager: %w", err)
	}
	return nil
}

// CheckBudget checks the budget of a campaign.
/*func (r *campaignRepository) CheckBudget(campaignID int) (float64, error) {
	var budgetDifference float64
	query := `
		SELECT estimated_cost - actual_cost AS budget_difference
		FROM campaigns
		WHERE campaign_id = $1
	`
	err := r.db.GetContext(r.ctx, &budgetDifference, query, campaignID)
	if err != nil {
		return 0, fmt.Errorf("failed to check budget: %w", err)
	}
	return budgetDifference, nil
}*/

func (r *campaignRepository) GetCampaignsByClientID(clientID int) ([]models.Campaign, error) {
	log.Printf("GetCampaignsByClientID: Fetching campaigns for client ID %d.\n", clientID)
	var campaigns []models.Campaign
	query := "SELECT * FROM campaigns WHERE client_id = $1"
	err := r.db.SelectContext(r.ctx, &campaigns, query, clientID)
	if err != nil {
		log.Printf("GetCampaignsByClientID: Failed to fetch campaigns for client ID %d: %v\n", clientID, err)
		return nil, fmt.Errorf("failed to get campaigns by client id: %w", err)
	}
	return campaigns, nil
}
