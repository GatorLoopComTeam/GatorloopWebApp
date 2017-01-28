package api

import (
	"database/sql"
	"math"

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
	row := database.DB.QueryRow("SELECT speed FROM gatorloop.wheel1speed ORDER BY idWheel1Speed DESC LIMIT 1")
	var res sql.NullFloat64
	err := row.Scan(&res)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
		}
	}
	var wheel1Speed float64
	if res.Valid {
		wheel1Speed = res.Float64
	} else {
		wheel1Speed = 0.0
	}

	row = database.DB.QueryRow("SELECT speed FROM gatorloop.wheel2speed ORDER BY idWheel2Speed DESC LIMIT 1")
	err = row.Scan(&res)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
		}
	}
	var wheel2Speed float64
	if res.Valid {
		wheel2Speed = res.Float64
	} else {
		wheel2Speed = 0.0
	}

	response.WriteEntity(Velocity{math.Max(wheel1Speed, wheel2Speed)})
}
