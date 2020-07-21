package model

type Link struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	REL  string `json:"rel"`
	HREF string `json:"href"`
}
