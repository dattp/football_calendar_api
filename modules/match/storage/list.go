package storage

import (
	"fmt"
	"football-calendar-api/modules/match/model"
)

func (s *sqlStore) GetList(condition map[string]interface{}) (*[]model.Match, error) {
	var result *[]model.Match
	if err := s.db.Where(condition).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (s *sqlStore) GetListMatchATeam(idTeam int, filter *model.Filter) ([]model.Match, error) {
	fmt.Println(filter)
	var result []model.Match
	var condition = make(map[string]interface{})

	db := s.db
	db = db.Where(
		db.Where("id_home_team=?", idTeam).
			Or("id_away_team=?", idTeam))

	if f := filter; f != nil {
		if f.IdCompetition > 0 {
			condition["id_competition"] = f.IdCompetition
		}
		if f.MatchStatus != "" {
			condition["match_status"] = f.MatchStatus
		}

		if !f.DateFrom.IsZero() {
			db = db.Where("match_datetime >= ?", f.DateFrom)
		}

		if !f.DateTo.IsZero() {
			db = db.Where("match_datetime <= ?", f.DateTo)
		}
	}

	if err := db.Where(condition).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
