package component

import "gorm.io/gorm"

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type appContext struct {
	dbConnection *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{dbConnection: db}
}

func (ctx *appContext) GetDBConnection() *gorm.DB {
	return ctx.dbConnection
}
