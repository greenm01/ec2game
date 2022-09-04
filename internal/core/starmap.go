package core

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

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

	fmt.Print("Finding homeworlds...")

	// Get a list of centroid clusters
	_, centroids := km.Train(sys, np, 50)

	hw := make([][]int, np)
	for i, cent := range centroids {
		hw[i] = []int{int(cent[0]), int(cent[1])}
	}

	fmt.Println("done!\n...")

	return hw

}

func (s StarMap) NumPlanets() int {
	return len(s.Systems)
}

// Init: Setup the game's starmap and planets
func (s *StarMap) InitMap(np int) error {

	fmt.Print("...\nGenerating starmap...")

	// Classic EC starmap size bound by number of players:
	// [4:18x18, 9:27x27, 16:36x36, 25:45x45]
	// Nonlinar power regression gives us a nice distribution
	// f(x) = 9*x^(-0.5)
	// https://www.desmos.com/calculator/9itjhjdmig
	r := 9 * math.Pow(float64(np), 0.5)
	s.GridSize = int(math.Ceil(r))

	// Generate starmap based on a random Poisson distribution
	// This generates a much nicer distribution over a pure random set
	// http://devmag.org.za/2009/05/03/poisson-disk-sampling/
	// Minimum distance between systems is PI. Why not?

	/* TODO: Consider adding perlin noise to distribution */

	gs := float64(s.GridSize)
	points := pd.Sample(0, 0, gs, gs, math.Pi, 50, nil)
	nodes := make([]km.Node, len(points))
	s.Systems = make([]int, len(nodes))

	for i, p := range points {
		nodes[i] = []float64{p.X, p.Y}
		c := CellPos(s.GridSize, int(p.X), int(p.Y))
		s.Systems[i] = c
	}

	fmt.Println("done!")

	for count := 1; count <= 10; count++ {

		// Employ k-means clustering to find homewords
		hw := kMeans(nodes, np, gs)

		s.HomeWorlds = make([]int, np)
		for i, h := range hw {
			s.HomeWorlds[i] = CellPos(s.GridSize, h[0], h[1])
		}
		// Make sure we have a unique set
		if countDuplicates(s.HomeWorlds) == 0 {
			break
		}

		if count > 9 {
			return errors.New("Error resolving homeworlds! Try fewer players.")
		}

	}

	s.Systems = append(s.HomeWorlds, s.Systems...)
	s.Systems = removeDuplicates(s.Systems)

	ppp := fmt.Sprintf("%.2f", float64(s.NumPlanets())/float64(np))
	fmt.Println("Grid size          =", gs, "x", gs)
	fmt.Println("Number of players  =", np)
	fmt.Println("Number of planets  =", s.NumPlanets())
	fmt.Println("Planets per player =", ppp)
	//fmt.Println("Systems =",s.Systems)
	//fmt.Println("Homeworlds =", s.HomeWorlds)

	// Seed the random number generator with system time
	rnd := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	s.Planets = make(map[int]*Planet)
	for _, c := range s.Systems {
		x, y := GridPos(c, s.GridSize)
		s.Planets[c] = &Planet{
			ID:      c,
			Pos:     XY{x, y},
			MaxProd: rnd.Intn(maxp-minp) + minp,
			Name:    "nameless",
		}
	}

	return nil

}

// https://www.dotnetperls.com/duplicates-go
func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func countDuplicates(dupArr []int) int {
	dupsize := len(dupArr)
	dupcount := 0
	for i := 0; i < dupsize; i++ {
		for j := i + 1; j < dupsize; j++ {
			if dupArr[i] == dupArr[j] {
				dupcount++
				break
			}
		}
	}
	return dupcount
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
