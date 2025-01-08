package models

type Campaign struct {
	CampaignID       int           `db:"campaign_id" json:"campaign_id"`
	ClientID         int           `db:"client_id" json:"client_id"`
	Title            string        `db:"title" json:"title"`
	StartDate        string        `db:"start_date" json:"start_date"`
	EndDate          string        `db:"end_date" json:"end_date"`
	EstimatedCost    float64       `db:"estimated_cost" json:"estimated_cost"`
	ActualCost       float64       `db:"actual_cost" json:"actual_cost"`
	CompletionStatus bool          `db:"completion_status" json:"completion_status"`
	CurrentState     CampaignState `db:"current_state" json:"current_state"`
	ManagerID        int           `db:"manager_id" json:"manager_id"`
	Budget           int           `db:"budget" json:"budget"`
}

type CampaignState string

const (
	StateNotStarted CampaignState = "not started"
	StateInProgress CampaignState = "in progress"
	StateCompleted  CampaignState = "completed"
	StateCancelled  CampaignState = "cancelled"
)
