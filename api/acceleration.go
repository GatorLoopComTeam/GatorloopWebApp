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
	row := database.DB.QueryRow("SELECT y FROM gatorloop.acc ORDER BY idAcc DESC LIMIT 1")
	var y sql.NullFloat64
	err := row.Scan(&y)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret Acceleration
	if y.Valid {
		ret = Acceleration{y.Float64}
	} else {
		ret = Acceleration{0}
	}
	response.WriteEntity(ret)
}
