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

type StaffGradeHandlers interface {
	GetAllGrades(c *gin.Context)
	CreateGrade(c *gin.Context)
	RemoveGrade(c *gin.Context)
	UpdateGrade(c *gin.Context)
}

type staffGradeHandlers struct {
	ctx          context.Context
	gradeService services.StaffGradeService
}

func NewStaffGradeHandlers(ctx context.Context, service services.StaffGradeService) StaffGradeHandlers {
	return &staffGradeHandlers{
		ctx:          ctx,
		gradeService: service,
	}
}

func (h *staffGradeHandlers) GetAllGrades(c *gin.Context) {
	grades, err := h.gradeService.FetchAllGrades()
	if err != nil {
		log.Printf("GetAllGrades: Failed to fetch grades: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch grades"})
		return
	}
	c.JSON(http.StatusOK, grades)
}

func (h *staffGradeHandlers) CreateGrade(c *gin.Context) {
	var grade models.StaffGrade

	if err := c.ShouldBindJSON(&grade); err != nil {
		log.Printf("CreateGrade: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.gradeService.AddGrade(&grade); err != nil {
		log.Printf("CreateGrade: Failed to add grade: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add grade"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "grade added"})
}

func (h *staffGradeHandlers) RemoveGrade(c *gin.Context) {
	gradeIDStr := c.Param("id")

	gradeID, err := strconv.Atoi(gradeIDStr)
	if err != nil {
		log.Printf("RemoveGrade: Invalid grade ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid grade id"})
		return
	}

	if err := h.gradeService.RemoveGrade(gradeID); err != nil {
		log.Printf("RemoveGrade: Failed to remove grade with ID %d: %v", gradeID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "grade successfully removed"})
}

func (h *staffGradeHandlers) UpdateGrade(c *gin.Context) {
	gradeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("UpdateGrade: Invalid grade ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid grade id"})
		return
	}

	var grade models.StaffGrade
	if err := c.ShouldBindJSON(&grade); err != nil {
		log.Printf("UpdateGrade: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var gradeNamePtr *string
	if grade.GradeName != "" {
		gradeNamePtr = &grade.GradeName
	}
	var payRatePtr *int
	if grade.PayRate != 0 {
		payRatePtr = &grade.PayRate
	}

	if err := h.gradeService.UpdateGrade(gradeID, gradeNamePtr, payRatePtr); err != nil {
		log.Printf("UpdateGrade: Failed to update grade with ID %d: %v", gradeID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "grade updated"})
}
