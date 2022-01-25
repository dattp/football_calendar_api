package teamstorage

import (
	teammodel "football-calendar-api/modules/team/model"
)

func (s *sqlStore) GetListTeamByIds(idTeams []int) (map[int]teammodel.Team, error) {

	result := make(map[int]teammodel.Team)

	var infoTeams []teammodel.Team

	if err := s.db.Where("id_team in (?)", idTeams).Find(&infoTeams).Error; err != nil {
		return nil, err
	}

	for _, item := range infoTeams {
		result[item.IdTeam] = item
	}

	return result, nil
}
