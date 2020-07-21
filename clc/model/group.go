package model

type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	LocationID   string `json:"locationId"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	ServersCount int    `json:"serversCount"`
	Groups       []Group
	Links        []Link
	ChangeInfo   ChangeInfo
	CustomFields []CustomField
}

func (g Group) GetServers() []string {
	serverNames := make([]string, 0)
	for _, link := range g.Links {
		if link.REL == "server" {
			serverNames = append(serverNames, link.ID)
		}
	}
	return serverNames
}

type ChangeInfo struct {
	CreatedDate  string `json:"createdDate"`
	CreatedBy    string `json:"createdBy"`
	ModifiedDate string `json:"modifiedDate"`
	ModifiedBy   string `json:"modifiedBy"`
}

type CustomField struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
}
