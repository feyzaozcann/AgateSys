package models

import "time"

type Advert struct {
	AdvertID   int       `db:"advert_id" json:"advert_id"`
	CampaignID int       `db:"campaign_id" json:"campaign_id"`
	Progress   string    `db:"progress" json:"progress"`
	RunDate    time.Time `db:"run_date" json:"run_date"`
}
