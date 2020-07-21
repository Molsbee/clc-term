package model

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Username      string   `json:"userName"`
	AccountAlias  string   `json:"accountAlias"`
	LocationAlias string   `json:"locationAlias"`
	Roles         []string `json:"roles"`
	BearerToken   string   `json:"bearerToken"`
}
