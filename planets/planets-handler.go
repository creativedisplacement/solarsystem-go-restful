package planets

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jmoiron/sqlx"
	"github.com/solarsystem-go-restful/planet"
)

var DB *sqlx.DB

func (p *Planets) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/planets").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(p.getPlanets))

	container.Add(ws)
}

// GET http://localhost:8000/v1/planets/
func (p Planets) getPlanets(request *restful.Request, response *restful.Response) {
	planets := []planet.Planet{}
	err := DB.Select(&planets, "select ID, NAME, DESCRIPTION, DENSITY, TILT, IMAGEURL, ROTATIONPERIOD, PERIOD, RADIUS, MOONS, AU, ECCENTRICITY, VELOCITY, MASS, INCLINATION, ORDINAL FROM planets")
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Planet could not be found.")
	} else {
		p.Planets = planets
		response.WriteEntity(p)
	}
}
