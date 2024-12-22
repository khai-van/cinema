package cinemaservice

type Cinema struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Rows    int    `json:"rows"`
	Columns int    `json:"columns"`
}

type Seat struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}
