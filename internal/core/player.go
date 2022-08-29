package core

type Player struct {
    UID int
    EmpireName string
    Planets []int
    Fleets []*Fleet
    Reports []*Report    
    Messages []*Message
}

func (p *Player) InitPlayer(name string, hw int) error {
    
    p.EmpireName = name
    p.Planets = append(p.Planets, hw)
    
    
    return nil    
}


 
