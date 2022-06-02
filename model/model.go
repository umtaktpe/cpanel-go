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

type GetGroupsParams struct {
	Expand string `json:"expand"`
}

type GetGroupsResponse struct {
	Version float64           `json:"version"`
	Status  int               `json:"status"`
	Groups  map[string]string `json:"groups"`
	Group   []struct {
		Country   interface{} `json:"country"`
		Name      string      `json:"name"`
		Addr2     interface{} `json:"addr2"`
		City      interface{} `json:"city"`
		Stateprov interface{} `json:"stateprov"`
		Addr1     interface{} `json:"addr1"`
		Groupid   string      `json:"groupid"`
		Postcode  interface{} `json:"postcode"`
		Maingroup int         `json:"maingroup"`
	} `json:"group"`
	Reason string `json:"reason"`
}
