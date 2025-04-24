package service

import (
	"main/interfaces"
	"main/models"
	"slices"
)

type CampaignService struct {
	Repo interfaces.CampaignPersistence
}

func NewCampaignService(repo interfaces.CampaignPersistence) *CampaignService {
	return &CampaignService{Repo: repo}
}

func (s *CampaignService) GetMatchingCampaigns(req models.AppRequest) ([]models.Campaign, error) {

	allRules, err := s.Repo.GetAllTargetingRules()
	if err != nil {
		return nil, err
	}

	allCampaigns, err := s.Repo.GetAllCampaigns()
	if err != nil {
		return nil, err
	}

	campaignMap := make(map[string]models.Campaign)
	for _, c := range allCampaigns {
		campaignMap[c.ID] = c
	}

	matchingCampaigns := []models.Campaign{}
	for _, rule := range allRules {
		if isValidRule(rule) && isTargetingRuleMatch(req, rule) {
			if campaign, found := campaignMap[rule.CampaignID]; found && campaign.Status == "ACTIVE" {
				matchingCampaigns = append(matchingCampaigns, campaign)
			}
		}
	}

	return matchingCampaigns, nil
}

func isTargetingRuleMatch(req models.AppRequest, rule models.TargetingRule) bool {
	// App
	if len(rule.IncludeApps) > 0 {
		if !contains(rule.IncludeApps, req.AppID) {
			return false
		}
	} else if len(rule.ExcludeApps) > 0 {
		if contains(rule.ExcludeApps, req.AppID) {
			return false
		}
	}

	// Country
	if len(rule.IncludeCountries) > 0 {
		if !contains(rule.IncludeCountries, req.Country) {
			return false
		}
	} else if len(rule.ExcludeCountries) > 0 {
		if contains(rule.ExcludeCountries, req.Country) {
			return false
		}
	}

	// OS
	if len(rule.IncludeOS) > 0 {
		if !contains(rule.IncludeOS, req.OS) {
			return false
		}
	} else if len(rule.ExcludeOS) > 0 {
		if contains(rule.ExcludeOS, req.OS) {
			return false
		}
	}

	return true
}

func contains(list []string, val string) bool {
	return slices.Contains(list, val)
}

func isValidRule(rule models.TargetingRule) bool {
	return !(len(rule.IncludeApps) > 0 && len(rule.ExcludeApps) > 0 ||
		len(rule.IncludeCountries) > 0 && len(rule.ExcludeCountries) > 0 ||
		len(rule.IncludeOS) > 0 && len(rule.ExcludeOS) > 0)
}
