package model

type PackageInfoParams struct {
	Expand string `json:"expand"`
}

type PackageInfoResponse struct {
	Packages map[string]string `json:"packages"`
	Version  float64           `json:"version"`
	Status   int               `json:"status"`
	Reason   string            `json:"reason"`
}

type GetRiskDataParams struct {
	Ip string `json:"ip"`
}

type GetRiskDataResponse struct {
	Version                   int    `json:"version"`
	Status                    int    `json:"status"`
	Reason                    string `json:"reason"`
	Clientreportingurl        string `json:"clientreportingurl"`
	RiskscoreMainScore        int    `json:"riskscore.main.score"`
	RiskscoreDirectorderScore int    `json:"riskscore.directorder.score"`
	RiskscoreAggregateScore   int    `json:"riskscore.aggregate.score"`
}

type ServiceCreditsParams struct {
	LiscID string `json:"liscid"`
}

type ServiceCreditsResponse struct {
	Reason  string `json:"reason"`
	Status  int    `json:"status"`
	Results []struct {
		LicenseID    string `json:"license_id"`
		Status       string `json:"status"`
		StatusReason string `json:"status_reason"`
	} `json:"results"`
}
