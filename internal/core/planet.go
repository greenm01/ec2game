package core

type Planet struct {
	ID        int    // Internal ID
	Pos       XY     // Starmap position
	Name      string // Planet's name
	Owner     int    // Current owner
	PrevOwner int    // Previous owner
	MaxProd   int    // Max production
	CurProd   int    // Current production
	BTC       int    // Bitcoin stored
	AR        int    // Armies
	GB        int    // Ground batteries
	Dock      []int  // Stardock contents
}

/* TODO: Make these global constants */
func (p *Planet) InitHomeworld(empire int) {
	p.Owner = empire
	p.PrevOwner = -1
	p.MaxProd = 100
	p.CurProd = 100
	p.BTC = 50
	p.AR = 100
	p.GB = 25
}

// PlanetDB holds an empire's planet database
type PlanetDB struct {
	// Map keys = Planet.ID
	Name        map[int]string
	YearScouted map[int]int
	MaxProd     map[int]int
	CurProd     map[int]int
	BTC         map[int]int
	Owner       map[int]int
	PrevOwner   map[int]int
	OwnedFor    map[int]int
	Dock        map[int][]int
	AR          map[int]int
	GB          map[int]int
	Pos         map[int]XY
}
