package api

import (
	"bufio"
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
)

// Stop : struct to hold stop acknowledgment
type Stop struct {
	Stop bool `json:"stop"`
}

// EmergencyStop : sends stop signal to python controller
func (s Stop) EmergencyStop(request *restful.Request, response *restful.Response) {
	conn, err := net.Dial("tcp", "127.0.0.1:6667")
	if err != nil {
		// handle error
		log.Error("Could not connect to emergency stop socket!")
		response.WriteEntity(Stop{false})
		return
	}
	fmt.Fprintf(conn, "EBRAKE\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Error("Error getting response from emergency stop socket")
		response.WriteEntity(Stop{false})
		return
	}

	if status == "EBRAKED\n" {
		response.WriteEntity(Stop{true})
	} else {
		log.Error(fmt.Sprintf("Unsuccessful stop. Got %s from controller.", status))
		response.WriteEntity(Stop{false})
	}

}
