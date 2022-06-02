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

type AddGroupParams struct {
	Group     string `json:"group"`
	MainGroup string `json:"maingroup"`
	Country   string `json:"country"`
	StateProv string `json:"stateprov"`
	City      string `json:"city"`
	Postcode  string `json:"postcode"`
	Addr1     string `json:"addr1"`
	Addr2     string `json:"addr2"`
}

type AddGroupResponse struct {
	Reason string `json:"reason"`
	Status int    `json:"status"`
}

type EditGroupParams struct {
	Group     string `json:"group"`
	GroupID   string `json:"groupid"`
	NewGroup  string `json:"newgroup"`
	MainGroup string `json:"maingroup"`
	Country   string `json:"country"`
	StateProv string `json:"stateprov"`
	City      string `json:"city"`
	Postcode  string `json:"postcode"`
	Addr1     string `json:"addr1"`
	Addr2     string `json:"addr2"`
}

type EditGroupResponse struct {
	Reason string `json:"reason"`
	Status int    `json:"status"`
}

type UpdateGroupParams struct {
	Ip         string `json:"ip"`
	Package    string `json:"package"`
	PackageID  string `json:"packageid"`
	OldGroup   string `json:"oldgroup"`
	OldGroupID string `json:"oldgroupid"`
	Group      string `json:"group"`
	GroupID    string `json:"groupid"`
}

type UpdateGroupResponse struct {
	Reason string `json:"reason"`
	Status int    `json:"status"`
}
