package heaputil

/* For a given index into a slice or array acting as a binary heap, 
   calculate the index of the parent node. We are guaranteed of not
   being presented an index less than 3. */
func getParentIndex(cIdx int) int {
    var pIdx int
    if cIdx%2 == 1 {
        pIdx = (cIdx-1)/2
    } else {
        pIdx = (cIdx-2)/2
    }
    return pIdx
}

func siftDown(h []int, parent int) {
   sliceLen := len(h)
   leftChild := (2*parent)+1;
   for ; sliceLen > leftChild; {
       rightChild := leftChild+1
       swap := parent
       if h[swap] < h[leftChild] {
           swap = leftChild
       }
       if (rightChild < sliceLen) && (h[swap] < h[rightChild]) {
           swap = rightChild
       }
       if swap == parent {
           break
       } else {
           h[parent], h[swap] = h[swap], h[parent]
           parent = swap
           leftChild = (2*parent)+1
       }
   }
}

func MaxIntHeapify(h []int) bool {
    sliceLen := len(h)
    // early bail-out conditions
    if sliceLen == 0 {
        //stupid is as stupid does
        return false
    }
    if sliceLen == 1 {
        // it's easy to heapify a single element
        return true
    }
    if sliceLen == 2 {
        if h[0] < h[1] {
            h[0], h[1] = h[1], h[0]
        }
        return true
    }

    lastIndex := sliceLen-1
    for idx := getParentIndex(lastIndex); -1 < idx; idx-- {
        siftDown(h, idx)
    }
    return true
}
