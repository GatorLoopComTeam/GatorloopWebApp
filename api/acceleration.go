package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// Acceleration : struct to hold acceleration values
type Acceleration struct {
	Val float64 `json:"acceleration"`
}

// GetRecent : gets the most recent acceleration
func (a Acceleration) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT acceleration FROM gatorloop.Acceleration ORDER BY idAcceleration DESC LIMIT 1")
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
	var ret Acceleration
	if res.Valid {
		ret = Acceleration{res.Float64}
	} else {
		ret = Acceleration{0}
	}
	response.WriteEntity(ret)
}
