package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"agate-project/models"
	"agate-project/services"

	"github.com/gin-gonic/gin"
)

type ClientHandlers interface {
	GetClients(c *gin.Context)
	CreateClient(c *gin.Context)
	RemoveClient(c *gin.Context)
	UpdateClient(c *gin.Context)
	GetClientByID(c *gin.Context)
}

type clientHandlers struct {
	ctx         context.Context
	userService services.ClientService
}

func NewClientHandlers(ctx context.Context, service services.ClientService) ClientHandlers {
	return &clientHandlers{
		ctx:         ctx,
		userService: service,
	}
}

func (h *clientHandlers) GetClients(c *gin.Context) {
	clients, err := h.userService.FetchAllClients()
	if err != nil {
		log.Printf("GetClients: Failed to fetch clients: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (h *clientHandlers) GetClientByID(c *gin.Context) {
	clientID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("GetClientByID: Invalid client ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client ID"})
		return
	}

	client, err := h.userService.GetClientByID(clientID)
	if err != nil {
		log.Printf("GetClientByID: Failed to fetch client with ID %d: %v", clientID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": client})
}

func (h *clientHandlers) CreateClient(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		log.Printf("CreateClient: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.userService.AddNewClient(&client); err != nil {
		log.Printf("CreateClient: Failed to add client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add client"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "client added"})
}

func (h *clientHandlers) RemoveClient(c *gin.Context) {
	clientIDStr := c.Param("id")

	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		log.Printf("RemoveClient: Invalid client ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client ID"})
		return
	}

	if err := h.userService.RemoveClient(clientID); err != nil {
		log.Printf("RemoveClient: Failed to delete client with ID %d: %v", clientID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "client deleted"})
}

func (h *clientHandlers) UpdateClient(c *gin.Context) {
	clientID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("UpdateClient: Invalid client ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client ID"})
		return
	}

	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		log.Printf("UpdateClient: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var namePtr *string
	if client.Name != "" {
		namePtr = &client.Name
	}

	var addressPtr *string
	if client.Address != "" {
		addressPtr = &client.Address
	}

	var contactDetailsPtr *string
	if client.ContactDetails != "" {
		contactDetailsPtr = &client.ContactDetails
	}

	if err := h.userService.UpdateClient(clientID, namePtr, addressPtr, contactDetailsPtr); err != nil {
		log.Printf("UpdateClient: Failed to update client with ID %d: %v", clientID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "client updated"})
}
