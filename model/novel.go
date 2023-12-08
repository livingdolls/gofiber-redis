package model

import "gorm.io/gorm"

type Novel struct {
	Id 				int	 `gorm:"type:int;primary_key" json:"id"`
	Name 			string `gorm:"type:varchar(50);not_null" json:"name"`
	Author 			string `gorm:"type:varchar(50);not_null" json:"author"`
	Description 	string `gorm:"type:varchar(50);not_null" json:"description"`
	*gorm.Model
}