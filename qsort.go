package qsort

type QSortable interface {
    LessThan(i, j int) bool
    Swap(i, j int)
    Len() int
}

func QSort(arr QSortable) {
    partition(arr, 0, arr.Len()-1)
}

func partition(arr QSortable, low, high int){
    l := high - low + 1

    if l < 2 {
        return
    }

    if l == 2 {
        if arr.LessThan(high, low) {
            arr.Swap(low, high)
        }
        return
    }

    lesserCount := 0
    for i := low+1; i < high+1; i++ {
        if arr.LessThan(i, low){
            lesserCount++
        }
    }

    for i := low+1; i < low + 1 + lesserCount; i++ {
        if arr.LessThan(i, low) {
            continue
        }

        for j := low+1+lesserCount; j < high+1; j++ {
            if arr.LessThan(j, low) {
                arr.Swap(i, j)
                break
            }
        }
    }

    arr.Swap(low, low+lesserCount)

    partition(arr, low, low+lesserCount-1)
    partition(arr, low+lesserCount+1, high)
}

