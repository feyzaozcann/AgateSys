package models

type Client struct {
	ClientID       int    `db:"client_id" json:"client_id"`
	Name           string `db:"name" json:"name"`
	Address        string `db:"address" json:"address"`
	ContactDetails string `db:"contact_details" json:"contact_details"`
}
