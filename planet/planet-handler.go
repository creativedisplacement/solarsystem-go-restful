package planet

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// Register adds paths and routes to container
func (p *Planet) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/planet").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{planet-id}").To(p.getPlanet))

	container.Add(ws)
}

// GET http://localhost:8000/v1/planet/1
func (p Planet) getPlanet(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("planet-id")
	err := DB.Get(&p, "SELECT ID, NAME, DESCRIPTION, DENSITY, TILT, IMAGEURL, ROTATIONPERIOD, PERIOD, RADIUS, MOONS, AU, ECCENTRICITY, VELOCITY, MASS, INCLINATION, ORDINAL FROM planets WHERE id=?", id)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Planet could not be found.")
	} else {
		response.WriteEntity(p)
	}
}
