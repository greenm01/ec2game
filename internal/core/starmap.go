package core

import (
    "errors"
    "math/rand"
    "time"
    "fmt"
)

type StarMap struct {
    Planets map[string]Planet
    Systems []int
    
}

// Return x,y coordinates of grid position
func GridPos(l int, grid int) (int, int) {
    return (l-1)%grid, (l-1)/grid
}

// Init: create starmap based on number of players
func (s *StarMap) Init(nPlayers int) error {
    var grid int
        
    // Map size based on number of players
    switch nPlayers {
    case 4:
        grid = 18
    case 9:
        grid = 27
    case 16:
        grid = 36
    case 25:
        grid = 45
    default:
        return errors.New("Invalid number of players; limited to 4,9,16,25")        
    }
    
    // Number of planets is five times the number players
    // TODO: Make nPlanets configurable
    nPlanets := nPlayers * 5   
    s.Systems = make([]int, nPlanets)
    
    // Random set of unque locations    
    rand.Seed(time.Now().UnixNano())
    p := rand.Perm(grid*grid)
    for i, r := range p[:nPlanets] {
        s.Systems[i] = r
        x,y := GridPos(r,grid)
        fmt.Println(r,":(",x,",",y,")")
    } 
              
    return nil   
        
}
