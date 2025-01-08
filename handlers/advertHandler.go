package handlers

import (
	"agate-project/models"
	"agate-project/services"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AdvertHandlers interface {
	GetAllAdverts(c *gin.Context)
	GetAdvertByID(c *gin.Context)
	CreateAdvert(c *gin.Context)
	RemoveAdvert(c *gin.Context)
	UpdateAdvert(c *gin.Context)
	GetAdvertsByCampaign(c *gin.Context)
}

type advertHandlers struct {
	ctx           context.Context
	advertService services.AdvertService
}

func NewAdvertHandlers(ctx context.Context, service services.AdvertService) AdvertHandlers {
	return &advertHandlers{
		ctx:           ctx,
		advertService: service,
	}
}

func (h *advertHandlers) GetAllAdverts(c *gin.Context) {
	log.Println("GetAllAdverts: Received request to fetch all adverts.")
	adverts, err := h.advertService.FetchAllAdverts()
	if err != nil {
		log.Printf("GetAllAdverts: Failed to fetch adverts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch adverts"})
		return
	}
	log.Printf("GetAllAdverts: Successfully fetched %d adverts.", len(adverts))
	c.JSON(http.StatusOK, adverts)
}

func (h *advertHandlers) GetAdvertByID(c *gin.Context) {
	log.Println("GetAdvertByID: Received request to fetch advert by ID.")
	advertID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("GetAdvertByID: Invalid advert ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid advert ID"})
		return
	}

	advert, err := h.advertService.GetAdvertByID(advertID)
	if err != nil {
		log.Printf("GetAdvertByID: Failed to fetch advert with ID %d: %v", advertID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetAdvertByID: Successfully fetched advert with ID %d.", advertID)
	c.JSON(http.StatusOK, advert)
}

func (h *advertHandlers) CreateAdvert(c *gin.Context) {
	log.Println("CreateAdvert: Received request to create an advert.")
	var advert models.Advert

	if err := c.ShouldBindJSON(&advert); err != nil {
		log.Printf("CreateAdvert: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.advertService.AddAdvert(&advert); err != nil {
		log.Printf("CreateAdvert: Failed to create advert: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("CreateAdvert: Advert created successfully.")
	c.JSON(http.StatusCreated, gin.H{"message": "advert added successfully"})
}

func (h *advertHandlers) RemoveAdvert(c *gin.Context) {
	log.Println("RemoveAdvert: Received request to remove an advert.")
	advertID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("RemoveAdvert: Invalid advert ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid advert ID"})
		return
	}

	if err := h.advertService.RemoveAdvert(advertID); err != nil {
		log.Printf("RemoveAdvert: Failed to remove advert with ID %d: %v", advertID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("RemoveAdvert: Successfully removed advert with ID %d.", advertID)
	c.JSON(http.StatusOK, gin.H{"message": "advert deleted"})
}

func (h *advertHandlers) UpdateAdvert(c *gin.Context) {
	log.Println("UpdateAdvert: Received request to update an advert.")
	advertID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("UpdateAdvert: Invalid advert ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid advert ID"})
		return
	}

	// Body'den gelen JSON'u models.Advert ile bağla
	var advert models.Advert
	if err := c.ShouldBindJSON(&advert); err != nil {
		log.Printf("UpdateAdvert: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var progressPtr *string
	if advert.Progress != "" {
		progressPtr = &advert.Progress
	}

	var runDatePtr *time.Time
	if !advert.RunDate.IsZero() {
		runDatePtr = &advert.RunDate
	}

	if err := h.advertService.UpdateAdvert(advertID, advert.CampaignID, progressPtr, runDatePtr); err != nil {
		log.Printf("UpdateAdvert: Failed to update advert with ID %d: %v", advertID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateAdvert: Successfully updated advert with ID %d.", advertID)
	c.JSON(http.StatusOK, gin.H{"message": "advert updated"})
}

func (h *advertHandlers) GetAdvertsByCampaign(c *gin.Context) {
	log.Println("GetAdvertsByCampaign: Received request to fetch adverts by campaign ID.")
	campaignID, err := strconv.Atoi(c.Param("campaignID"))
	if err != nil {
		log.Printf("GetAdvertsByCampaign: Invalid campaign ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid campaign id"})
		return
	}

	adverts, err := h.advertService.GetAdvertsByCampaign(campaignID)
	if err != nil {
		log.Printf("GetAdvertsByCampaign: Failed to fetch adverts for campaign ID %d: %v", campaignID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetAdvertsByCampaign: Successfully fetched %d adverts for campaign ID %d.", len(adverts), campaignID)
	c.JSON(http.StatusOK, adverts)
}
