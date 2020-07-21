package singer


type AddEntity struct {
	SingerId int `json:"singer_id" form:"singer_id" gorm:"column:singer_id"`
	SingerName string `json:"singer_name" form:"singer_name" gorm:"column:singer_name"`
	SingerImg string ` json:"singer_img" form:"singer_img"gorm:"column:singer_img"`
}
