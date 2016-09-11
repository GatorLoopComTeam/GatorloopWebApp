package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
	"github.com/gatorloopwebapp/server/constants"
)

// Temperature : struct to hold pressure values
type Temperature struct {
	Temperature float64 `json:"temperature"`
}

// GetRecent : gets the average of most recent pressures
func (t Temperature) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT AVG(tmp.Temperature) FROM (SELECT Temperature FROM gatorloop.Temperature ORDER BY idTemperature DESC LIMIT " + constants.NumEntriesToAvg + ") as tmp;")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		log.Errorf("Row scan failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	var ret Temperature
	if res.Valid {
		ret = Temperature{res.Float64}
	} else {
		ret = Temperature{0}
	}
	response.WriteEntity(ret)
}
