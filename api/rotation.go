package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// Rotations : struct to hold rotation values
type Rotations struct {
	Roll  float64 `json:"roll"`
	Pitch float64 `json:"pitch"`
	Yaw   float64 `json:"yaw"`
}

// GetRecent : gets the most recent rotation
func (r Rotations) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT roll, pitch, yaw FROM gatorloop.Rotation ORDER BY idRotation DESC LIMIT 1")
	var roll, pitch, yaw sql.NullFloat64
	err := row.Scan(&roll, &pitch, &yaw)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("No Rows found. Returning 0.")
		} else {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret Rotations
	if roll.Valid && pitch.Valid && yaw.Valid {
		ret = Rotations{roll.Float64, pitch.Float64, yaw.Float64}
	} else {
		ret = Rotations{0, 0, 0}
	}
	response.WriteEntity(ret)
}
