package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// Position : struct to hold position values
type Position struct {
	Position           float64 `json:"position"`
	PositionPercentage float64 `json:"position_percentage"`
}

// GetRecent : gets the most recent position
func (p Position) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT position FROM gatorloop.wheel1speed ORDER BY idWheel1Speed DESC LIMIT 1")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret Position
	if res.Valid {
		ret = Position{res.Float64, res.Float64 / TotalTrackLength}
	} else {
		ret = Position{0, 0}
	}
	response.WriteEntity(ret)
}
