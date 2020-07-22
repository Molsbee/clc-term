package clc

import (
	"fmt"
	"github.com/Molsbee/clc-term/clc/model"
)

type CLC struct {
	accountAlias string
	baseURL      string
	client       client
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

func (c CLC) GetDataCenters() []model.DataCenter {
	url := fmt.Sprintf("%s/v2/datacenters/%s", c.baseURL, c.accountAlias)
	var dataCenters []model.DataCenter
	c.client.Get(url, &dataCenters)
	return dataCenters
}

func (c CLC) GetDataCenter(id string) model.DataCenter {
	url := fmt.Sprintf("%s/v2/datacenters/%s/%s?groupLinks=true", c.baseURL, c.accountAlias, id)
	dataCenter := model.DataCenter{}
	c.client.Get(url, &dataCenter)
	return dataCenter
}

func (c CLC) GetGroup(id string) model.Group {
	url := fmt.Sprintf("%s/v2/groups/%s/%s", c.baseURL, c.accountAlias, id)
	group := model.Group{}
	c.client.Get(url, &group)
	return group
}

func (c CLC) GetServer(serverName string) model.Server {
	url := fmt.Sprintf("%s/v2/servers/%s/%s", c.baseURL, c.accountAlias, serverName)
	server := model.Server{}
	c.client.Get(url, &server)
	return server
}

func (c CLC) GetServerCredentials(serverName string) model.Credentials {
	url := fmt.Sprintf("%s/v2/servers/%s/%s/credentials", c.baseURL, c.accountAlias, serverName)
	credentials := model.Credentials{}
	c.client.Get(url, &credentials)
	return credentials
}
