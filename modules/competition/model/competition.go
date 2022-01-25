package model

const EntityName = "competition"

type Competition struct {
	IdCompetiton int    `json:"id_competition" gorm:"column:id_competition;"`
	IdArea       int    `json:"id_area" gorm:"column:id_area;"`
	Name         string `json:"name" gorm:"column:name;"`
	Code         string `json:"code" gorm:"column:code;"`
	createdAt    string `json:"created_at" gorm:"column:created_at;"`
}

func (c *Competition) TableName() string {
	return EntityName
}
