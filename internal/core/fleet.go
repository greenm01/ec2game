package core

type Ship struct {
	Name     string
	Type     int
	Class    int
	Attack   int
	Defense  int
	Hull     int
	MaxSpeed int
}

type Destroyer struct {
	Ship
}

type Scout struct {
	Ship
}

type Cruiser struct {
	Ship
}

type Battleship struct {
	Ship
}

type ETAC struct {
	Ship
}

type TroopTransport struct {
	Ship
	AR int 						// Armies
}

type Fleet struct {
	ID       int
	Ships    []Ship
	Pos      XY
	Speed    int
	MaxSpeed int
	ROE      int
	ETA      int
	Orders   Mission
}
