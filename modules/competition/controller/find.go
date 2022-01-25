package controller

import (
	"football-calendar-api/common"
	"football-calendar-api/modules/competition/model"
)

type GetCompetitionStore interface {
	FindById(id *int) (*model.Competition, error)
}

type getCompetitionCtrl struct {
	store GetCompetitionStore
}

func NewGetCompetitionCtrl(store GetCompetitionStore) *getCompetitionCtrl {
	return &getCompetitionCtrl{store: store}
}

func (ctrl *getCompetitionCtrl) GetCompetitionById(id *int) (*model.Competition, error) {
	data, err := ctrl.store.FindById(id)
	//fmt.Println(data)
	//var myArr []int
	//fmt.Println(myArr[2])
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	if data == nil {
		return nil, common.ErrInvalidRequest(nil)
	}
	return data, err
}
