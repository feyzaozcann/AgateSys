package repositories

import (
	"context"
	"fmt"
	"log"

	"agate-project/models"

	"github.com/jmoiron/sqlx"
)

type ClientRepository interface {
	GetAllClients() ([]models.Client, error)
	AddClient(client *models.Client) error
	RemoveClient(ClientID int) error
	GetClientByID(clientID int) (models.Client, error)
	UpdateClient(clientID int, name *string, address *string, contact_details *string) error
}

type clientRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewClientRepository(ctx context.Context, db *sqlx.DB) ClientRepository {
	return &clientRepository{
		db:  db,
		ctx: ctx,
	}
}

// r clientRepository'e ait bir pointer receiver
func (r *clientRepository) GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	query := "SELECT client_id, name, address, contact_details FROM clients"

	if err := r.db.SelectContext(r.ctx, &clients, query); err != nil {
		log.Printf("GetAllClients: Failed to retrieve clients: %v", err)
		return nil, fmt.Errorf("failed to retrieve clients: %w", err)
	}
	return clients, nil
}

func (r *clientRepository) AddClient(client *models.Client) error {
	query := "INSERT INTO clients (name, address, contact_details) VALUES ($1, $2, $3) RETURNING client_id"

	if err := r.db.GetContext(r.ctx, &client.ClientID, query, client.Name, client.Address, client.ContactDetails); err != nil {
		log.Printf("AddClient: Failed to add client: %v", err)
		return fmt.Errorf("failed to add client: %w", err)
	}
	return nil
}

func (r *clientRepository) RemoveClient(clientID int) error {
	query := "DELETE FROM clients WHERE client_id = $1"
	if _, err := r.db.ExecContext(r.ctx, query, clientID); err != nil {
		log.Printf("RemoveClient: Failed to delete client with ID %d: %v", clientID, err)
		return fmt.Errorf("failed to delete client with id %d: %w", clientID, err)
	}
	return nil
}

func (s *clientRepository) GetClientByID(clientID int) (models.Client, error) {
	query := `SELECT client_id, name, address, contact_details
			  FROM clients
			  WHERE client_id = $1`
	var client models.Client
	err := s.db.GetContext(s.ctx, &client, query, clientID)
	if err != nil {
		log.Printf("GetClientByID: Failed to get client with ID %d: %v", clientID, err)
		return client, fmt.Errorf("failed to get client with ID %d: %w", clientID, err)
	}
	return client, nil
}

func (r *clientRepository) UpdateClient(clientID int, name *string, address *string, contact_details *string) error {
	existingClient, err := r.GetClientByID(clientID)
	if err != nil {
		log.Printf("UpdateClient: Failed to fetch existing client with ID %d: %v", clientID, err)
		return err
	}

	// Gelen değer boş değilse güncelle, boşsa mevcut değeri koru
	newClientName := existingClient.Name
	if name != nil {
		newClientName = *name
	}

	newClientAddress := existingClient.Address
	if address != nil {
		newClientAddress = *address
	}

	newContactDetails := existingClient.ContactDetails
	if contact_details != nil {
		newContactDetails = *contact_details
	}

	// Güncelleme sorgusu
	query := `
		UPDATE clients
		   SET name = $1, address = $2, contact_details = $3
		 WHERE client_id = $4
	`
	_, err = r.db.ExecContext(r.ctx, query, newClientName, newClientAddress, newContactDetails, clientID)
	if err != nil {
		log.Printf("UpdateClient: Failed to update client with ID %d: %v", clientID, err)
		return fmt.Errorf("failed to update client with id %d: %w", clientID, err)
	}

	return nil
}
