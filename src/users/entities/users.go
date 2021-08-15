package entities

type Users struct {
	UserID   string `gorm:"column:user_id;primary_key"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (t *Users) TableName() string {
	return "users"
}
