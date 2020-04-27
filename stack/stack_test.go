package stack

import (
    "testing"
    "fmt"
)

//func TestPush(t *testing.T) {
//    for iterations := 0; iterations < 100; iterations++ {
//        newStack := &Stack{}
//        for loops := 0; loops < 10; loops++ {
//            newVal := loops+iterations
//            if newStack.Push(newVal) {
//                if newStack.Head.Data != newVal || newStack.GetNth(0).Data != newVal {
//                    t.Errorf("Did not push to stack")
//                }
//            }
//        }
//    }
//}
//
//func TestPop(t *testing.T) {
//    for iterations := 0; iterations < 100; iterations++ {
//        newStack := &Stack{}
//        for loops := 0; loops < 10; loops++ {
//            newVal := loops+iterations
//            if newStack.Push(newVal) {
//                if newStack.Head.Data != newVal || newStack.Pop() != newVal {
//                    t.Errorf("Did not push to stack and pop right")
//                }
//            }
//        }
//    }
//}

func TestCapacity(t *testing.T) {
    newStack := new(Stack)
    otherStack := new(Stack)
    fmt.Printf("faaaaaaaa %+v\n", newStack)
    fmt.Printf("faaaaaaaa %+v\n", otherStack)
}
