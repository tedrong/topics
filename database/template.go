package database

import (
	"github.com/jinzhu/gorm"
)

var (
	CJ_Trends = "GoogleTrends"
)

type BasicDTO struct{}

//Get the database
func (b *BasicDTO) Get(flag DBFlag) *gorm.DB {
	return DBSet[flag].db
}
