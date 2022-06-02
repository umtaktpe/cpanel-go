package client

import "github.com/umtaktpe/cpanel-go/model"

func (c *client) GetGroups(params *model.GetGroupsParams) (*model.GetGroupsResponse, error) {
	response := &model.GetGroupsResponse{}
	if err := c.request("GET", "/XMLgroupInfo.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
