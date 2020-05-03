package search

import (
    "testing"
    "../util"
)

func makeTestGrid(size int) *Grid {
    randSize := util.GetRandIntn(size)+20
    randNmines := randSize/3
    return BuildGrid(randSize, randNmines)
}

func TestGetNeighbors(t *testing.T) {
    for iterations := 0; iterations < 100; iterations++ {
        newGrid := makeTestGrid(iterations)
        for x := 0; x < newGrid.width; x++ {
            for y := 0; y < newGrid.height; y++ {
                currCell := newGrid.Rows[x][y]
                neighbors := 0

                for xPos := x-1; xPos < newGrid.width && xPos <= x+1; xPos++ {
                    for yPos := y-1; yPos < newGrid.height && yPos <= y+1; yPos++ {
                        if ((xPos >= 0 && yPos >= 0) && (xPos != x && yPos != y) &&
                            newGrid.Rows[xPos][yPos].mined) {
                            neighbors++
                        }
                    }
                }

                if (currCell.neighbors != neighbors) {
                    t.Errorf("Wrong count of neighbors for cell at %d, %d, expected %d, but got %d", x, y, neighbors, currCell.neighbors)
                }
            }
        }
    }
}

// test:
// 1. each cell is only visited only once
// check that resultQ doesn't have dupes
// 2. order of traversal
// if currCell.x < lastCell.x-3 or > lastCell.x+3 ||
// currCell.y < lastCell.y-3 or > lastCell.y+3 it is outside of current neighboring range
func TestBFS(t *testing.T) {
    var lastCell *Cell
    var newGrid *Grid
    for iterations := 0; iterations < 100; iterations++ {
        newGrid = makeTestGrid(iterations+10)
        randInt := util.GetRandIntn(newGrid.width*newGrid.height)
        randX := randInt/newGrid.width
        randY := randInt%newGrid.height
        visited := make([]*Cell, newGrid.width*newGrid.height)
        resultQ := newGrid.BFS(randX, randY)

        if resultQ != nil {
            for i := 0; resultQ.Size() != 0; i++ {
                currCell := resultQ.Dequeue().(*Cell)
                if i == 0 {
                    if currCell != newGrid.Rows[randX][randY] {
                        t.Errorf("Incorrect order of traversal on initial node, should be %+v but got %+v", newGrid.Rows[randX][randY], currCell)
                    }
                } else {
                    // if currcell is not neighboring current x, y
                    // and its not neighboring the lastest set of visited
                    if ((currCell.x < lastCell.x-3 || currCell.x > lastCell.x+3) ||
                        (currCell.y < lastCell.y-3 || currCell.y > lastCell.y+3)) {
                            t.Errorf("Incorrect order of traversal, curr node is not neighboring current x, y and not neighboring last visited set of neighbors")
                    } else {
                        for k := 0; k < len(visited); k++ {
                            if visited[k] == currCell {
                                t.Errorf("Node already visited at position %d, %+v", k, currCell)
                            }
                        }
                    }
                }

                visited[i] = currCell
                lastCell = currCell
            }
        }
    }
}
