package repository

import "gorm.io/gorm"

type Repositories struct {
	User     UserRepository
	Category CategoryRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db),
		Category: NewCategoryRepository(db),
	}
}
