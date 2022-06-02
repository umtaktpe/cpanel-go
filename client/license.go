package client

import "github.com/umtaktpe/cpanel-go/model"

func (c *client) AddLicense(params *model.AddLicenseParams) (*model.AddLicenseResponse, error) {
	response := &model.AddLicenseResponse{}
	if err := c.request("GET", "/XMLlicenseAdd.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) CancelLicenseTransfer(params *model.CancelLicenseTransferParams) (*model.CancelLicenseTransferResponse, error) {
	response := &model.CancelLicenseTransferResponse{}
	if err := c.request("GET", "/XMLtransferRequest.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) ChangeLicenseIp(params *model.ChangeLicenseIpParams) (*model.ChangeLicenseIpResponse, error) {
	response := &model.ChangeLicenseIpResponse{}
	if err := c.request("GET", "/XMLtransfer.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) ChangeLicensePackage(params *model.ChangeLicensePackageParams) (*model.ChangeLicensePackageResponse, error) {
	response := &model.ChangeLicensePackageResponse{}
	if err := c.request("GET", "/XMLpackageUpdate.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) ExpireLicense(params *model.ExpireLicenseParams) (*model.ExpireLicenseResponse, error) {
	response := &model.ExpireLicenseResponse{}
	if err := c.request("GET", "/XMLlicenseExpire.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) RequestLicenseTransfer(params *model.RequestLicenseTransferParams) (*model.RequestLicenseTransferResponse, error) {
	response := &model.RequestLicenseTransferResponse{}
	if err := c.request("GET", "/XMLtransferRequest.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

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

func (c *client) ServiceCredits(params *model.ServiceCreditsParams) (*model.ServiceCreditsResponse, error) {
	response := &model.ServiceCreditsResponse{}
	if err := c.request("GET", "/XMLserviceCredit.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
