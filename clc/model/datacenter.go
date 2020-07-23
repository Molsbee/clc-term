package model

import "strings"

type DataCenter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

func (d DataCenter) SanitizedName() string {
	name := d.Name
	if strings.Contains(name, "[") {
		name = strings.Replace(name, "[", "(", -1)
	}
	if strings.Contains(name, "]") {
		name = strings.Replace(name, "]", ")", -1)
	}
	return name
}

func (d DataCenter) GetHardwareGroupID() string {
	for _, link := range d.Links {
		if link.REL == "group" {
			return link.ID
		}
	}

	return ""
}
