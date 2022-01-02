package entities

type UserRefreshToken struct {
	RefreshToken string
	Expired      int64
}
