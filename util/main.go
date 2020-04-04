package util

import (
    "math/rand"
    "time"
)

func GetRandIntn(max int) int {
    if max == 0 {
        return max
    }

    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max)
}
