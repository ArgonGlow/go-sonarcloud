package hotspots

import paging "github.com/ArgonGlow/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ChangeStatusRequest Change the status of a Security Hotpot.<br/>Requires the 'Administer Security Hotspot' permission.
type ChangeStatusRequest struct {
	Comment    string `form:"comment,omitempty"`    // Comment text.
	Hotspot    string `form:"hotspot,omitempty"`    // Key of the Security Hotspot
	Resolution string `form:"resolution,omitempty"` // Resolution of the Security Hotspot when new status is REVIEWED, otherwise must not be set.
	Status     string `form:"status,omitempty"`     // New status of the Security Hotspot.
}

// SearchRequest Search for Security Hotpots.
type SearchRequest struct {
	FileUuids       string `form:"fileUuids,omitempty"`       // Comma-separated list of file uuids. Returns only hotspots found in those files. If set, 'files' must not be set.
	Files           string `form:"files,omitempty"`           // Comma-separated list of file paths. Returns only hotspots found in those files. If set, 'fileUuids' must not be set.
	Hotspots        string `form:"hotspots,omitempty"`        // Comma-separated list of Security Hotspot keys. This parameter is required unless projectKey is provided.
	OnlyMine        string `form:"onlyMine,omitempty"`        // If 'projectKey' is provided, returns only Security Hotspots assigned to the current user
	ProjectKey      string `form:"projectKey,omitempty"`      // Key of the project or application. This parameter is required unless hotspots is provided.
	Resolution      string `form:"resolution,omitempty"`      // If 'projectKey' is provided and if status is 'REVIEWED', only Security Hotspots with the specified resolution are returned.
	SinceLeakPeriod string `form:"sinceLeakPeriod,omitempty"` // If '%s' is provided, only Security Hotspots created since the leak period are returned.
	Status          string `form:"status,omitempty"`          // If 'projectKey' is provided, only Security Hotspots with the specified status are returned.
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Components []struct {
		Key          string `json:"key,omitempty"`
		LongName     string `json:"long_name,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Path         string `json:"path,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Hotspots []struct {
		Assignee         string  `json:"assignee,omitempty"`
		Author           string  `json:"author,omitempty"`
		Component        string  `json:"component,omitempty"`
		CreationDate     string  `json:"creation_date,omitempty"`
		Key              string  `json:"key,omitempty"`
		Line             float64 `json:"line,omitempty"`
		Message          string  `json:"message,omitempty"`
		Project          string  `json:"project,omitempty"`
		RuleKey          string  `json:"rule_key,omitempty"`
		SecurityCategory string  `json:"security_category,omitempty"`
		Status           string  `json:"status,omitempty"`
		TextRange        struct {
			EndLine     float64 `json:"end_line,omitempty"`
			EndOffset   float64 `json:"end_offset,omitempty"`
			StartLine   float64 `json:"start_line,omitempty"`
			StartOffset float64 `json:"start_offset,omitempty"`
		} `json:"text_range,omitempty"`
		UpdateDate               string `json:"update_date,omitempty"`
		VulnerabilityProbability string `json:"vulnerability_probability,omitempty"`
	} `json:"hotspots,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Components []struct {
		Key          string `json:"key,omitempty"`
		LongName     string `json:"long_name,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Path         string `json:"path,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Hotspots []struct {
		Assignee         string  `json:"assignee,omitempty"`
		Author           string  `json:"author,omitempty"`
		Component        string  `json:"component,omitempty"`
		CreationDate     string  `json:"creation_date,omitempty"`
		Key              string  `json:"key,omitempty"`
		Line             float64 `json:"line,omitempty"`
		Message          string  `json:"message,omitempty"`
		Project          string  `json:"project,omitempty"`
		RuleKey          string  `json:"rule_key,omitempty"`
		SecurityCategory string  `json:"security_category,omitempty"`
		Status           string  `json:"status,omitempty"`
		TextRange        struct {
			EndLine     float64 `json:"end_line,omitempty"`
			EndOffset   float64 `json:"end_offset,omitempty"`
			StartLine   float64 `json:"start_line,omitempty"`
			StartOffset float64 `json:"start_offset,omitempty"`
		} `json:"text_range,omitempty"`
		UpdateDate               string `json:"update_date,omitempty"`
		VulnerabilityProbability string `json:"vulnerability_probability,omitempty"`
	} `json:"hotspots,omitempty"`
}

// ShowRequest Provides the details of a Security Hotspot.
type ShowRequest struct {
	Hotspot string `form:"hotspot,omitempty"` // Key of the Security Hotspot
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Assignee        string `json:"assignee,omitempty"`
	Author          string `json:"author,omitempty"`
	CanChangeStatus bool   `json:"can_change_status,omitempty"`
	Changelog       []struct {
		Avatar       string `json:"avatar,omitempty"`
		CreationDate string `json:"creation_date,omitempty"`
		Diffs        []struct {
			Key      string `json:"key,omitempty"`
			NewValue string `json:"new_value,omitempty"`
			OldValue string `json:"old_value,omitempty"`
		} `json:"diffs,omitempty"`
		IsUserActive bool   `json:"is_user_active,omitempty"`
		User         string `json:"user,omitempty"`
		UserName     string `json:"user_name,omitempty"`
	} `json:"changelog,omitempty"`
	Comment []struct {
		CreatedAt string `json:"created_at,omitempty"`
		HtmlText  string `json:"html_text,omitempty"`
		Key       string `json:"key,omitempty"`
		Login     string `json:"login,omitempty"`
		Markdown  string `json:"markdown,omitempty"`
	} `json:"comment,omitempty"`
	Component struct {
		Key          string `json:"key,omitempty"`
		LongName     string `json:"long_name,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Path         string `json:"path,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"component,omitempty"`
	CreationDate string  `json:"creation_date,omitempty"`
	Key          string  `json:"key,omitempty"`
	Line         float64 `json:"line,omitempty"`
	Message      string  `json:"message,omitempty"`
	Project      struct {
		Key          string `json:"key,omitempty"`
		LongName     string `json:"long_name,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"project,omitempty"`
	Rule struct {
		Key                      string `json:"key,omitempty"`
		Name                     string `json:"name,omitempty"`
		SecurityCategory         string `json:"security_category,omitempty"`
		VulnerabilityProbability string `json:"vulnerability_probability,omitempty"`
	} `json:"rule,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdateDate string `json:"update_date,omitempty"`
	Users      []struct {
		Active bool   `json:"active,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}
