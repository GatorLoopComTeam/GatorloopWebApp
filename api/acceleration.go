package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
	"github.com/gatorloopwebapp/server/constants"
)

// Acceleration : struct to hold acceleration values
type Acceleration struct {
	Val float64 `json:"acceleration"`
}

// GetRecent : gets the average of most recent accelerations
func (a Acceleration) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT AVG(tmp.Acceleration) FROM (SELECT Acceleration FROM gatorloop.Acceleration ORDER BY idAcceleration DESC LIMIT " + constants.NumEntriesToAvg + ") as tmp;")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		log.Errorf("Row scan failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	var ret Acceleration
	if res.Valid {
		ret = Acceleration{res.Float64}
	} else {
		ret = Acceleration{0}
	}
	response.WriteEntity(ret)
}
