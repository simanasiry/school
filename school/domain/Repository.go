package domain

import "gorm.io/gorm"

type Reposiory interface {
	Create(interface{}) *gorm.DB
	Model(interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
}
