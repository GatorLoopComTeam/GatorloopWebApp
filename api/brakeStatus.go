package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// BrakeStatus : struct to hold the current state of primary and aux brakes
type BrakeStatus struct {
	PrimaryEngaged   bool `json:"primary_engaged"`
	AuxiliaryEngaged bool `json:"auxiliary_engaged"`
}

// GetRecent : gets the most recent state value
func (s BrakeStatus) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT primaryEngaged, auxiliaryEngaged FROM gatorloop.BrakeStatus ORDER BY idBrakeStatus DESC LIMIT 1")
	var resPrimaryEngaged, resAuxiliaryEngaged sql.NullBool
	err := row.Scan(&resPrimaryEngaged, &resAuxiliaryEngaged)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("No Rows found. Returning false.")
		} else {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret BrakeStatus
	if resPrimaryEngaged.Valid && resAuxiliaryEngaged.Valid {
		ret = BrakeStatus{resPrimaryEngaged.Bool, resAuxiliaryEngaged.Bool}
	} else {
		ret = BrakeStatus{false, false}
	}
	response.WriteEntity(ret)
}
