package singer

type Singer struct {
	SingerId int `gorm:"column:singer_id"`
	SingerName string `gorm:"column:singer_name"`
	SingerImg string `gorm:"column:singer_img"`
}
