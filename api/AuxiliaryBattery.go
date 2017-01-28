package api

import (
	"database/sql"
	"math"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
)

// AuxiliaryBattery : struct to hold auxiliary battery values
type AuxiliaryBattery struct {
	Voltage   float64 `json:"voltage"`
	SOC       float64 `json:"soc"`
	Pack1Temp float64 `json:"pack1_temp"`
	Pack2Temp float64 `json:"pack2_temp"`
	Pack3Temp float64 `json:"pack3_temp"`
	AmpHours  float64 `json:"amp_hours"`
}

func validOrZero(val sql.NullFloat64) float64 {
	if val.Valid {
		return val.Float64
	}
	return 0
}

func validOrErrCode(val sql.NullFloat64) float64 {
	if val.Valid {
		return val.Float64
	}
	return -65001
}

// GetRecent : gets the most recent PrimaryBattery values
func (a AuxiliaryBattery) GetRecent(request *restful.Request, response *restful.Response) {
	row := database.DB.QueryRow("SELECT vs FROM gatorloop.bms ORDER BY idBMS DESC LIMIT 1")
	var resVoltage sql.NullFloat64
	err := row.Scan(&resVoltage)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	pack1Row := database.DB.QueryRow("SELECT idBatteryA1Temp, temp FROM gatorloop.battery_a1_temp ORDER BY idBatteryA1Temp DESC LIMIT 1")
	var resPack1 sql.NullFloat64
	var id sql.NullInt64
	err = pack1Row.Scan(&id, &resPack1)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	pack2Row := database.DB.QueryRow("SELECT idBatteryA2Temp, temp FROM gatorloop.battery_a2_temp ORDER BY idBatteryA2Temp DESC LIMIT 1")
	var resPack2 sql.NullFloat64
	err = pack2Row.Scan(&id, &resPack2)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	pack3Row := database.DB.QueryRow("SELECT idBatteryA3Temp, temp FROM gatorloop.battery_a3_temp ORDER BY idBatteryA3Temp DESC LIMIT 1")
	var resPack3 sql.NullFloat64
	err = pack3Row.Scan(&id, &resPack3)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("Row scan failed. %v", err)
		}
	}

	var ret AuxiliaryBattery
	soc := 0.0
	if resVoltage.Valid {
		soc = math.Max((resVoltage.Float64-MinVoltage)/(MaxVoltage-MinVoltage), 0)
	}
	ret = AuxiliaryBattery{
		validOrZero(resVoltage) / 1000,
		soc * 100,
		validOrErrCode(resPack1),
		validOrErrCode(resPack2),
		validOrErrCode(resPack3),
		soc * TotalAmpHours,
	}
	response.WriteEntity(ret)
}
