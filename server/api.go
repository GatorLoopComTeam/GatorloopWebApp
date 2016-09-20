package server

import (
	"net/http"
	"path"

	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/gatorloopwebapp/api"
	"github.com/gatorloopwebapp/server/constants"
)

// RegisterFileServer : registers file serving routes with go-restful
func RegisterFileServer(path string, container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/static").To(staticFromQueryParam).
		Doc("Serves static files").
		Param(ws.PathParameter("resource", "the path to a resource within /static")))

	restful.Add(ws)
}

// RegisterAPI : registers api routes with go-restful
func RegisterAPI(apiPath string, container *restful.Container) {

	ws := new(restful.WebService)

	ws.Path(apiPath).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	var v api.Velocity
	ws.Route(ws.GET("/velocity").
		To(v.GetRecent).
		Doc("Get an average of the last " + constants.NumEntriesToAvg + " velocities."))

	var a api.Acceleration
	ws.Route(ws.GET("/acceleration").
		To(a.GetRecent).
		Doc("Get an average of the last " + constants.NumEntriesToAvg + " accelerations."))

	var p api.Position
	ws.Route(ws.GET("/position").
		To(p.GetRecent).
		Doc("Get an average of the last " + constants.NumEntriesToAvg + " positions."))

	var r api.Rotations
	ws.Route(ws.GET("/rotation").
		To(r.GetRecent).
		Doc("Get an average of the last " + constants.NumEntriesToAvg + " rotations."))

	var t api.Temperature
	ws.Route(ws.GET("/temperature").
		To(t.GetRecent).
		Doc("Get an average of the last " + constants.NumEntriesToAvg + " temperatures."))

	var pr api.Pressure
	ws.Route(ws.GET("/pressure").
		To(pr.GetRecent).
		Doc("Get an average of the last " + constants.NumEntriesToAvg + " pressures."))

	var s api.Stop
	ws.Route(ws.GET("/stop").
		To(s.EmergencyStop).
		Doc("Initiates an emergency stop of the pod."))

	restful.Add(ws)

	config := swagger.Config{
		WebServices:    restful.DefaultContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/Users/gavin/gocode/src/github.com/gatorloopwebapp/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, restful.DefaultContainer)
}

// handler to serve static files
func staticFromQueryParam(req *restful.Request, resp *restful.Response) {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join("./static", req.QueryParameter("resource")))
}
