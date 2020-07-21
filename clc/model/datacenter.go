package model

type DataCenter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

func (d DataCenter) GetHardwareGroupID() string {
	for _, link := range d.Links {
		if link.REL == "group" {
			return link.ID
		}
	}

	return ""
}
