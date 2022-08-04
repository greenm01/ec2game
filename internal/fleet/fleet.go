package fleet

include (
	"github.com/grenem01/ec2/internal/planet"
)

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
	Armies []planet.Army
}

typw Mission struct {
	ID int
	Destination [2]int
}

type Fleet struct {
	ID int
	Ships []Ship
	Location [2]int
	Speed int
	MaxSpeed int
	ROE int
	ETA int
	Orders  Mission
}
