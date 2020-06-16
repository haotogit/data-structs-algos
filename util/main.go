package util

import (
	"math/rand"
	"reflect"
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
	var nilVal bool
	switch el.(type) {
	case string:
		nilVal = "" == el.(string)
	case int:
		nilVal = 0 == el.(int)
	case uint32:
		nilVal = 0 == el.(uint32)
	}

	return nilVal
}

func Greater(el, el1 interface{}) int {
	compareVal := 0
	switch el.(type) {
	case string:
		elTyped := el.(string)
		el1Typed := el1.(string)
		if elTyped > el1Typed {
			compareVal = 1
		} else if elTyped < el1Typed {
			compareVal = -1
		}
	case int:
		elTyped := el.(int)
		el1Typed := el1.(int)
		if elTyped > el1Typed {
			compareVal = 1
		} else if elTyped < el1Typed {
			compareVal = -1
		}
	case uint32:
		elTyped := el.(uint32)
		el1Typed := el1.(uint32)
		if elTyped > el1Typed {
			compareVal = 1
		} else if elTyped < el1Typed {
			compareVal = -1
		}
	}

	return compareVal
}

func ObterItens(qty int, typeFlag string) []interface{} {
	var list []interface{}
	if typeFlag == "string" {
		for qty > 0 {
			currLen := GetRandIntn(1, 12)
			bytes := make([]byte, currLen)
			for i := 0; i < currLen; i++ {
				bytes[i] = byte(GetRandIntn(48, 122))
			}

			currStr := string(bytes)
			list = append(list, currStr)
			qty--
		}
	} else {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(qty)
		for i := 0; i < qty; i++ {
			list = append(list, tmp[i]*(i+1)*2)
		}
	}

	return list
}

func GetType(el interface{}) string {
	return reflect.TypeOf(el).String()
}
