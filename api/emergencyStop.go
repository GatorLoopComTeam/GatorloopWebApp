package api

import restful "github.com/emicklei/go-restful"

// Stop : struct to hold stop acknowledgment
type Stop struct {
	Stop bool `json:"stop"`
}

// EmergencyStop : handles actions necessarry in the event of an emergency stop
func (s Stop) EmergencyStop(request *restful.Request, response *restful.Response) {
	// TODO : figure out what to do here
	response.WriteEntity(Stop{true})
}
