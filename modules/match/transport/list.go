package matchtransport

import (
	"fmt"
	"football-calendar-api/common"
	"football-calendar-api/component"
	"football-calendar-api/modules/match/controller"
	"football-calendar-api/modules/match/model"
	"football-calendar-api/modules/match/storage"
	teamstorage "football-calendar-api/modules/team/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListMatches(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("====>check")
		condition := make(map[string]interface{})

		condition["id_home_team"] = 66

		db := ctx.GetDBConnection()
		store := storage.NewStore(db)
		teamStore := teamstorage.NewStore(db)
		ctrl := controller.NewListMatches(store, teamStore)
		data, err := ctrl.GetMatches(condition)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}

func ListMatchATeam(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idTeam, err := strconv.Atoi(c.Param("id_team"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fmt.Println("====>filter: ", filter)

		db := ctx.GetDBConnection()
		store := storage.NewStore(db)
		teamStore := teamstorage.NewStore(db)
		ctrl := controller.NewListMatches(store, teamStore)
		data, err := ctrl.GetMatchesATeam(idTeam, &filter)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.ResponseSuccessLite(data))
	}
}
