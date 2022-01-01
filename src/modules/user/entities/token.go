package entities

type UserToken struct {
	RefreshToken string
	Expired      int64
}
