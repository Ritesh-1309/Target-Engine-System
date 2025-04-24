package models

type Campaign struct {
	ID     string `json:"cid"`
	Name   string `json:"name"`
	Image  string `json:"img"`
	CTA    string `json:"cta"`
	Status string `json:"status"`
}

type TargetingRule struct {
	CampaignID       string
	IncludeApps      []string
	ExcludeApps      []string
	IncludeCountries []string
	ExcludeCountries []string
	IncludeOS        []string
	ExcludeOS        []string
}

type AppRequest struct {
	AppID   string
	Country string
	OS      string
}
