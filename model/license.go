package model

type AddLicenseParams struct {
	Ip            string `json:"ip"`
	PackageID     string `json:"packageid"`
	GroupID       string `json:"groupid"`
	Force         string `json:"force"`
	DryRun        string `json:"dryrun"`
	MaxUsers      string `json:"maxusers"`
	ExternalNotes string `json:"externalnotes"`
}

type AddLicenseResponse struct {
	MonthlyPrice string `json:"monthly_price"`
	Status       int    `json:"status"`
	Yearly       int    `json:"yearly"`
	Reason       string `json:"reason"`
	Price        string `json:"price"`
	Promoinfo    string `json:"promoinfo"`
	Licenseid    string `json:"licenseid"`
}

type CancelLicenseTransferParams struct {
	Ip      string `json:"ip"`
	Cancel  string `json:"cancel"`
	GroupID string `json:"groupid"`
}

type CancelLicenseTransferResponse struct {
	Version float64 `json:"version"`
	Status  int     `json:"status"`
	Reason  string  `json:"reason"`
}

type ChangeLicenseIpParams struct {
	OldIP     string `json:"oldip"`
	NewIP     string `json:"newip"`
	PackageID string `json:"packageid"`
	Force     string `json:"force"`
	DryRun    string `json:"dryrun"`
}

type ChangeLicenseIpResponse struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
	Oldip  string `json:"oldip"`
	Newip  string `json:"newip"`
}

type ChangeLicensePackageParams struct {
	Ip           string `json:"ip"`
	NewPackageID string `json:"newpackageid"`
	OldPackageID string `json:"oldpackageid"`
	MaxUsers     string `json:"maxusers"`
}

type ChangeLicensePackageResponse struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
}

type ExpireLicenseParams struct {
	LiscID  string `json:"liscid"`
	Reason  string `json:"reason"`
	Expcode string `json:"expcode"`
}

type ExpireLicenseResponse struct {
	Status    int    `json:"status"`
	Reason    string `json:"reason"`
	Result    string `json:"result"`
	Licenseid string `json:"licenseid"`
}

type RequestLicenseTransferParams struct {
	GroupID   string `json:"groupid"`
	PackageID string `json:"packageid"`
	Ip        string `json:"ip"`
}

type RequestLicenseTransferResponse struct {
	Version float64 `json:"version"`
	Status  int     `json:"status"`
	Reason  string  `json:"reason"`
}

type LicenseInfoParams struct {
	Expired   string `json:"expired"`
	GroupID   string `json:"groupid"`
	Group     string `json:"group"`
	Maxage    string `json:"maxage"`
	PackageID string `json:"packageid"`
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
	PackageID string `json:"packageid"`
	All       string `json:"all"`
}

type SearchLicenseIDResponse struct {
	Status    int    `json:"status"`
	Reason    string `json:"reason"`
	Licenseid int    `json:"licenseid"`
}
