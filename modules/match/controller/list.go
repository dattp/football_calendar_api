package controller

import (
	"football-calendar-api/modules/match/model"
	teammodel "football-calendar-api/modules/team/model"
)

type MatchStore interface {
	GetList(condition map[string]interface{}) (*[]model.Match, error)
	GetListMatchATeam(idTeam int, filter *model.Filter) ([]model.Match, error)
}

type TeamStore interface {
	GetListTeamByIds(idTeams []int) (map[int]teammodel.Team, error)
}

type ListMatch struct {
	store     MatchStore
	teamStore TeamStore
}

func NewListMatches(store MatchStore, teamStore TeamStore) *ListMatch {
	return &ListMatch{store: store, teamStore: teamStore}
}

func (ctrl *ListMatch) GetMatches(condition map[string]interface{}) (*[]model.Match, error) {
	data, err := ctrl.store.GetList(condition)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ctrl *ListMatch) GetMatchesATeam(idTeam int, filter *model.Filter) ([]model.Match, error) {
	data, err := ctrl.store.GetListMatchATeam(idTeam, filter)
	if err != nil {
		return nil, err
	}

	idTeams := make([]int, len(data))
	for i := range data {
		if data[i].IdHomeTeam != idTeam {
			idTeams[i] = data[i].IdHomeTeam
		}
		if data[i].IdAwayTeam != idTeam {
			idTeams[i] = data[i].IdAwayTeam
		}
	}
	idTeams = append(idTeams, idTeam)

	mapInfoTeams, err := ctrl.teamStore.GetListTeamByIds(idTeams)

	if err != nil {
		return nil, err
	}

	for i, item := range data {
		homeTeam := mapInfoTeams[item.IdHomeTeam]
		awayTeam := mapInfoTeams[item.IdAwayTeam]

		data[i].HomeTeamName = homeTeam.Name
		data[i].HomeTeamCrest = homeTeam.CrestUrl
		data[i].AwayTeamName = awayTeam.Name
		data[i].AwayTeamCrest = awayTeam.CrestUrl
		data[i].Stadium = homeTeam.Stadium

	}

	return data, nil
}
