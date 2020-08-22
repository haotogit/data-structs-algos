package search

import (
    "github.com/haotogit/data-structs-algos/util"
    "github.com/haotogit/data-structs-algos/queue"
)

type Gridd interface {
    SetCells()
    MineCells()
    SetNeighbors()
    BFS(x, y int) *queue.Queue
    DFS(x, y int) *queue.Queue
    OuttaRange(x, y int) bool
}

type Grid struct {
    width int
    height int
    Rows [][]*Cell
}

func (g *Grid) SetCells() {
    for x := 0; x < g.width; x++ {
        g.Rows[x] = make([]*Cell, g.height)
        for y := 0; y < g.height; y++ {
            g.Rows[x][y] = NewCell(x, y) 
        }
    }
}

func (g *Grid) MineCells(nMines int) {
    for nMines > 0 {
        randInt := util.GetRandIntn(0, g.width*g.height)
        randX := randInt/g.width
        randY := randInt%g.height
        g.Rows[randX][randY].PlantMine()
        nMines--
    }
}

func (g *Grid) SetNeighbors(x, y int) {
    currCell := g.Rows[x][y]
    neighbors := 0
    for xPos := x-1; xPos < g.width && xPos <= x+1; xPos++ {
        for yPos := y-1; yPos < g.height && yPos <= y+1; yPos++ {
            if (xPos > 0 && yPos > 0) && ((xPos != x || yPos != y) &&
                g.Rows[xPos][yPos].mined) {
                neighbors++
            }
        }
    }

    currCell.SetNeighbors(neighbors)

    if (x == g.width-1 && y == g.height-1) {
        return
    } else if (x == g.width-1 && y < g.height-1) {
        g.SetNeighbors(0, y+1)
    } else {
        g.SetNeighbors(x+1, y)
    }
}

func (g *Grid) OuttaRange(x, y int) bool {
    return (x < 0 || x > g.width-1) ||
        (y < 0 || y > g.height-1)
}

func (g *Grid) BFS(x, y int) *queue.Queue {
    if g.OuttaRange(x, y) || g.Rows[x][y].mined {
        return nil
    } 

    resultQ := queue.NewQ(g.width*g.height)
    searchQ := queue.NewQ(g.width*g.height)
    currCell := g.Rows[x][y]
    currCell.visited = true
    searchQ.Enqueue(currCell)

    for searchQ.Size() != 0 {
        currCell = searchQ.Dequeue().(*Cell)
        resultQ.Enqueue(currCell)
        if currCell.GetNeighbors() == 0 {
            for xPos := currCell.x-1; xPos < g.width && xPos <= currCell.x+1; xPos++ {
                for yPos := currCell.y-1; yPos < g.height && yPos <= currCell.y+1; yPos++ {
                    if (xPos > 0 && yPos > 0) && (xPos != currCell.x || yPos != currCell.y) {
                        nextCell := g.Rows[xPos][yPos]
                        if !nextCell.visited {
                            nextCell.visited = true
                            searchQ.Enqueue(nextCell)
                        }
                    }
                }
            }
        }
    }

    return resultQ
}

func (g *Grid) DFS(x, y int, resultQ *queue.Queue) *queue.Queue {
    if g.OuttaRange(x, y) || g.Rows[x][y].mined {
        return nil
    }

    currCell := g.Rows[x][y]
    if resultQ == nil {
        resultQ = queue.NewQ(g.width*g.height)
        currCell.visited = true
    }

    resultQ.Enqueue(currCell)
    if currCell.GetNeighbors() == 0 {
        for xPos := x-1; xPos < g.width && xPos <= x+1; xPos++ {
            for yPos := y-1; yPos < g.height && yPos <= y+1; yPos++ {
                if (xPos >= 0 && yPos >= 0) && (xPos != x && yPos != y) {
                    nextCell := g.Rows[xPos][yPos]
                    if !nextCell.visited {
                        nextCell.visited = true
                        g.DFS(xPos, yPos, resultQ)
                    }
                }
            }
        }
    }

    return resultQ
}

func NewGrid(size, nMines int) *Grid {
    newGrid := &Grid{
        size,
        size,
        make([][]*Cell, size),
    }

    newGrid.SetCells()
    newGrid.MineCells(nMines)
    newGrid.SetNeighbors(0, 0)
    return newGrid
}
