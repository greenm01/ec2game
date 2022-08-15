package core

type Army struct {
	Attack  int
	Defense int
	Health  int
}

type groundBattery struct {
	Attack  int
	Defense int
	Health  int
}

type starDock struct {
	Ships  []Ship
	Armies []Army
}

type Planet struct {
	Pos           pos
	Name          string
	Owner         string
	MaxProduction int
	CurProduction int
	storedPoints  int
	AR            []Army
	GB            []groundBattery
	Dock          starDock
}
