package core

type Empire struct {
	UID       int
	Name      string
	Planets   []int
	Fleets    map[int]*Fleet
	Reports   []Report
	Messages  []Message
	Tax       int
	PDB       PlanetDB
	Autopilot bool
}
