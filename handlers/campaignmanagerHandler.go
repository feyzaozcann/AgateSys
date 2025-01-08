package handlers

import (
	"agate-project/models"
	"agate-project/services"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignManagerHandlers interface {
	GetAllManagers(c *gin.Context)
	CreateManager(c *gin.Context)
	DeleteManager(c *gin.Context)
	//UpdateManager(c *gin.Context)
	//AssignStaffToManager(c *gin.Context)
}

type campaignManagerHandlers struct {
	ctx            context.Context
	managerService services.CampaignManagerService
}

func NewCampaignManagerHandlers(ctx context.Context, service services.CampaignManagerService) CampaignManagerHandlers {
	return &campaignManagerHandlers{
		ctx:            ctx,
		managerService: service,
	}
}

func (h *campaignManagerHandlers) GetAllManagers(c *gin.Context) {
	managers, err := h.managerService.GetAllCampaignManager()
	if err != nil {
		log.Printf("GetAllManagers: Failed to fetch campaign managers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch campaign managers"})
		return
	}
	c.JSON(http.StatusOK, managers)
}

func (h *campaignManagerHandlers) CreateManager(c *gin.Context) {
	var manager models.CampaignManager

	if err := c.ShouldBindJSON(&manager); err != nil {
		log.Printf("CreateManager: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.managerService.AddCampaignManager(&manager); err != nil {
		log.Printf("CreateManager: Failed to add campaign manager: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add campaign manager"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "campaign manager added"})
}

func (h *campaignManagerHandlers) DeleteManager(c *gin.Context) {
	managerIDStr := c.Param("id")

	managerID, err := strconv.Atoi(managerIDStr)
	if err != nil {
		log.Printf("DeleteManager: Invalid manager ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid manager id"})
		return
	}

	if err := h.managerService.DeleteCampaignManager(managerID); err != nil {
		log.Printf("DeleteManager: Failed to delete campaign manager with ID %d: %v", managerID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "campaign manager deleted"})
}

/*func (h *campaignManagerHandlers) AssignStaffToManager(c *gin.Context) {
	var assignRequest struct {
		StaffID   int `json:"staff_id"`
		ManagerID int `json:"manager_id"`
	}

	if err := c.ShouldBindJSON(&assignRequest); err != nil {
		log.Printf("AssignStaffToManager: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.managerService.AssignStaffToManager(assignRequest.StaffID, assignRequest.ManagerID); err != nil {
		log.Printf("AssignStaffToManager: Failed to assign staff to manager: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "staff assigned to manager"})
}*/
