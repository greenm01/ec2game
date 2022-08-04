package ga

type Ship struct {
	Name string
	Type int
	Class int
	Attack int
	Defense int
	Hull int
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
	Armies []Army
}

type Fleet struct {
	ID int
	Ships []Ship
	Pos pos
	Speed int
	MaxSpeed int
	ROE int
	ETA int
	Orders  Mission
}
