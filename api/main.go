package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"github.com/solarsystem-go-restful/data"
)

var DB *sql.DB

type Planet struct {
	ID             int     `json:id`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Density        float32 `json:"density"`
	Tilt           float32 `json:"tilt"`
	ImageUrl       string  `json:"imageUrl"`
	RotationPeriod float32 `json:"rotationperiod"`
	Period         float32 `json:"period"`
	Radius         int64   `json:"radius"`
	Moons          int8    `json:"moons"`
	AU             float32 `json:"au"`
	Eccentricity   float32 `json:"eccentricity"`
	Velocity       float32 `json:"velocity"`
	Mass           float32 `json:"mass"`
	Inclination    float32 `json:"inclination"`
	Ordinal        int8    `json:"order"`
}

type Moon struct {
	ID        int    `json:id`
	Name      string `json:"name"`
	Planet_Id int32  `json:"planetid"`
}

// Register adds paths and routes to container
func (p *Planet) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/planets").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	ws.Route(ws.GET("/").To(p.getPlanets))
	ws.Route(ws.GET("/{planet-id}").To(p.getPlanet))

	container.Add(ws)
}

// GET http://localhost:8000/v1/planets/
func (p Planet) getPlanets(request *restful.Request, response *restful.Response) {
	err := DB.QueryRow("select ID, NAME, DESCRIPTION, DENSITY, TILT, IMAGEURL, ROTATIONPERIOD, PERIOD, RADIUS, MOONS, AU, ECCENTRICITY, VELOCITY, MASS, INCLINATION, ORDINAL FROM planets").Scan(&p.ID, &p.Name, &p.Description, &p.Density, &p.Tilt, &p.ImageUrl, &p.RotationPeriod, &p.Period, &p.Radius, &p.Moons, &p.AU, &p.Eccentricity, &p.Velocity, &p.Mass, &p.Inclination, &p.Ordinal)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Planet could not be found.")
	} else {
		response.WriteEntity(p)
	}
}

// GET http://localhost:8000/v1/planets/1
func (p Planet) getPlanet(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("planet-id")
	err := DB.QueryRow("SELECT ID, NAME, DESCRIPTION, DENSITY, TILT, IMAGEURL, ROTATIONPERIOD, PERIOD, RADIUS, MOONS, AU, ECCENTRICITY, VELOCITY, MASS, INCLINATION, ORDINAL FROM planets WHERE id=?", id).Scan(&p.ID, &p.Name, &p.Description, &p.Density, &p.Tilt, &p.ImageUrl, &p.RotationPeriod, &p.Period, &p.Radius, &p.Moons, &p.AU, &p.Eccentricity, &p.Velocity, &p.Mass, &p.Inclination, &p.Ordinal)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Planet could not be found.")
	} else {
		response.WriteEntity(p)
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "../solarsystem.db")
	if err != nil {
		log.Println("Problem creating the database")
	}
	defer DB.Close()
	data.Initialize(DB)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := Planet{}
	t.Register(wsContainer)

	log.Printf("start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
