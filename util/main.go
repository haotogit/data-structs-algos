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

func IsNil(el interface{}) bool {
    var nilVal interface{}
    switch el.(type) {
        case string:
            nilVal = ""
        case int:
            nilVal = 0
        default:
            nilVal = nil
    }

    return nilVal == el
}
