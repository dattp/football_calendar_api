package competitiontransport

import (
	"fmt"
	"football-calendar-api/common"
	"football-calendar-api/component"
	"football-calendar-api/modules/competition/controller"
	"football-calendar-api/modules/competition/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCompetition(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id_competition"))
		fmt.Println(id)
		if err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(ctx.GetDBConnection())
		competitionCtrl := controller.NewGetCompetitionCtrl(store)

		data, err := competitionCtrl.GetCompetitionById(&id)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.ResponseSuccessLite(data))
		return
	}
}
