package services

import (
	"fmt"
	"log"

	"agate-project/models"
	"agate-project/repositories"
)

type ClientService interface {
	FetchAllClients() ([]models.Client, error)
	AddNewClient(client *models.Client) error
	RemoveClient(clientID int) error
	UpdateClient(clientID int, name *string, address *string, contactDetails *string) error
	GetClientByID(clientID int) (models.Client, error)
}

type clientService struct {
	repo repositories.ClientRepository
}

func NewClientService(repo repositories.ClientRepository) ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) FetchAllClients() ([]models.Client, error) {
	clients, err := s.repo.GetAllClients()
	if err != nil {
		log.Printf("FetchAllClients: Error fetching clients: %v", err)
		return nil, fmt.Errorf("fetching clients failed: %w", err)
	}
	return clients, nil
}

func (s *clientService) GetClientByID(clientID int) (models.Client, error) {
	if clientID <= 0 {
		log.Printf("GetClientByID: Invalid client ID: %d", clientID)
		return models.Client{}, fmt.Errorf("invalid client ID")
	}

	client, err := s.repo.GetClientByID(clientID)
	if err != nil {
		log.Printf("GetClientByID: Error fetching client with ID %d: %v", clientID, err)
		return models.Client{}, fmt.Errorf("failed to fetch client with ID %d: %w", clientID, err)
	}

	return client, nil
}

func (s *clientService) AddNewClient(client *models.Client) error {
	if err := s.repo.AddClient(client); err != nil {
		log.Printf("AddNewClient: Error adding client: %v", err)
		return fmt.Errorf("adding client failed: %w", err)
	}
	return nil
}

func (s *clientService) RemoveClient(clientID int) error {
	if err := s.repo.RemoveClient(clientID); err != nil {
		log.Printf("RemoveClient: Error removing client with ID %d: %v", clientID, err)
		return fmt.Errorf("failed to remove client with ID %d: %w", clientID, err)
	}
	return nil
}

func (s *clientService) UpdateClient(clientID int, name *string, address *string, contactDetails *string) error {
	if clientID <= 0 {
		log.Printf("UpdateClient: Invalid client ID: %d", clientID)
		return fmt.Errorf("invalid client ID")
	}

	if err := s.repo.UpdateClient(clientID, name, address, contactDetails); err != nil {
		log.Printf("UpdateClient: Error updating client with ID %d: %v", clientID, err)
		return fmt.Errorf("failed to update client with ID %d: %w", clientID, err)
	}

	return nil
}
