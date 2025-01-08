package repositories

import (
	"agate-project/models"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type CampaignManagerRepository interface {
	GetAllCampaignManager() ([]models.CampaignManager, error)
	AddCampaignManager(staffGrade *models.CampaignManager) error
	DeleteCampaignManager(managerID int) error
	//AssignStaffToManager(staffID int, managerID int) error
	//UpdateCampaignManager(managerID int, staffID *int) error
}

type campaignManagerRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewCampaignManagerRepository(ctx context.Context, db *sqlx.DB) CampaignManagerRepository {
	return &campaignManagerRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *campaignManagerRepository) GetAllCampaignManager() ([]models.CampaignManager, error) {
	var campaignManager []models.CampaignManager
	query := "SELECT manager_id, staff_id FROM campaign_manager"

	if err := r.db.SelectContext(r.ctx, &campaignManager, query); err != nil {
		log.Printf("GetAllCampaignManager: Failed to retrieve managers: %v", err)
		return nil, fmt.Errorf("failed to retrieve managers: %w", err)
	}

	return campaignManager, nil
}

func (r *campaignManagerRepository) AddCampaignManager(staffGrade *models.CampaignManager) error {
	query := `INSERT INTO campaign_manager (staff_id) 
              VALUES ($1) RETURNING manager_id`
	err := r.db.GetContext(r.ctx, &staffGrade.ManagerID, query, staffGrade.StaffID)
	if err != nil {
		log.Printf("AddCampaignManager: Failed to add campaign manager: %v", err)
		return fmt.Errorf("failed to add campaign manager: %w", err)
	}
	return nil
}

func (r *campaignManagerRepository) DeleteCampaignManager(managerID int) error {
	query := `DELETE FROM campaign_manager WHERE manager_id = $1`
	_, err := r.db.ExecContext(r.ctx, query, managerID)
	if err != nil {
		log.Printf("DeleteCampaignManager: Failed to delete campaign manager with ID %d: %v", managerID, err)
		return fmt.Errorf("failed to delete campaign manager: %w", err)
	}
	return nil
}

/*func (r *campaignManagerRepository) AssignStaffToManager(staffID int, managerID int) error {
	query := `UPDATE staff SET manager_id = $1 WHERE staff_id = $2`
	_, err := r.db.ExecContext(r.ctx, query, managerID, staffID)
	if err != nil {
		return fmt.Errorf("failed to assign staff to manager: %w", err)
	}
	return nil
}*/

func (r *campaignManagerRepository) UpdateCampaignManager(managerID int, staffID *int) error {
	query := `UPDATE campaign_manager 
	SET staff_id = $1 
	WHERE manager_id = $2`
	_, err := r.db.ExecContext(r.ctx, query, staffID, managerID)
	if err != nil {
		log.Printf("UpdateCampaignManager: Failed to update campaign manager with ID %d: %v", managerID, err)
		return fmt.Errorf("failed to update campaign manager: %w", err)
	}
	return nil
}

/*func (r *campaignManagerRepository) UpdateCampaignManager(managerID int, staffID *int) error {
	query := `UPDATE campaign_manager
	SET staff_id = $1,
	WHERE manager_id = $2`
	_, err := r.db.ExecContext(r.ctx, query, staffID,managerID)
	if err != nil {
		return fmt.Errorf("failed to update campaign manager: %w", err)
	}
	return nil
}*/
