package search

import (
    "testing"
    "../util"
)

func TestGetNeighbors(t *testing.T) {
    for iterations := 0; iterations < 100; iterations++ {
        randSize := util.GetRandIntn(iterations)+20
        randNmines := randSize/3
        newGrid := BuildGrid(randSize, randNmines)

        for x := 0; x < randSize; x++ {
            for y := 0; y < randSize; y++ {
                currCell := newGrid.Rows[x][y]
                neighbors := 0

                for xPos := x-1; xPos < randSize && xPos <= x+1; xPos++ {
                    for yPos := y-1; yPos < randSize && yPos <= y+1; yPos++ {
                        if ((xPos >= 0 && yPos >= 0) && (xPos != x && yPos != y) &&
                            newGrid.Rows[xPos][yPos].mined) {
                            neighbors++
                        }
                    }
                }

                if (currCell.Neighbors != neighbors) {
                    t.Errorf("Wrong count of neighbors for cell at %d, %d, expected %d, but got %d", x, y, neighbors, currCell.Neighbors)
                }
            }
        }
    }
}
