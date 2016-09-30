package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// Velocity : struct to hold velocity values
type Velocity struct {
	Velocity float64 `json:"velocity"`
}

// GetRecent : gets the most recent velocity
func (v Velocity) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT velocity FROM gatorloop.Velocity ORDER BY idVelocity DESC LIMIT 1")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("No Rows found. Returning 0.")
		} else {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret Velocity
	if res.Valid {
		ret = Velocity{res.Float64}
	} else {
		ret = Velocity{0}
	}
	response.WriteEntity(ret)
}
