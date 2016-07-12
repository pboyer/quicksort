package qsort

import (
    "testing"
    "math/rand"
)

type f64Arr []float64

func (arr f64Arr) Swap(i,j int){
    arr[j], arr[i] = arr[i], arr[j]
}

func (arr f64Arr) Len() int {
    return len(arr)
}

func (arr f64Arr) LessThan(i, j int) bool {
    return arr[i] < arr[j]
}

func TestBasicSort(t *testing.T){
    c := 10
    f := make([]float64, c)
    for i := 0; i < c; i++ {
        f[i] = rand.Float64()
    }

    QSort(f64Arr(f))

    for i := 1; i < c; i++ {
        if f[i] < f[i-1] {
            t.Fatalf("not sorted %v", f)
        }
    }
}