package planet

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
