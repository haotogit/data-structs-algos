package sorters

import (
    "sort"
    "strings"
    "os"
    "bufio"
    "testing"
)

func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

func obterItens(path string) ([]string, []string) {
    var list, proofList []string
    arquivo, err := os.Open(path)
    checkErr(err)
    defer arquivo.Close()
    escaneadora := bufio.NewScanner(arquivo)

    for escaneadora.Scan() {
        list = append(list, escaneadora.Text())
        proofList = append(proofList, escaneadora.Text())
    }

    checkErr(escaneadora.Err())
    return list, proofList
}

func TestBubbles(t *testing.T) {
    toSort, proof := obterItens("./small-wordlist.txt")
    Bubbler(toSort)
    sort.Strings(proof)
    for x := 0; x < len(toSort); x++ {
        if strings.Compare(toSort[x], proof[x]) != 0 {
            t.Errorf("Failed bubble sort, expected %+v, but got %+v", proof[x], toSort[x])
        }
    }
}
