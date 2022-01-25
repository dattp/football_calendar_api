package model

import "time"

type FilerQuery struct {
	MatchStatus   string    `json:"match_status" form:"match_status"`
	IdCompetition int       `json:"id_competition" form:"id_competition"`
	DateFrom      time.Time `json:"date_from" form:"date_from" time_format:"2006-01-02"`
	DateTo        time.Time `json:"date_to" form:"date_to"`
}

type Filter struct {
	MatchStatus   string    `json:"match_status" form:"match_status"`
	IdCompetition int       `json:"id_competition" form:"id_competition"`
	DateFrom      time.Time `json:"date_from" form:"date_from" time_format:"2006-01-02"`
	DateTo        time.Time `json:"date_to" form:"date_to" time_format:"2006-01-02"`
}
