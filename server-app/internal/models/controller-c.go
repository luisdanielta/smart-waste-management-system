package models

type Location_t struct {
	City    string
	Country string
	Address string
	Lat     float64
	Lng     float64
}

type Status_t struct {
	Battery int
	Signal  int
	Online  bool
}

type ControllerC struct {
	Name     string
	Version  string
	Ip       string
	Status   Status_t
	Location Location_t
	Devices  []Container
}
