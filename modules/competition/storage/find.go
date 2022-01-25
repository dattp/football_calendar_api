package storage

import (
	"football-calendar-api/modules/competition/model"
)

func (s *sqlStore) FindById(id *int) (*model.Competition, error) {
	var result model.Competition
	db := s.db
	if err := db.Where("id_competition = ?", *id).Take(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
