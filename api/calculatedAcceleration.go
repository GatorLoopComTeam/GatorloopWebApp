package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// CalculatedAcceleration : struct to hold calculated acceleration values
type CalculatedAcceleration struct {
	Val float64 `json:"acceleration"`
}

// GetRecent : gets the most recent calculated acceleration
func (c CalculatedAcceleration) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT acc FROM gatorloop.calc_acc ORDER BY idCalcAcc DESC LIMIT 1")
	var acc sql.NullFloat64
	err := row.Scan(&acc)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret CalculatedAcceleration
	if acc.Valid {
		ret = CalculatedAcceleration{acc.Float64}
	} else {
		ret = CalculatedAcceleration{0}
	}
	response.WriteEntity(ret)
}
