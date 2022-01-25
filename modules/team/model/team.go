package teammodel

type Team struct {
	IdTeam   int    `json:"id_team" gorm:"column:id_team"`
	Name     string `json:"name" gorm:"colum:name"`
	Tla      string `json:"tla" gorm:"column:tla"`
	Stadium  string `json:"stadium" gorm:"column:stadium"`
	CrestUrl string `json:"crest_url" gorm:"column:crest_url"`
}

func (t *Team) TableName() string {
	return "team"
}
