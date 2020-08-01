package clc

import (
	"fmt"
	"github.com/Molsbee/clc-term/clc/model"
)

type CLC struct {
	accountAlias string
	baseURL      string
	client       restClient
}

func New(accountAlias string) (c CLC, err error) {
	client, err := newClient()
	if err != nil {
		return
	}

	c = CLC{
		accountAlias: accountAlias,
		baseURL:      "https://api.ctl.io",
		client:       client,
	}
	return
}

func (c CLC) GetAccountAlias() string {
	return c.accountAlias
}

func (c CLC) GetDataCenters() (dataCenters []model.DataCenter) {
	c.client.Get(fmt.Sprintf("%s/v2/datacenters/%s", c.baseURL, c.accountAlias), &dataCenters)
	return
}

func (c CLC) GetDataCenter(id string) (dataCenter model.DataCenter) {
	c.client.Get(fmt.Sprintf("%s/v2/datacenters/%s/%s?groupLinks=true", c.baseURL, c.accountAlias, id), &dataCenter)
	return
}

func (c CLC) GetGroup(id string) (group model.Group) {
	c.client.Get(fmt.Sprintf("%s/v2/groups/%s/%s", c.baseURL, c.accountAlias, id), &group)
	return
}

func (c CLC) GetServer(serverName string) (server model.Server) {
	c.client.Get(fmt.Sprintf("%s/v2/servers/%s/%s", c.baseURL, c.accountAlias, serverName), &server)
	return
}

func (c CLC) GetServerCredentials(serverName string) (credentials model.Credentials) {
	c.client.Get(fmt.Sprintf("%s/v2/servers/%s/%s/credentials", c.baseURL, c.accountAlias, serverName), &credentials)
	return
}

func (c CLC) GetCrossDataCenterFirewallPolicies(dataCenter string) (policies []model.CrossDataCenterFirewallPolicy) {
	c.client.Get(fmt.Sprintf("%s/v2-experimental/crossDcFirewallPolicies/%s/%s", c.baseURL, c.accountAlias, dataCenter), &policies)
	return
}

func (c CLC) GetIntraDataCenterFirewallPolicies(dataCenter string) (policies []model.IntraDataCenterFirewallPolicy) {
	c.client.Get(fmt.Sprintf("%s/v2-experimental/firewallPolicies/%s/%s", c.baseURL, c.accountAlias, dataCenter), &policies)
	return
}
