package core

type StarDock struct {
	Ships []Ship // Stardock contents
	AR    int    // Armies
}

type Planet struct {
	ID      int      // Internal ID
	Pos     XY       // Starmap position
	Name    string   // Planet's name
	Empire  int   // Current owner
	MaxProd int      // Max production
	CurProd int      // Current production
	BTC     int      // Bitcoin stored
	AR      int      // Armies
	GB      int      // Ground batteries
	Dock    StarDock // Stardock
}

func (p *Planet) InitHomeworld(empire int) {
	p.Empire = empire
	p.MaxProd = 100
	p.CurProd = 100
	p.AR = 100
	p.GB = 25
}
