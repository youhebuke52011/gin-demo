package user

type User struct {
	Id int `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age int `gorm:"column:age"`
}