package entity

type Users struct {
	Id             string `gorm:"column:id"`
	UserId         string `gorm:"column:user_id"`
	Name           string `gorm:"column:name"`
	LevelId        string `gorm:"column:level_id"`
	Password       string `gorm:"column:password"`
	PersonalNumber string `gorm:"column:personal_number"`
	Email          string `gorm:"column:email"`
}

type UsersRequest struct {
	Request Users `json:"request"`
}

// TableName :
func (m *Users) TableName() string {
	return "users"
}
