package ships

type Ship struct {
	Name string
	Type int
	Class int
	Attack int
	Defense int
	Hull int
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
