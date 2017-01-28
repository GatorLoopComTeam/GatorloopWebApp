package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// ProcTemp : struct to hold processor temperature values
type ProcTemp struct {
	Temperature float64 `json:"temperature"`
}

// GetRecent : gets the most recent processor temperature
func (p ProcTemp) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT proc_temp FROM gatorloop.proc_temp ORDER BY idProcTemp DESC LIMIT 1")
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
	var ret ProcTemp
	if res.Valid {
		ret = ProcTemp{res.Float64}
	} else {
		ret = ProcTemp{0}
	}
	response.WriteEntity(ret)
}
