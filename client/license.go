package client

import "github.com/umtaktpe/cpanel-go/model"

func (c *client) LicenseInfo(params *model.LicenseInfoParams) (*model.LicenseInfoResponse, error) {
	response := &model.LicenseInfoResponse{}
	if err := c.request("GET", "/XMLlicenseInfo.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) SearchLicenseID(params *model.SearchLicenseIDParams) (*model.SearchLicenseIDResponse, error) {
	response := &model.SearchLicenseIDResponse{}
	if err := c.request("GET", "/XMLlookup.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) PackageInfo(params *model.PackageInfoParams) (*model.PackageInfoResponse, error) {
	response := &model.PackageInfoResponse{}
	if err := c.request("GET", "/XMLpackageInfo.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) GetRiskData(params *model.GetRiskDataParams) (*model.GetRiskDataResponse, error) {
	response := &model.GetRiskDataResponse{}
	if err := c.request("GET", "/XMLsecverify.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
