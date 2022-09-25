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
	YearViewed  map[int]int
	MaxProd     map[int]int
	CurProd     map[int]int
	BTC         map[int]int
	Owner       map[int]string
	PrevOwner   map[int]string
	OwnedFor    map[int]int
	Dock        map[int][]int
	AR          map[int]int
	GB          map[int]int
	Pos         map[int]XY
}

func (p *PlanetDB) Init() {
	p.Name = make(map[int]string)
	p.Owner = make(map[int]string)
	p.PrevOwner = make(map[int]string)
	p.YearScouted = make(map[int]int)
	p.YearViewed = make(map[int]int)
	p.MaxProd = make(map[int]int)
	p.CurProd = make(map[int]int)
	p.BTC = make(map[int]int)
	p.OwnedFor = make(map[int]int)
	p.AR = make(map[int]int)
	p.GB = make(map[int]int)
	p.Pos = make(map[int]XY)
	p.Dock = make(map[int][]int)
}
