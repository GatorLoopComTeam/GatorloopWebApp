package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// State : struct to hold State value
type State struct {
	State string `json:"state"`
}

// GetRecent : gets the most recent state value
func (s State) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT state FROM gatorloop.states ORDER BY idState DESC LIMIT 1")
	var resState sql.NullString
	err := row.Scan(&resState)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("No Rows found. Returning empty string.")
		} else {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret State
	if resState.Valid {
		ret = State{resState.String}
	} else {
		ret = State{""}
	}
	response.WriteEntity(ret)
}
