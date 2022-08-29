package core


type StarDock struct {
	Ships  []Ship				// Stardock contents
	AR int						// Armies
}

type Planet struct {
	ID			  int			// Internal ID
	Pos           XY			// Starmap position
	Name          string		// Planet's name
	Owner         string		// Current owner
	MaxProd 	  int			// Max production
	CurProd 	  int			// Current production
	BTC  		  int			// Bitcoin stored
	AR            int			// Armies
	GB            int			// Ground batteries
	Dock          StarDock		// Stardock
}
