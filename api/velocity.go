package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// Velocity : struct to hold velocity values
type Velocity struct {
	Val float64 `json:"val"`
}

type avgVelocity struct {
	AvgVelocity float64 `json:"avgVelocity"`
}

// GetRecent : gets the average of 5 most recent velocities
func (v Velocity) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT AVG(tmp.Velocity) FROM (SELECT velocity FROM gatorloop.Velocity ORDER BY idVelocity DESC LIMIT 5) as tmp;")
	var res float64
	err := row.Scan(&res)
	if err != nil {
		log.Errorf("Row scan failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	ret := avgVelocity{res}
	response.WriteEntity(ret)
}

// CreateVelocity : create a new velocity entry
func (v Velocity) CreateVelocity(request *restful.Request, response *restful.Response) {
	vel := new(Velocity)
	err := request.ReadEntity(&vel)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}
	stmt, err := database.DB.Prepare("INSERT INTO Velocity (velocity) VALUES(?)")
	if err != nil {
		log.Errorf("DB prepare failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		err = stmt.Close()
		if err != nil {
			log.Errorf("DB stmt close failed. %v", err)
		}
		return
	}

	_, err = stmt.Exec(vel.Val)
	if err != nil {
		log.Errorf("DB stmt exec failed. %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteEntity(vel)
}
