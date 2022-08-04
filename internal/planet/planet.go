package planet

include (
	"github.com/greenm01/ec2/internal/ship"
)

type army struct {
	Attack int
	Defense int
	Health int
}

type groundBattery struct {
	Attack int
	Defense int
	Health int
}

type starDock struct {
	Ships []ship.Ship
	Armies []Army
	
}

type Planet struct {
	Location [2]int
	Name string
	Owner string
	MaxProduction int
	CurProduction int
	storedPoints int
	AR []army
	GB []groundBattery
	Dock StarDock
}
