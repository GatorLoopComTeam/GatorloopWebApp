package server

import (
	"net/http"
	"path"

	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/gatorloopwebapp/api"
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

	ws.Route(ws.GET("/node_modules").To(staticNMFromQueryParam).
		Doc("Serves static node_modules files").
		Param(ws.PathParameter("resource", "the path to a resource within /node_modules")))

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
		Doc("Get the most recent velocity."))

	var a api.Acceleration
	ws.Route(ws.GET("/acceleration").
		To(a.GetRecent).
		Doc("Get the most recent acceleration."))

	var p api.Position
	ws.Route(ws.GET("/position").
		To(p.GetRecent).
		Doc("Get the most recent position."))

	var t api.Temperature
	ws.Route(ws.GET("/temperature").
		To(t.GetRecent).
		Doc("Get the most recent temperature."))

	var pBat api.PrimaryBattery
	ws.Route(ws.GET("/primarybattery").
		To(pBat.GetRecent).
		Doc("Get the most recent primary battery values"))

	var aBat api.AuxiliaryBattery
	ws.Route(ws.GET("/auxbattery").
		To(aBat.GetRecent).
		Doc("Get the most recent auxiliary battery values"))

	var s api.Stop
	ws.Route(ws.GET("/stop").
		To(s.EmergencyStop).
		Doc("Initiates an emergency stop of the pod."))

	var k api.KillPower
	ws.Route(ws.GET("/killpower").
		To(k.SendKillPower).
		Doc("Throws the circuit breaker on the pod to kill power."))

	var state api.State
	ws.Route(ws.GET("/state").
		To(state.GetRecent).
		Doc("Gets the current state of the pod."))

	restful.Add(ws)

	config := swagger.Config{
		WebServices:    restful.DefaultContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/home/pi/go/src/github.com/gatorloopwebapp/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, restful.DefaultContainer)
}

// handler to serve static files
func staticFromQueryParam(req *restful.Request, resp *restful.Response) {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join("./static", req.QueryParameter("resource")))
}

// handler to serve static node_modules files
func staticNMFromQueryParam(req *restful.Request, resp *restful.Response) {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join("./node_modules", req.QueryParameter("resource")))
}
