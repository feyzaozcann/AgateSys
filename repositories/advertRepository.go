package repositories

import (
	"agate-project/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type AdvertRepository interface {
	GetAllAdverts() ([]models.Advert, error)
	GetAdvertById(advertID int) (models.Advert, error)
	AddAdvert(advert *models.Advert) error
	DeleteAdvert(advertID int) error
	UpdateAdvert(advertID int, campaign_id int, progress *string, runDate *time.Time) error
	GetAdvertsByCampaign(campaignID int) ([]models.Advert, error)
}

type advertRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewAdvertRepository(ctx context.Context, db *sqlx.DB) AdvertRepository {
	return &advertRepository{
		db:  db,
		ctx: ctx,
	}
}

func (s *advertRepository) GetAllAdverts() ([]models.Advert, error) {
	log.Println("GetAllAdverts: Fetching all adverts.")
	query := `SELECT advert_id, campaign_id, progress, run_date FROM adverts`
	var adverts []models.Advert
	err := s.db.SelectContext(s.ctx, &adverts, query)
	if err != nil {
		log.Printf("GetAllAdverts: Failed to fetch adverts: %v", err)
		return nil, fmt.Errorf("failed to get all adverts: %w", err)
	}
	return adverts, nil
}

func (s *advertRepository) GetAdvertById(advertID int) (models.Advert, error) {
	log.Printf("GetAdvertById: Fetching advert with ID %d.", advertID)
	query := `SELECT advert_id, campaign_id, progress, run_date
			  FROM adverts
			  WHERE advert_id = $1`
	var advert models.Advert
	err := s.db.GetContext(s.ctx, &advert, query, advertID)
	if err != nil {
		log.Printf("GetAdvertById: Failed to fetch advert with ID %d: %v", advertID, err)
		return advert, fmt.Errorf("failed to get advert with id %d: %w", advertID, err)
	}
	return advert, nil
}

func (s *advertRepository) AddAdvert(advert *models.Advert) error {
	log.Printf("AddAdvert: Adding a new advert for campaign ID %d.", advert.CampaignID)
	query := `INSERT INTO adverts (campaign_id, progress, run_date) 
	VALUES ($1, $2, $3) RETURNING advert_id`
	err := s.db.GetContext(s.ctx, &advert.AdvertID, query, advert.CampaignID, advert.Progress, advert.RunDate)
	if err != nil {
		log.Printf("AddAdvert: Failed to add advert: %v", err)
		return fmt.Errorf("failed to add advert: %w", err)
	}
	return nil
}

func (s *advertRepository) DeleteAdvert(advertID int) error {
	log.Printf("DeleteAdvert: Deleting advert with ID %d.", advertID)
	query := `DELETE FROM adverts WHERE advert_id = $1`
	_, err := s.db.ExecContext(s.ctx, query, advertID)
	if err != nil {
		log.Printf("DeleteAdvert: Failed to delete advert with ID %d: %v", advertID, err)
		return fmt.Errorf("failed to delete advert: %w", err)
	}
	return nil
}

// run date ve progress tek başına da güncellenebilir
func (s *advertRepository) UpdateAdvert(advertID int, campaign_id int, progress *string, runDate *time.Time) error {
	log.Printf("UpdateAdvert: Updating advert with ID %d.", advertID)

	existingAdvert, err := s.GetAdvertById(advertID)
	if err != nil {
		log.Printf("UpdateAdvert: Failed to fetch existing advert with ID %d: %v", advertID, err)
		return err
	}

	newProgress := existingAdvert.Progress
	if progress != nil {
		newProgress = *progress
	}

	newRunDate := existingAdvert.RunDate
	if runDate != nil {
		newRunDate = *runDate
	}

	query := `
		UPDATE adverts
		   SET progress = $1, run_date = $2
		 WHERE advert_id = $3
	`
	_, err = s.db.ExecContext(s.ctx, query, newProgress, newRunDate, advertID)
	if err != nil {
		log.Printf("UpdateAdvert: Failed to update advert with ID %d: %v", advertID, err)
		return fmt.Errorf("failed to update advert: %w", err)
	}
	return nil
}

func (s *advertRepository) GetAdvertsByCampaign(campaignID int) ([]models.Advert, error) {
	log.Printf("GetAdvertsByCampaign: Fetching adverts for campaign ID %d.", campaignID)
	var adverts []models.Advert
	query := `SELECT advert_id, campaign_id, progress, run_date
			  FROM adverts
			  WHERE campaign_id = $1`
	err := s.db.SelectContext(s.ctx, &adverts, query, campaignID)
	if err != nil {
		log.Printf("GetAdvertsByCampaign: Failed to fetch adverts for campaign ID %d: %v", campaignID, err)
		return nil, fmt.Errorf("failed to get adverts for campaign id %d: %w", campaignID, err)
	}
	return adverts, nil
}
