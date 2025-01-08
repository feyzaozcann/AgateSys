package models

type CampaignManager struct {
	ManagerID int `db:"manager_id" json:"manager_id"`
	StaffID   int `db:"staff_id" json:"staff_id"`
}
