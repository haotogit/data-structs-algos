package search

import (
    "testing"
    "../util"
    "../queue"
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
    var subTree *queue.Queue
    var visited *queue.Queue

    for iterations := 0; iterations < 100; iterations++ {
        var currCell, papaCell *Cell
        newGrid = makeTestGrid(iterations+20)
        visited = queue.NewQ(newGrid.width*newGrid.height)
        randInt := util.GetRandIntn(newGrid.width*newGrid.height)
        randX := randInt/newGrid.width
        randY := randInt%newGrid.height
        resultQ := newGrid.BFS(randX, randY)

        if resultQ != nil && resultQ.Size() > 1 {
            for resultQ.Size() != 0 {
                subTree = queue.NewQ(newGrid.height*newGrid.width)
                if papaCell == nil {
                    papaCell = resultQ.Dequeue().(*Cell)
                }

                checkVisited(t, papaCell, visited)
                visited.Enqueue(papaCell)
                if papaCell.minedNeighbors == 0 {
                    currCell = resultQ.Dequeue().(*Cell)

                    // get papaCell direct children and put them in subtree
                    // last currCell will be first child of first node in subtree
                    for resultQ.Size() != 0 &&
                        ((currCell.x >= papaCell.x - 1 && currCell.x <= papaCell.x + 1) &&
                        (currCell.y >= papaCell.y - 1 && currCell.y <= papaCell.y +1)) {
                        subTree.Enqueue(currCell)
                        checkVisited(t, currCell, visited)
                        visited.Enqueue(currCell)
                        currCell = resultQ.Dequeue().(*Cell)
                    }

                    // get children of subtree and subsequent children
                    for resultQ.Size() != 0 && subTree.Size() != 0 {
                        mamaCell := subTree.Dequeue().(*Cell)
                        if mamaCell.minedNeighbors == 0 {
                            for (currCell.x >= mamaCell.x - 1 &&
                                currCell.x <= mamaCell.x + 1) &&
                                (currCell.y >= mamaCell.y -1 && currCell.y <= mamaCell.y + 1) {
                                if currCell.minedNeighbors == 0 {
                                    subTree.Enqueue(currCell)
                                }

                                checkVisited(t, currCell, visited)
                                visited.Enqueue(currCell)
                                currCell = resultQ.Dequeue().(*Cell)

                                // if resultQ.Size == 0 check last child
                                // hasn't been visited
                                if resultQ.Size() == 0 {
                                    checkVisited(t, currCell, visited)
                                    visited.Enqueue(currCell)
                                    break
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}

func checkVisited(t *testing.T, currCell *Cell, visited *queue.Queue) {
    for visited.Size() != 0 {
        currVisited := visited.Dequeue().(*Cell)
        if currCell == currVisited {
            t.Errorf("Node already visited %+v", currCell)
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
