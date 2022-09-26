package core

type Empire struct {
	ID        	int
	Name      	string
	Bio			string
	Planets   	[]int
	PrevPlanets int
	Fleets    	map[int]*Fleet
	Reports   	[]Report
	Messages  	[]Message
	Tax       	int
	PDB       	PlanetDB
	Autopilot 	bool
	Status    	string
	CurProd   	int
	PrevProd  	int
}
