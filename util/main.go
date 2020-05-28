package util

import (
    "math/rand"
    "time"
)

func GetRandIntn(min, max int) int {
    if max == 0 {
        return max
    }

    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
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

func ElComparer(el, el1 interface{}) int {
    compareVal := 0
    switch el.(type) {
        case string:
            elTyped := el.(string)
            el1Typed := el1.(string)
            if elTyped > el1Typed {
                compareVal = 1
            } else if elTyped < el1Typed {
                compareVal = - 1
            }
        case int:
            elTyped := el.(int)
            el1Typed := el1.(int)
            if elTyped > el1Typed {
                compareVal = 1
            } else if elTyped < el1Typed {
                compareVal = - 1
            }
            //elTyped = el.(int)

        //default:
        //    elTyped := el.(*Objecto)
        //    fmt.Printf("%+v", elTyped)
    }

    return compareVal
}

