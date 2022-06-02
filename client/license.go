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
