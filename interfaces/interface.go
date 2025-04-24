package interfaces

import (
	"database/sql"
	"main/campaign"
	"main/models"

	_ "github.com/go-sql-driver/mysql"
)

type CampaignPersistence interface {
	GetAllCampaigns() ([]models.Campaign, error)
	GetAllTargetingRules() ([]models.TargetingRule, error)
}

func NewSQLCampaignPersistence(db *sql.DB) *campaign.SQLCampaignPersistence {
	return &campaign.SQLCampaignPersistence{DB: db}
}
