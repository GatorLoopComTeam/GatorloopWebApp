package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
	"github.com/gatorloopwebapp/server/constants"
)

// Pressure : struct to hold pressure values
type Pressure struct {
	Pressure float64 `json:"pressure"`
}

// GetRecent : gets the average of most recent pressures
func (p Pressure) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT AVG(tmp.Pressure) FROM (SELECT Pressure FROM gatorloop.Pressure ORDER BY idPressure DESC LIMIT " + constants.NumEntriesToAvg + ") as tmp;")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		log.Errorf("Row scan failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	var ret Pressure
	if res.Valid {
		ret = Pressure{res.Float64}
	} else {
		ret = Pressure{0}
	}
	response.WriteEntity(ret)
}
