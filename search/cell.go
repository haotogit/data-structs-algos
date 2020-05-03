package search

type Cell struct {
    x int
    y int
    mined bool
    neighbors int
    visited bool
}

func NewCell(x int, y int) *Cell {
    return &Cell{ x, y, false, 0, false }
}

func (c *Cell) PlantMine() {
    c.mined = true
}

func (c *Cell) SetNeighbors(n int) {
    c.neighbors = n
}

func (c *Cell) GetNeighbors() int {
    return c.neighbors
}
