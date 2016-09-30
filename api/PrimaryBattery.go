package api

import (
	"database/sql"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// PrimaryBattery : struct to hold primaryBattery values
type PrimaryBattery struct {
	Voltage     float64 `json:"voltage"`
	SOC         float64 `json:"soc"`
	Temperature float64 `json:"temperature"`
	AmpHours    float64 `json:"amp_hours"`
}

// GetRecent : gets the most recent PrimaryBattery values
func (p PrimaryBattery) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT voltage, soc, temperature, amp_hour FROM gatorloop.PrimaryBattery ORDER BY idPrimaryBattery DESC LIMIT 1")
	var resVoltage, resSOC, resTemperature, resAmpHour sql.NullFloat64
	err := row.Scan(&resVoltage, &resSOC, &resTemperature, &resAmpHour)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("No Rows found. Returning 0.")
		} else {
			log.Errorf("Row scan failed. %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}
	var ret PrimaryBattery
	if resVoltage.Valid && resSOC.Valid && resTemperature.Valid && resAmpHour.Valid {
		ret = PrimaryBattery{resVoltage.Float64, resSOC.Float64, resTemperature.Float64, resAmpHour.Float64}
	} else {
		ret = PrimaryBattery{0, 0, 0, 0}
	}
	response.WriteEntity(ret)
}
