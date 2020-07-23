package model

type CrossDataCenterFirewallPolicy struct {
	ID                  string `json:"id"`
	Status              string `json:"status"`
	Enabled             bool   `json:"enabled"`
	SourceCidr          string `json:"sourceCidr"`
	SourceAccount       string `json:"sourceAccount"`
	SourceLocation      string `json:"sourceLocation"`
	DestinationCidr     string `json:"destinationCidr"`
	DestinationAccount  string `json:"destinationAccount"`
	DestinationLocation string `json:"destinationLocation"`
	Links               []Link `json:"links"`
}

type IntraDataCenterFirewallPolicy struct {
	ID                 string   `json:"id"`
	Status             string   `json:"status"`
	Enabled            bool     `json:"enabled"`
	Source             []string `json:"source"`
	Destination        []string `json:"destination"`
	DestinationAccount string   `json:"destinationAccount"`
	Ports              []string `json:"ports"`
	Links              []Link   `json:"links"`
}
