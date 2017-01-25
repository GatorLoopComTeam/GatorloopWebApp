package api

import (
	"database/sql"
	"math"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// PrimaryBattery : struct to hold primaryBattery values
type PrimaryBattery struct {
	Voltage   float64 `json:"voltage"`
	SOC       float64 `json:"soc"`
	Pack1Temp float64 `json:"pack1_temp"`
	Pack2Temp float64 `json:"pack2_temp"`
	Pack3Temp float64 `json:"pack3_temp"`
	AmpHours  float64 `json:"amp_hours"`
}

// GetRecent : gets the most recent PrimaryBattery values
func (p PrimaryBattery) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT v FROM gatorloop.bms ORDER BY idBMS DESC LIMIT 1")
	var resVoltage sql.NullFloat64
	err := row.Scan(&resVoltage)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Errorf("No Rows found.")
		} else {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	pack1Row := database.DB.QueryRow("SELECT idBatteryM1Temp, temp FROM gatorloop.battery_m1_temp ORDER BY idBatteryM1Temp DESC LIMIT 1")
	var resPack1 sql.NullFloat64
	var id sql.NullInt64
	err = pack1Row.Scan(&id, &resPack1)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("No Rows found.")
		} else {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	pack2Row := database.DB.QueryRow("SELECT idBatteryM2Temp, temp FROM gatorloop.battery_m2_temp ORDER BY idBatteryM2Temp DESC LIMIT 1")
	var resPack2 sql.NullFloat64
	err = pack2Row.Scan(&id, &resPack2)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("No Rows found.")
		} else {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	pack3Row := database.DB.QueryRow("SELECT idBatteryM3Temp, temp FROM gatorloop.battery_m3_temp ORDER BY idBatteryM3Temp DESC LIMIT 1")
	var resPack3 sql.NullFloat64
	err = pack3Row.Scan(&id, &resPack3)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("No Rows found.")
		} else {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	var ret PrimaryBattery
	soc := 0.0
	if resVoltage.Valid {
		soc = math.Max((resVoltage.Float64-MinVoltage)/(MaxVoltage-MinVoltage), 0)
	}
	ret = PrimaryBattery{
		validOrZero(resVoltage) / 1000,
		soc * 100,
		validOrZero(resPack1),
		validOrZero(resPack2),
		validOrZero(resPack3),
		soc * TotalAmpHours,
	}
	response.WriteEntity(ret)
}
