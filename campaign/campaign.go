package campaign

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"main/models"
)

type SQLCampaignPersistence struct {
	DB *sql.DB
}

func (s *SQLCampaignPersistence) GetAllTargetingRules() ([]models.TargetingRule, error) {
	rows, err := s.DB.Query(`
		SELECT campaign_id, include_apps, exclude_apps, include_countries, exclude_countries, include_os, exclude_os
		FROM targeting_rules
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.TargetingRule
	for rows.Next() {
		var rule models.TargetingRule
		var incApp, excApp, incCountry, excCountry, incOS, excOS sql.NullString

		if err := rows.Scan(
			&rule.CampaignID,
			&incApp, &excApp,
			&incCountry, &excCountry,
			&incOS, &excOS,
		); err != nil {
			fmt.Println("Error scanning targeting rule row:", err)
			return nil, err
		}

		rule.IncludeApps = parseJSONStringArray(incApp.String)
		rule.ExcludeApps = parseJSONStringArray(excApp.String)
		rule.IncludeCountries = parseJSONStringArray(incCountry.String)
		rule.ExcludeCountries = parseJSONStringArray(excCountry.String)
		rule.IncludeOS = parseJSONStringArray(incOS.String)
		rule.ExcludeOS = parseJSONStringArray(excOS.String)

		rules = append(rules, rule)
	}
	fmt.Println("Targeting Rules:= ", rules)
	return rules, nil
}

func parseJSONStringArray(input string) []string {
	if input == "" {
		return []string{}
	}
	var result []string
	if err := json.Unmarshal([]byte(input), &result); err != nil {
		fmt.Printf("Error parsing JSON string array: %v\n", err)
		return []string{}
	}
	return result
}

func (s *SQLCampaignPersistence) GetAllCampaigns() ([]models.Campaign, error) {
	rows, err := s.DB.Query("SELECT id, name, image, cta, status FROM campaigns")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []models.Campaign
	for rows.Next() {
		var c models.Campaign
		if err := rows.Scan(&c.ID, &c.Name, &c.Image, &c.CTA, &c.Status); err != nil {
			fmt.Println("Error scanning campaign row:", err)
			return nil, err
		}
		campaigns = append(campaigns, c)
	}
	return campaigns, nil
}
