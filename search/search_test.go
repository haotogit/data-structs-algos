package search

import (
    "testing"
    "github.com/haotogit/data-structs-algos/util"
    "github.com/haotogit/data-structs-algos/queue"
)

func makeTestGrid(size int) *Grid {
    randSize := util.GetRandIntn(0, size)+20
    randNmines := randSize
    return NewGrid(randSize, randNmines)
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
                        if (xPos > 0 && yPos > 0) && ((xPos != x || yPos != y) &&
                            newGrid.Rows[xPos][yPos].mined) {
                            neighbors++
                        }
                    }
                }

                if (currCell.minedNeighbors != neighbors) {
                    t.Errorf("Wrong count of neighbors for cell at %d, %d, expected %d, but got %d", x, y, neighbors, currCell.minedNeighbors)
                }
            }
        }
    }
}

// test:
// 1. each cell is only visited only once
// 2. order of traversal and that it visits the right neighbors
// traverse resultQ by getting subtree top node's unmined neighbors,
// then each of those neighbors unmined neighbors subsequently. 
// until descending neighbors have no unmined
func TestBFS(t *testing.T) {
    var newGrid *Grid
    for iterations := 0; iterations < 100; iterations++ {
        var currCell, papaCell *Cell
        newGrid = makeTestGrid(iterations+20)
        randInt := util.GetRandIntn(0, newGrid.width*newGrid.height)
        randX := randInt/newGrid.width
        randY := randInt%newGrid.height
        resultQ := newGrid.BFS(randX, randY)

        if resultQ != nil && resultQ.Size() > 1 {
            resultSize := resultQ.Size()
            checkVisited(t, resultQ)
            count := 0
            for resultQ.Size() != 0 {
                if papaCell == nil {
                    papaCell = resultQ.Dequeue().(*Cell)
                    count++
                } else {
                    papaCell = currCell
                }

                currCell = resultQ.Dequeue().(*Cell)
                count++
                for ((currCell.x >= papaCell.x - 1 && currCell.x <= papaCell.x + 1) &&
                    (currCell.y >= papaCell.y - 1 && currCell.y <= papaCell.y +1)) {
                    if resultQ.Size() == 0 {
                        break
                    } else {
                        currCell = resultQ.Dequeue().(*Cell)
                        count++
                    }
                }
            }

            if resultSize != count || resultQ.Size() != 0 {
                t.Errorf("Traversed incorrect number of nodes resultSize %d, but got %d, and should be 0 %d", resultSize, count, resultQ.Size())
            }
        }
    }
}

func checkVisited(t *testing.T, resultQ *queue.Queue) {
    resultSize := resultQ.Size()
    visited := make([]*Cell, 0, resultSize)
    for x := 0; x < resultSize; x++ {
        currCell := resultQ.Dequeue().(*Cell)
        for i := 0; i < len(visited); i++ {
            if currCell == visited[i] {
                t.Errorf("Visited incorrectly, found dupe")
            }
        }

        visited = visited[:x+1]
        visited[x] = currCell
        resultQ.Enqueue(currCell)
    }
}
