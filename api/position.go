package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
	"github.com/gatorloopwebapp/server/constants"
)

// Position : struct to hold position values
type Position struct {
	Position float64 `json:"position"`
}

// GetRecent : gets the average of most recent positions
func (p Position) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT AVG(tmp.Position) FROM (SELECT position FROM gatorloop.Position ORDER BY idPosition DESC LIMIT " + constants.NumEntriesToAvg + ") as tmp;")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		log.Errorf("Row scan failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	var ret Position
	if res.Valid {
		ret = Position{res.Float64}
	} else {
		ret = Position{0}
	}
	response.WriteEntity(ret)
}
