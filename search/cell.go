package search

type Cell struct {
    row int
    col int
    mined bool
    Neighbors int
}

func NewCell(x int, y int) *Cell {
    return &Cell{ x, y, false, 0 }
}

func (c *Cell) PlantMine() {
    c.mined = true
}

func (c *Cell) WelcomeNeighbors(n int) {
    c.Neighbors = n
}
