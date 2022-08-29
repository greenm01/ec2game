package main

import (
	"fmt"
	"math"
	"github.com/mash/gokmeans"
)

/* See Numbers spreadsheet for classic EC Starmap example. The below listed "nodes" in the
   obersrvations set represent X,Y coordinates for existing planets. This Algo groups them 
   into 'n' different nodes based on the K-Means clustering algo (a.k.a. Lloy's algorithm). The foo set is
   filled with 'n' nodes at the exact center of our sarmap to intialize the algo. There are four players on
   this map, so we want four clusters.  The returned set finds our homeworlds in the optimal location, 
   relative to other players and the distributin of planets. 
*/

var observations []gokmeans.Node = []gokmeans.Node {
	gokmeans.Node{1, 14},
	gokmeans.Node{3, 13},
	gokmeans.Node{3, 15},
	gokmeans.Node{5, 4},
	gokmeans.Node{6, 5},
	gokmeans.Node{6, 16},
	gokmeans.Node{9, 1},
	gokmeans.Node{10, 1},
	gokmeans.Node{11, 17},
	gokmeans.Node{12, 4},
	gokmeans.Node{12, 8},
	gokmeans.Node{12, 12},
	gokmeans.Node{13, 8},
	gokmeans.Node{14, 6},
	gokmeans.Node{14, 16},
	gokmeans.Node{15, 16},
	gokmeans.Node{16, 10},
	gokmeans.Node{17, 17},
	gokmeans.Node{18, 7},
}

var foo []gokmeans.Node = []gokmeans.Node {
	gokmeans.Node{9, 9},
	gokmeans.Node{9, 9},
	gokmeans.Node{9, 9},
	gokmeans.Node{9, 9},
}

func main() {
	// Get a list of centroids and output the values
	if success, centroids := gokmeans.Train2(observations, -1, 50, foo); success {
		// Show the centroids
		fmt.Println("The centroids are")
		for _, cent := range centroids {
			cent[0] = math.Round(cent[0])
			cent[1] = math.Round(cent[1])
			fmt.Println(cent)
		}

		// Output the clusters
		fmt.Println("...")
		for _, observation := range observations {
			index := gokmeans.Nearest(observation, centroids)
			fmt.Println(observation, "belongs in cluster", index+1, ".")
		}
		
		fmt.Println("finding homeworlds...")
		for _, cent := range centroids {
			i := gokmeans.Nearest(cent, observations)
			fmt.Println(cent, "nearest to", observations[i])
		}
		
	}
}