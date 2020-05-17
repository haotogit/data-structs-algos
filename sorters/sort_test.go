package sorters

import (
    "fmt"
    "os"
    "bufio"
    "testing"
    "flag"
    "sort"
    "strings"
    "time"
)

const (
    document = "small-wordlist.txt"
    startQty = 5000
    incBy = 10000
    iterations = 5
    sortAlg = 0
    sortBy = ""
    sortDesc = true
)

var docFlag, sortByFlag string
var startQtyFlag, incByFlag int
var sortAlgFlag, iterationsFlag int
var sortDescFlag bool

func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

func obterItens(path string, qty int) ([]string, []string) {
    var list, proofList []string
    arquivo, err := os.Open(path)
    checkErr(err)
    defer arquivo.Close()
    escaneadora := bufio.NewScanner(arquivo)

    for i := 0; i < qty && escaneadora.Scan(); i++ {
        list = append(list, escaneadora.Text())
        proofList = append(proofList, escaneadora.Text())
    }

    checkErr(escaneadora.Err())
    for len(list) < qty {
        batch := (qty - len(list)) % (len(list)-1)
        list = append(list, list[:batch]...)
        proofList = append(proofList, proofList[:batch]...)
    }

    return list, proofList
}

func TestSort(t *testing.T) {
    fmt.Printf("Sorting %s with sortAlg: %s\n", docFlag, AlgMap[sortAlgFlag])
    for i := 0; i < iterationsFlag; i++ {
        currSort := SortererCriador(sortAlgFlag, sortByFlag, sortDescFlag)
        toSort, proof := obterItens(docFlag, startQty+(incByFlag*i))
        sort.Strings(proof)
        startTime := time.Now()
        currSort.SortIt(toSort)
        endTime := time.Now()
        elapsed := endTime.Sub(startTime)
        fmt.Printf("%d: Sorted %d items in %d ms\n", i+1, len(toSort), elapsed.Milliseconds())
        for x := 0; x < len(toSort); x++ {
            if strings.Compare(toSort[x], proof[x]) != 0 {
                t.Errorf("Failed %s sort at %d, expected %+v, but got %+v", AlgMap[sortAlgFlag], x, proof[x], toSort[x])
            }
        }

        fmt.Println("<<<<<<<<<<<<<<<<<<<<< ================== >>>>>>>>>>>>>>>>>>>>>")
    }
}

func init() {
    flag.StringVar(&docFlag, "document", document, "Document to sort")
    flag.StringVar(&sortByFlag, "sortBy", sortBy, "Sort by criteria")
    flag.IntVar(&startQtyFlag, "startQty", startQty, "Starting quantity to sort")
    flag.IntVar(&incByFlag, "incBy", incBy, "Increase each iteration by")
    flag.IntVar(&iterationsFlag, "iterations", iterations, "Number of iterations")
    flag.IntVar(&sortAlgFlag, "sortAlg", sortAlg, "Which sort, 0=bubblesort, 1=insertion, 2=selection, 3=merge, 4=quick, 5=heap")
    flag.BoolVar(&sortDescFlag, "sortDesc", sortDesc, "Sort desc|asc")
}
