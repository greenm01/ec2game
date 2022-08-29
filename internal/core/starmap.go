package core

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"errors"

	pd "github.com/fogleman/poissondisc"
	km "github.com/mash/gokmeans"
)

const (
	maxp = 150 // Max planet production
	minp = 10  // Minimum planet production
)

type StarMap struct {
	Planets    map[int]*Planet
	GridSize   int
	Systems    []int
	HomeWorlds []int
}

// Return x,y grid coordinates of cell number
// Cell 1 starts at (0,0)
func GridPos(c int, grid int) (int, int) {
	return (c - 1) % grid, (c - 1) / grid
}

// Return the cell number of given x,y coordinates
// (0,0) starts at cell 1
func CellPos(grid int, x int, y int) int {
	return grid*y + x + 1
}

// kMeans: k-means clustering (LLoyd's algo) to find optimal
// distribution of homeworlds.
func kMeans(sys []km.Node, np int, gs float64) [][]int {

	// Generate our centroid seeds
	cg := gs / 2
	seed := make([]km.Node, np)
	for i, _ := range seed {
		seed[i] = []float64{cg, cg}
	}

	// Get a list of centroid clusters
	_, centroids := km.Train2(sys, -1, 50, seed)

	fmt.Println("...\nFinding homeworlds:")

	hw := make([][]int, np)
	for i, cent := range centroids {
		n := km.Nearest(cent, sys)
		hw[i] = []int{int(sys[n][0]), int(sys[n][1])}
	}

	fmt.Println(hw)

	return hw

}

func (s StarMap) NumPlanets() int {
	return len(s.Systems)
}

// Init: Setup the game's starmap and planets
func (s *StarMap) InitMap(np int) error {

	// Starmap size bound by number of players
	// Linear regression of original EC map sizing is approx:
	// [4:18x18, 9:27x27, 16:36x36, 25:45x45]
	// y = 1.27x + 14.42
	// We want some extra room and a more planets
	r := 1.5*float64(np) + 15
	// Round to nearest even number.
	s.GridSize = int(math.RoundToEven(math.Floor(r) + 0.5))
	s.Planets = make(map[int]*Planet)

	// Generate starmap based on a random Poisson distribution
	// This generates a much nicer distribution over a pure random set
	// http://devmag.org.za/2009/05/03/poisson-disk-sampling/
	// Minimum distance between systems is PI. Why not?

	gs := float64(s.GridSize)
	points := pd.Sample(0, 0, gs, gs, math.Pi, 32, nil)
	nodes := make([]km.Node, len(points)-1)
	s.Systems = make([]int, len(nodes))

	// Seed the random number generator with system time
	rnd := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	for i, p := range points[1:] {
		x := math.RoundToEven(p.X)
		y := math.RoundToEven(p.Y)
		nodes[i] = []float64{x,y}
		c := CellPos(s.GridSize, int(x), int(y))
		s.Systems[i] = c
		s.Planets[c] = &Planet{
			ID:      c,
			Pos:     XY{int(x), int(y)},
			MaxProd: rnd.Intn(maxp-minp) + minp,
			Name:    "nameless",
		}
	}

	ppp := fmt.Sprintf("%.2f", float64(s.NumPlanets())/float64(np))

	fmt.Println("...\nGenerating starmap:")
	fmt.Println("Grid size          =", gs, "x", gs)
	fmt.Println("Number of players  =", np)
	fmt.Println("Number of planets  =", s.NumPlanets())
	fmt.Println("Planets per player =", ppp)

	// Employ k-means clustering to find homewords
	hw := kMeans(nodes, np, gs)

	s.HomeWorlds = make([]int, np)
	for i, h := range hw {
		s.HomeWorlds[i] = CellPos(s.GridSize, h[0], h[1])
	}
	
	// Validate homeworlds just in case
	// Don't expect any rounding errors... :-O
	if !subset(s.HomeWorlds, s.Systems) {
		return errors.New("Homeworld validation error!")
	}	
	
	return nil
	
}

// subset returns true if the first array is completely
// contained in the second array.
func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}