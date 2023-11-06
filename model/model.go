package model

type Respo struct {
	AlertID string `json:"alert_id"`
	Error   string `json:"error"`
}
type Service struct {
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	// Alerts      []Alerts `json:"alerts"`
}
type ResService struct {
	ServiceID   string   `json:"service_id"`
	ServiceName string   `json:"service_name"`
	Alerts      []Alerts `json:"alerts"`
}

type Alerts struct {
	AlertID        string `json:"alert_id"`
	Model          string `json:"model"`
	AlertType      string `json:"alert_type"`
	AlertTs        int64  `json:"alert_ts"`
	Severity       string `json:"severity"`
	AlertServiceID string `json:"alertServiceID"`
	TeamSlack      string `json:"team_slack"`
}
type ReqBody struct {
	AlertID     string `json:"alert_id"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Model       string `json:"model"`
	AlertType   string `json:"alert_type"`
	AlertTs     string `json:"alert_ts"`
	Severity    string `json:"severity"`
	TeamSlack   string `json:"team_slack"`
}
