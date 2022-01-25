package teamstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
