package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/solarsystem-go-restful/data"
)

var DB *sqlx.DB

//Planet - planet struct with properties
type Planet struct {
	ID             int     `json:id`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Density        float32 `json:"density"`
	Tilt           float32 `json:"tilt"`
	ImageUrl       string  `json:"imageUrl" db:"imageUrl"`
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

//Planets - collection of planet structs
type Planets struct {
	Planets []Planet
}

//Moon - moon struct
type Moon struct {
	ID        int    `json:id`
	Name      string `json:"name"`
	Planet_Id int32  `json:"planetid"`
}

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
	planets := []Planet{}
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

func main() {
	var err error
	DB, err = sqlx.Connect("sqlite3", "../solarsystem.db")
	if err != nil {
		log.Println("Problem creating the database")
	}
	defer DB.Close()
	data.Initialize(DB)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	p := Planet{}
	p.Register(wsContainer)
	ps := Planets{}
	ps.Register(wsContainer)

	log.Printf("start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
