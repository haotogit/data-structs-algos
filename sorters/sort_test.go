package sorters

import (
    "fmt"
    "testing"
    "flag"
    "sort"
    "reflect"
    "time"
    "math/rand"
    "../util"
)

const (
    document = "random-strings.txt"
    startQty = 5000
    incBy = 10000
    iterations = 5
    sortAlg = 4
    sortBy = ""
    sortDesc = true
    sortType = "string"
)

var docFlag, sortByFlag, typeFlag string
var startQtyFlag, incByFlag int
var sortAlgFlag, iterationsFlag int
var sortDescFlag bool

//func checkErr(e error) {
//    if e != nil {
//        panic(e)
//    }
//}

func obterItens(path string, qty int) ([]interface{}, []interface{}) {
    var list, proofList []interface{}
    
    if typeFlag == "string" {
        for qty > 0 {
            currLen := util.GetRandIntn(1, 12)
            bytes := make([]byte, currLen)
            for i := 0; i < currLen; i++ {
                bytes[i] = byte(util.GetRandIntn(48, 122))
            }

            currStr := string(bytes)
            list = append(list, currStr)
            qty--
        }
        //arquivo, err := os.Open(path)
        //checkErr(err)
        //defer arquivo.Close()
        //// TODO refactor this to read without bufio
        //escaneadora := bufio.NewScanner(arquivo)

        //for i := 0; i < qty && escaneadora.Scan(); i++ {
        //    list = append(list, escaneadora.Text())
        //    proofList = append(proofList, escaneadora.Text())
        //}

        //checkErr(escaneadora.Err())
        //for len(list) < qty {
        //    batch := (qty - len(list)) % (len(list)-1)
        //    list = append(list, list[:batch]...)
        //    proofList = append(proofList, proofList[:batch]...)
        //}
    } else {
        rand.Seed(time.Now().UnixNano())
        tmp := rand.Perm(qty)
        for i := 0; i < qty; i++ {
            list = append(list, tmp[i]*i+1)
        }
    }

    proofList = append(proofList, list...)
    return list, proofList
}

func TestSort(t *testing.T) {
    fmt.Printf("Sorting %ss with sortAlg: %s\n", typeFlag, AlgMap[sortAlgFlag])
    for i := 0; i < iterationsFlag; i++ {
        currSort := SortererCriador(sortAlgFlag)
        toSort, proof := obterItens(docFlag, startQty+(incByFlag*i))
        sort.Slice(proof, func(i, j int) bool { return util.ElComparer(proof[i], proof[j]) == -1 })
        startTime := time.Now()
        currSort.SortIt(toSort)
        endTime := time.Now()
        elapsed := endTime.Sub(startTime)
        if !reflect.DeepEqual(toSort, proof) {
            t.Errorf("Did not sort correctly")
        }
        fmt.Printf("%d: Sorted %d items in %d ms\n", i+1, len(toSort), elapsed.Milliseconds())
        fmt.Println("<<<<<<<<<<<<<<<<<<<<< ================== >>>>>>>>>>>>>>>>>>>>>")
    }
}

func BenchmarkSort(b *testing.B) {
    b.StopTimer()
    for i := 0; i < iterationsFlag; i++ {
        currSort := SortererCriador(sortAlgFlag)
        toSort, proof := obterItens(docFlag, startQty+(incByFlag*i))
        sort.Slice(proof, func(i, j int) bool { return util.ElComparer(proof[i], proof[j]) == -1 })
        b.StartTimer()
        currSort.SortIt(toSort)
        b.StopTimer()
        if !reflect.DeepEqual(toSort, proof) {
            b.Errorf("Did not sort correctly")
        }
    }
}

func init() {
    //flag.StringVar(&docFlag, "doc", document, "Document to sort")
    flag.StringVar(&sortByFlag, "sortBy", sortBy, "Sort by criteria")
    flag.StringVar(&typeFlag, "t", sortType, "Sort typeof el")
    flag.IntVar(&startQtyFlag, "startQty", startQty, "Starting quantity to sort")
    flag.IntVar(&incByFlag, "incBy", incBy, "Increase each iteration by")
    flag.IntVar(&iterationsFlag, "iter", iterations, "Number of iterations")
    flag.IntVar(&sortAlgFlag, "sortAlg", sortAlg, "Which sort, 0=bubblesort, 1=selection, 2=insertion, 3=merge, 4=quick, 5=heap")
    flag.BoolVar(&sortDescFlag, "sortDesc", sortDesc, "Sort desc|asc")
}
