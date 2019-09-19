package moon

//Moon - moon struct
type Moon struct {
	ID        int    `json:id`
	Name      string `json:"name"`
	Planet_Id int32  `json:"planetid"`
}
