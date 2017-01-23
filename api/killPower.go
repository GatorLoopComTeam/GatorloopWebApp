package api

import (
	"bufio"
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
)

// KillPower : struct to hold kill power acknowledgment
type KillPower struct {
	KillPower bool `json:"kill_power"`
}

// KillPower : sends kill power signal to python controller
func (k KillPower) SendKillPower(request *restful.Request, response *restful.Response) {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		// handle error
		log.Error("Could not connect to kill power socket!")
		response.WriteEntity(Stop{false})
		return
	}
	fmt.Fprintf(conn, "KILL\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Error("Error getting response from kill power socket")
		response.WriteEntity(KillPower{false})
		return
	}

	if status == "KILLED\n" {
		response.WriteEntity(KillPower{true})
	} else {
		log.Error(fmt.Sprintf("Unsuccessful stop. Got %s from controller.", status))
		response.WriteEntity(KillPower{false})
	}

}
