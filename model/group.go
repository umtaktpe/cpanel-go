package model

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
