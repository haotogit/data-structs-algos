package search

import (
    "fmt"
    "testing"
    "../util"
)

func makeTestGrid(size int) *Grid {
    randSize := util.GetRandIntn(size)+20
    randNmines := randSize
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
                        if (xPos > 0 && yPos > 0) && ((xPos != x || yPos != y) &&
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
// if currCell.x < lastCell.x-1 or > lastCell.x+1 ||
// currCell.y < lastCell.y-1 or > lastCell.y+1 it is outside of current neighboring range
// 3. test that it visits the right neighbors
// 
func TestBFS(t *testing.T) {
    var currCell, papaCell *Cell
    var newGrid *Grid
    var visited []*Cell

    for iterations := 1; iterations < 11; iterations++ {
        newGrid = makeTestGrid(iterations+10)
        visited = make([]*Cell, newGrid.height*newGrid.width)
        randInt := util.GetRandIntn(newGrid.width*newGrid.height)
        randX := randInt/newGrid.width
        randY := randInt%newGrid.height
        resultQ := newGrid.BFS(randX, randY)
        if resultQ != nil && resultQ.Size() > 1 {
            fmt.Printf("=========resultsize %d\n", resultQ.Size())
            i := 0
            for ; resultQ.Size() != 0; {
                if papaCell == nil {
                    papaCell = resultQ.Dequeue().(*Cell)
                    visited[i] = papaCell
                    i++
                }

                fmt.Printf("newPapa %+v----------------\n", papaCell)
                currCell = resultQ.Dequeue().(*Cell)

                for (currCell.x >= papaCell.x - 1 && currCell.x <= papaCell.x + 1) &&
                    (currCell.y >= papaCell.y - 1 && currCell.y <= papaCell.y + 1) && 
                    resultQ.Size() != 0 {
                    visited[i] = currCell
                    i++
                    fmt.Printf("child %+v----------------\n", currCell)
                    currCell = resultQ.Dequeue().(*Cell)
                    fmt.Printf("nextchild %+v----------------\n", currCell)
                }

                papaCell = currCell
                visited[i] = currCell
                i++
                fmt.Printf("nextPapa %+v----------------\n", papaCell)

            }
        fmt.Printf("=========visited %d\n", len(visited))
        fmt.Printf("=========finalqqq %d\n", resultQ.Size())
        fmt.Printf("=========i%d\n", i)
        }
    }
}

// to test:
// 1. each cell only visited once
// 2. order of traversal
// if currCell.x < lastCell.x -1 or > lastCell.x ||
// currCell.y < lastCell
//func TestDFS(t *testing.T) {
//    var lastCell *Cell
//    var newGrid *Grid
//    var visited []*Cell
//    newGrid = makeTestGrid(10)
//    visited = make([]*Cell, newGrid.width*newGrid.height)
//    randInt := util.GetRandIntn(newGrid.width*newGrid.height)
//    randX := randInt/newGrid.width
//    randY := randInt%newGrid.height
//
//    resultQ := newGrid.DFS(randX, randY, nil)
//
//    if resultQ != nil {
//        for i := 0; resultQ.Size() != 0; i++ {
//            currCell := resultQ.Dequeue().(*Cell)
//            if i == 0 {
//                if currCell != newGrid.Rows[randX][randY] {
//                    t.Errorf("Incorrect order of traversal on initial node, should be %+v but got %+v", newGrid.Rows[randX][randY], currCell)
//                }
//            } else {
//                if (currCell.x < newGrid.width-1 || currCell.x > lastCell.x+1) ||
//                    (currCell.y < lastCell.y-1 || currCell.y > lastCell.x+1) {
//                        t.Errorf("Incorrect order of traversal, curr node at %+v is not direct descendant of %+v", currCell, lastCell)
//                } else {
//                    for k := 0; k < len(visited); k++ {
//                        if visited[k] == currCell {
//                            t.Errorf("Node already visited at position %d, %+v", k, currCell)
//                        }
//                    }
//                }
//            }
//        }
//    }
//}
