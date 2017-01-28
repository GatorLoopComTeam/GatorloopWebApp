package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// Temperature : struct to hold temperature values
type Temperature struct {
	Temperature float64 `json:"temperature"`
}

// GetRecent : gets the most recent temperature
func (t Temperature) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT temperature FROM gatorloop.Temperature ORDER BY idTemperature DESC LIMIT 1")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret Temperature
	if res.Valid {
		ret = Temperature{res.Float64}
	} else {
		ret = Temperature{0}
	}
	response.WriteEntity(ret)
}
