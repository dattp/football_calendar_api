package model

const Entity = "match"

type Match struct {
	Id            int    `json:"id" gorm:"column:id_match;"`
	IdCompetition int    `json:"id_competition" gorm:"column:id_competition;"`
	MatchDatetime string `json:"match_datetime" gorm:"column:match_datetime";`
	IdHomeTeam    int    `json:"id_home_team" gorm:"column:id_home_team;"`
	HomeTeamName  string `json:"home_team_name" gorm:"-"`
	HomeTeamCrest string `json:"home_team_crest" gorm:"-"`
	IdAwayTeam    int    `json:"id_away_team" gorm:"column:id_away_team;"`
	AwayTeamName  string `json:"away_team_name" gorm:"-"`
	AwayTeamCrest string `json:"away_team_crest" gorm:"-"`
	Stadium       string `json:"stadium" gorm:"-"`
	MatchStatus   string `json:"match_status" gorm:"column:match_status"`
	Stage         string `json:"stage" gorm:"column:stage;"`
}

func (m Match) TableName() string {
	return Entity
}
