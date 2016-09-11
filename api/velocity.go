package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
	"github.com/gatorloopwebapp/server/constants"
)

// Velocity : struct to hold velocity values
type Velocity struct {
	Velocity float64 `json:"velocity"`
}

// GetRecent : gets the average of 5 most recent velocities
func (v Velocity) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT AVG(tmp.Velocity) FROM (SELECT velocity FROM gatorloop.Velocity ORDER BY idVelocity DESC LIMIT " + constants.NumEntriesToAvg + ") as tmp;")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		log.Errorf("Row scan failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	var ret Velocity
	if res.Valid {
		ret = Velocity{res.Float64}
	} else {
		ret = Velocity{0}
	}
	response.WriteEntity(ret)
}
