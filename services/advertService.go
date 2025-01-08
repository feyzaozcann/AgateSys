package services

import (
	"agate-project/models"
	"agate-project/repositories"
	"fmt"
	"log"
	"time"
)

type AdvertService interface {
	FetchAllAdverts() ([]models.Advert, error)
	GetAdvertByID(advertID int) (models.Advert, error)
	AddAdvert(advert *models.Advert) error
	RemoveAdvert(advertID int) error
	UpdateAdvert(advertID int, campaignID int, progress *string, runDate *time.Time) error
	GetAdvertsByCampaign(campaignID int) ([]models.Advert, error)
}

type advertService struct {
	repo repositories.AdvertRepository
}

func NewAdvertService(repo repositories.AdvertRepository) AdvertService {
	return &advertService{repo: repo}
}

func (s *advertService) FetchAllAdverts() ([]models.Advert, error) {
	log.Println("FetchAllAdverts: Fetching all adverts.")
	adverts, err := s.repo.GetAllAdverts()
	if err != nil {
		log.Printf("FetchAllAdverts: Failed to fetch adverts: %v", err)
		return nil, fmt.Errorf("fetching adverts failed: %w", err)
	}
	return adverts, nil
}

func (s *advertService) GetAdvertByID(advertID int) (models.Advert, error) {
	log.Printf("GetAdvertByID: Fetching advert with ID %d.", advertID)
	advert, err := s.repo.GetAdvertById(advertID)
	if err != nil {
		log.Printf("GetAdvertByID: Failed to fetch advert with ID %d: %v", advertID, err)
		return advert, fmt.Errorf("fetching advert with id %d failed: %w", advertID, err)
	}
	return advert, nil
}

func (s *advertService) AddAdvert(advert *models.Advert) error {
	log.Printf("AddAdvert: Adding a new advert for campaign ID %d.", advert.CampaignID)
	if advert.CampaignID == 0 || advert.Progress == "" {
		log.Println("AddAdvert: Invalid advert data: campaign ID or progress is missing.")
		return fmt.Errorf("invalid advert data: campaign ID or progress is missing")
	}

	if err := s.repo.AddAdvert(advert); err != nil {
		log.Printf("AddAdvert: Error adding advert: %v", err)
		return fmt.Errorf("adding advert failed: %w", err)
	}
	return nil
}

func (s *advertService) RemoveAdvert(advertID int) error {
	log.Printf("RemoveAdvert: Removing advert with ID %d.", advertID)
	if err := s.repo.DeleteAdvert(advertID); err != nil {
		log.Printf("RemoveAdvert: Failed to remove advert with ID %d: %v", advertID, err)
		return fmt.Errorf("failed to remove advert with id %d: %w", advertID, err)
	}
	return nil
}

func (s *advertService) UpdateAdvert(advertID int, campaignID int, progress *string, runDate *time.Time) error {
	log.Printf("UpdateAdvert: Updating advert with ID %d.", advertID)
	if err := s.repo.UpdateAdvert(advertID, campaignID, progress, runDate); err != nil {
		log.Printf("UpdateAdvert: Failed to update advert with ID %d: %v", advertID, err)
		return fmt.Errorf("failed to update advert with id %d: %w", advertID, err)
	}
	return nil
}

func (s *advertService) GetAdvertsByCampaign(campaignID int) ([]models.Advert, error) {
	log.Printf("GetAdvertsByCampaign: Fetching adverts for campaign ID %d.", campaignID)
	adverts, err := s.repo.GetAdvertsByCampaign(campaignID)
	if err != nil {
		log.Printf("GetAdvertsByCampaign: Failed to fetch adverts for campaign ID %d: %v", campaignID, err)
		return nil, fmt.Errorf("fetching adverts for campaign id %d failed: %w", campaignID, err)
	}
	return adverts, nil
}
