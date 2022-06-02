package model

type LicenseInfoParams struct {
	Expired   string `json:"expired"`
	GroupId   string `json:"groupid"`
	Group     string `json:"group"`
	Maxage    string `json:"maxage"`
	PackageId string `json:"packageid"`
	Package   string `json:"package"`
}

type LicenseInfoResponse struct {
	Version  float64                           `json:"version"`
	Status   int                               `json:"status"`
	Reason   string                            `json:"reason"`
	Licenses map[string]map[string]interface{} `json:"licenses"`
}

type SearchLicenseIDParams struct {
	Ip        string `json:"ip"`
	PackageId string `json:"packageid"`
	All       string `json:"all"`
}

type SearchLicenseIDResponse struct {
	Status    int    `json:"status"`
	Reason    string `json:"reason"`
	Licenseid int    `json:"licenseid"`
}

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
