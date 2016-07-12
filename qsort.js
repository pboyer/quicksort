function quicksort(arr){
    partition(arr, 0, arr.length-1)
}

function swap(arr, i0, i1){
    var t = arr[i0];
    arr[i0] = arr[i1];
    arr[i1] = t;
}

function partition(arr, low, high){
    var diff = high - low + 1;

    if (diff < 2) {
        return
    }

    if (diff === 2) {
        if (arr[high] < arr[low]){
            swap(arr, low, high)
            return
        }
    }

    var r = Math.floor(Math.random() * diff) + low;

    swap(arr, low, r);

    var lesserCount = 0;

    for (var i = low+1; i < high+1; i++){
        if (arr[i] < arr[low]){
            lesserCount++;
        }
    }

    for (var i = low+1; i < lesserCount + low+1; i++){
        if (arr[i] > arr[low]) {
            for (var j = lesserCount+low+1; j < high+1; j++){
                if (arr[j] < arr[low]) {
                    swap(arr, i, j);
                    break;
                }
            }
        }
    }

    swap(arr, low, low + lesserCount);

    partition(arr, low, low + lesserCount - 1);
    partition(arr, low + lesserCount + 1, high);
}