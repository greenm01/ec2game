package core

// Ship Classes:
// 01 = Destroyer
// 02 = Cruiser
// 03 = Battleship
// 04 = Scout
// 05 = Troop Transport
// 06 = ETAC

type Ship struct {
	ID    int
	Class int
	AR    int
}

type Fleet struct {
	ID       int
	Ships    []Ship
	Pos      int
	Speed    int
	MaxSpeed int
	ROE      int
	ETA      int
	Orders   int
}

func (f *Fleet) Armies() int {
	a := 0
	for i, s := range f.Ships {
		// Only troop transports can have armies!
		if s.Class == 5 {
			a += s.AR
		} else {
			f.Ships[i].AR = 0
		}
	}
	return a
}
