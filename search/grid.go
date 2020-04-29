package search

import (
    "fmt"
    "../util"
)

type Gridd interface {
    SetCells()
    //BFS()
    //DFS()
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
        randInt := util.GetRandIntn(g.width*g.height)
        randX := randInt/g.width
        randY := randInt%g.height
        g.Rows[randX][randY].PlantMine()
        nMines--
    }
}

func (g *Grid) GetNeighbors(x, y int) {
    currCell := g.Rows[x][y]
    neighbors := 0

    for xPos := x-1; xPos < g.width && xPos <= x+1; xPos++ {
        for yPos := y-1; yPos < g.height && yPos <= y+1; yPos++ {
            if ((xPos >= 0 && yPos >= 0) && (xPos != x && yPos != y) &&
                g.Rows[xPos][yPos].mined) {
                neighbors++
            }
        }
    }

    currCell.WelcomeNeighbors(neighbors)

    if (x == g.width-1 && y == g.height-1) {
        return
    } else if (x == g.width-1 && y < g.height-1) {
        g.GetNeighbors(0, y+1)
    } else {
        g.GetNeighbors(x+1, y)
    }
}

func BuildGrid(size int, nMines int) *Grid {
    newGrid := &Grid{
        size,
        size,
        make([][]*Cell, size),
    }

    newGrid.SetCells()
    newGrid.MineCells(nMines)
    newGrid.GetNeighbors(0, 0)
    return newGrid
}

func (g *Grid) Printer() {
    for x := 0; x < g.width; x++ {
        for y := 0; y < g.height; y++ {
            msg := "Cell %+v "
            if (y == g.height-1) {
                msg+= "\n"
            }

            fmt.Printf(msg, g.Rows[x][y])
        }
    }
}
