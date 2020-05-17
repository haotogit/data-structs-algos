package search

type Cell struct {
    x int
    y int
    mined bool
    minedNeighbors int
    visited bool
}

func (c *Cell) PlantMine() {
    c.mined = true
}

func (c *Cell) SetNeighbors(n int) {
    c.minedNeighbors = n
}

func (c *Cell) GetNeighbors() int {
    return c.minedNeighbors
}

func NewCell(x int, y int) *Cell {
    return &Cell{ x, y, false, 0, false }
}

