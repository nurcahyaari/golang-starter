package entities

type Users struct {
	UserID   string `gorm:"column:user_id;primary_key" json:"user_id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func (t *Users) TableName() string {
	return "users"
}
