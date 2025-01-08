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

type StaffHandlers interface {
	GetStaff(c *gin.Context)
	GetStaffByID(c *gin.Context)
	CreateStaff(c *gin.Context)
	RemoveStaff(c *gin.Context)
	UpdateStaff(c *gin.Context)
}

type staffHandlers struct {
	ctx         context.Context
	userService services.StaffService
}

func NewStaffHandlers(ctx context.Context, service services.StaffService) StaffHandlers {
	return &staffHandlers{
		ctx:         ctx,
		userService: service,
	}
}

func (h *staffHandlers) GetStaff(c *gin.Context) {
	staff, err := h.userService.FetchAllStaff()
	if err != nil {
		log.Printf("GetStaff: Failed to fetch staff: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch staff"})
		return
	}
	c.JSON(http.StatusOK, staff)
}

func (h *staffHandlers) GetStaffByID(c *gin.Context) {
	staffID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("GetStaffByID: Invalid staff ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid staff ID"})
		return
	}

	staff, err := h.userService.GetStaffByID(staffID)
	if err != nil {
		log.Printf("GetStaffByID: Failed to retrieve staff with ID %d: %v", staffID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, staff)
}

func (h *staffHandlers) CreateStaff(c *gin.Context) {
	var staff models.Staff

	if err := c.ShouldBindJSON(&staff); err != nil {
		log.Printf("CreateStaff: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.userService.AddStaff(&staff); err != nil {
		log.Printf("CreateStaff: Failed to add staff: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add staff"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "staff added"})
}

func (h *staffHandlers) RemoveStaff(c *gin.Context) {
	staffIDStr := c.Param("id")

	staffID, err := strconv.Atoi(staffIDStr)
	if err != nil {
		log.Printf("RemoveStaff: Invalid staff ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid staff id"})
		return
	}

	if err := h.userService.RemoveStaff(staffID); err != nil {
		log.Printf("RemoveStaff: Failed to delete staff with ID %d: %v", staffID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "staff deleted"})
}

func (h *staffHandlers) UpdateStaff(c *gin.Context) {
	staffID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("UpdateStaff: Invalid staff ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid staff id"})
		return
	}

	var updatedDetails models.Staff
	if err := c.ShouldBindJSON(&updatedDetails); err != nil {
		log.Printf("UpdateStaff: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.userService.UpdateStaff(staffID, &updatedDetails); err != nil {
		log.Printf("UpdateStaff: Failed to update staff with ID %d: %v", staffID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "staff updated"})
}
