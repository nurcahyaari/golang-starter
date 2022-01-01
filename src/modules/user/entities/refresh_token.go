package entities

type UserRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
	Expired      int64  `json:"expired"`
}
