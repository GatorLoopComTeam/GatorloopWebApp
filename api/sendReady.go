package api

import (
	"bufio"
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
)

// Ready : struct to hold ready acknowledgment
type Ready struct {
	Ready bool `json:"ready"`
}

// SendReady : sends ready signal to python controller
func (r Ready) SendReady(request *restful.Request, response *restful.Response) {
	conn, err := net.Dial("tcp", "127.0.0.1:6668")
	if err != nil {
		// handle error
		log.Error("Could not connect to ready socket!")
		response.WriteEntity(Ready{false})
		return
	}
	fmt.Fprintf(conn, "READY\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Error("Error getting response from ready socket")
		response.WriteEntity(Ready{false})
		return
	}

	if status == "READY\n" {
		response.WriteEntity(Ready{true})
	} else {
		log.Error(fmt.Sprintf("Unsuccessful ready. Got %s from controller.", status))
		response.WriteEntity(Ready{false})
	}

}
