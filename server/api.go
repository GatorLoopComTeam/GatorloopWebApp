package server

import (
	"html/template"
	"net/http"
	"path"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/gatorloopwebapp/api"
)

// RegisterAPI : registers routes with go-restful
func RegisterAPI(apiPath string, container *restful.Container) {

	ws := new(restful.WebService)

	ws.Path(apiPath).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/static").To(staticFromQueryParam).
		Doc("Serves static files").
		Param(ws.PathParameter("resource", "the path to a resource within /static")))

	ws.Route(ws.GET("/home").To(home).
		Doc("Serves index.html"))

	var v api.Velocity
	ws.Route(ws.GET("/velocity").
		To(v.GetRecent).
		Doc("Get an average of the last 5 velocities."))

	ws.Route(ws.POST("/velocity").
		To(v.CreateVelocity).
		Doc("Create a new velocity entry.").
		Param(ws.BodyParameter("velocity", "a velocity value").DataType("float")))

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

// handler to serve index.html on /
func home(req *restful.Request, resp *restful.Response) {
	// you might want to cache compiled templates
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Errorf("Template gave: %s", err)
	}
	t.Execute(resp.ResponseWriter, t)
}
