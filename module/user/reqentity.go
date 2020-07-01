package user

type GetEntity struct {
	Id   int    `json:"id" form:"id" binding:"required,gt=18"`
}

type AddEntity struct {
	Id   int    `json:"id" form:"id" gorm:"column:id" binding:"required,gt=18"`
	Name string `json:"name" form:"name" gorm:"column:name" binding:"required"`
}
