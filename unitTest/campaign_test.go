package main

import (
	"main/models"
	"testing"

	"slices"

	"gopkg.in/go-playground/assert.v1"
)

func TestIsTargetingRuleMatch(t *testing.T) {
	tests := []struct {
		name     string
		req      models.AppRequest
		rule     models.TargetingRule
		expected bool
	}{
		{"IncludeApps match", models.AppRequest{AppID: "app1"}, models.TargetingRule{IncludeApps: []string{"app1"}}, true},
		{"IncludeApps no match", models.AppRequest{AppID: "app2"}, models.TargetingRule{IncludeApps: []string{"app1"}}, false},
		{"ExcludeApps match", models.AppRequest{AppID: "app1"}, models.TargetingRule{ExcludeApps: []string{"app1"}}, false},
		{"IncludeCountries match", models.AppRequest{Country: "US"}, models.TargetingRule{IncludeCountries: []string{"US"}}, true},
		{"IncludeCountries no match", models.AppRequest{Country: "UK"}, models.TargetingRule{IncludeCountries: []string{"US"}}, false},
		{"ExcludeCountries match", models.AppRequest{Country: "US"}, models.TargetingRule{ExcludeCountries: []string{"US"}}, false},
		{"IncludeOS match", models.AppRequest{OS: "android"}, models.TargetingRule{IncludeOS: []string{"android"}}, true},
		{"IncludeOS no match", models.AppRequest{OS: "ios"}, models.TargetingRule{IncludeOS: []string{"android"}}, false},
		{"ExcludeOS match", models.AppRequest{OS: "android"}, models.TargetingRule{ExcludeOS: []string{"android"}}, false},
		{"Empty rule = match everything", models.AppRequest{AppID: "any", Country: "any", OS: "any"}, models.TargetingRule{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isTargetingRuleMatch(tt.req, tt.rule)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMatchingCampaigns_Filtering(t *testing.T) {
	allCampaigns := []models.Campaign{
		{ID: "1", Status: "ACTIVE"},
		{ID: "2", Status: "INACTIVE"},
	}
	allRules := []models.TargetingRule{
		{CampaignID: "1", IncludeApps: []string{"app1"}},
		{CampaignID: "2", IncludeApps: []string{"app2"}},
	}

	req := models.AppRequest{AppID: "app1"}

	expected := []models.Campaign{
		{ID: "1", Status: "ACTIVE"},
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

	assert.Equal(t, expected, matchingCampaigns)
}

func isTargetingRuleMatch(req models.AppRequest, rule models.TargetingRule) bool {

	if len(rule.IncludeApps) > 0 {
		if !contains(rule.IncludeApps, req.AppID) {
			return false
		}
	} else if len(rule.ExcludeApps) > 0 {
		if contains(rule.ExcludeApps, req.AppID) {
			return false
		}
	}

	if len(rule.IncludeCountries) > 0 {
		if !contains(rule.IncludeCountries, req.Country) {
			return false
		}
	} else if len(rule.ExcludeCountries) > 0 {
		if contains(rule.ExcludeCountries, req.Country) {
			return false
		}
	}

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
