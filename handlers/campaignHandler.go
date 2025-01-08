package handlers

import (
	"agate-project/models"
	"agate-project/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandlers interface {
	CreateCampaign(c *gin.Context)
	GetCampaignByID(c *gin.Context)
	UpdateCampaign(c *gin.Context)
	RemoveCampaign(c *gin.Context)
	AssignManager(c *gin.Context)
	GetAllCampaigns(c *gin.Context)
	//CheckBudget(c *gin.Context)
	GetCampaignsByClientID(c *gin.Context)
}

type campaignHandlers struct {
	service services.CampaignService
}

func NewCampaignHandlers(service services.CampaignService) CampaignHandlers {
	return &campaignHandlers{service: service}
}

func (h *campaignHandlers) CreateCampaign(c *gin.Context) {
	log.Println("CreateCampaign: Received request to create a campaign.")
	var campaign models.Campaign

	if err := c.ShouldBindJSON(&campaign); err != nil {
		log.Printf("CreateCampaign: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.service.CreateCampaign(campaign); err != nil {
		log.Printf("CreateCampaign: Failed to create campaign: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "campaign created"})
}

func (h *campaignHandlers) GetCampaignByID(c *gin.Context) {
	log.Println("GetCampaignByID: Received request to fetch campaign by ID.")
	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("GetCampaignByID: Invalid campaign ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid campaign ID"})
		return
	}

	campaign, err := h.service.GetCampaignByID(campaignID)
	if err != nil {
		log.Printf("GetCampaignByID: Failed to fetch campaign by ID %d: %v", campaignID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func (h *campaignHandlers) UpdateCampaign(c *gin.Context) {
	log.Println("UpdateCampaign: Received request to update a campaign.")

	// URL'den campaign_id'yi al
	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("UpdateCampaign: Invalid campaign ID in URL: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid campaign ID"})
		return
	}
	var campaign models.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		log.Printf("UpdateCampaign: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	campaign.CampaignID = campaignID
	if err := h.service.UpdateCampaign(campaign); err != nil {
		log.Printf("UpdateCampaign: Failed to update campaign with ID %d: %v", campaign.CampaignID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "campaign updated"})
}

func (h *campaignHandlers) RemoveCampaign(c *gin.Context) {
	log.Println("RemoveCampaign: Received request to remove an campaign.")
	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("RemoveCampaign: Invalid campaign ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid campaign ID"})
		return
	}

	if err := h.service.RemoveCampaign(campaignID); err != nil {
		log.Printf("RemoveAdvert: Failed to remove campaign with ID %d: %v", campaignID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "campaign deleted"})
}

func (h *campaignHandlers) AssignManager(c *gin.Context) {
	log.Println("AssignManager: Received request to assign a manager to a campaign.")
	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("AssignManager: Invalid campaign ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid campaign id"})
		return
	}

	managerID, err := strconv.Atoi(c.Param("managerID"))
	if err != nil {
		log.Printf("AssignManager: Invalid manager ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid manager id"})
		return
	}

	if err := h.service.AssignManager(campaignID, managerID); err != nil {
		log.Printf("AssignManager: Failed to assign manager with ID %d to campaign with ID %d: %v", managerID, campaignID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "manager assigned to campaign"})
}

func (h *campaignHandlers) GetAllCampaigns(c *gin.Context) {
	log.Println("GetAllCampaigns: Received request to fetch all campaigns.")
	campaigns, err := h.service.FetchAllCampaigns()
	if err != nil {
		log.Printf("GetAllCampaigns: Failed to fetch campaigns: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch campaigns"})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

// CheckBudget handles checking the budget for a campaign
/*func (h *campaignHandlers) CheckBudget(c *gin.Context) {
	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	budgetDifference, err := h.service.CheckBudget(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"budget_difference": budgetDifference})
}*/

func (h *campaignHandlers) GetCampaignsByClientID(c *gin.Context) {
	log.Println("GetCampaignsByClientID: Received request to fetch campaigns by client ID.")
	clientID, err := strconv.Atoi(c.Param("clientID"))
	if err != nil {
		log.Printf("GetCampaignsByClientID: Invalid client ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	campaigns, err := h.service.GetCampaignsByClientID(clientID)
	if err != nil {
		log.Printf("GetCampaignsByClientID: Failed to fetch campaigns for client ID %d: %v", clientID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}
