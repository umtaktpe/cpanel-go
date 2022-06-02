package client

import "github.com/umtaktpe/cpanel-go/model"

func (c *client) GetGroups(params *model.GetGroupsParams) (*model.GetGroupsResponse, error) {
	response := &model.GetGroupsResponse{}
	if err := c.request("GET", "/XMLgroupInfo.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) AddGroup(params *model.AddGroupParams) (*model.AddGroupResponse, error) {
	response := &model.AddGroupResponse{}
	if err := c.request("GET", "/XMLgroupAdd.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) EditGroup(params *model.EditGroupParams) (*model.EditGroupResponse, error) {
	response := &model.EditGroupResponse{}
	if err := c.request("GET", "/XMLgroupEdit.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) UpdateGroup(params *model.UpdateGroupParams) (*model.UpdateGroupResponse, error) {
	response := &model.UpdateGroupResponse{}
	if err := c.request("GET", "/XMLgroupUpdate.cgi", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
