package models

type Books struct {
	Id          int    `json:"id" form:"id" gorm:"primaryKey"`
	Title       string `json:"title" form:"title" binding:"required"`
	Author      string `json:"author" form:"author" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Stock       int    `json:"stock" form:"stock" binding:"required"`
}
