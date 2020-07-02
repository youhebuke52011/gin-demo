package user

import "gin-demo/common"

type GetEntity struct {
	common.PublicArg
	Id   int    `json:"id" form:"id" binding:"required"`
}

type AddEntity struct {
	common.PublicArg
	Id   int    `json:"id" form:"id" gorm:"column:id" binding:"required"`
	Name string `json:"name" form:"name" gorm:"column:name" binding:"required"`
	Age int `json:"age" form:"id" gorm:"column:age" binding:"required,gt=18,enum=19-20-21"`
}
