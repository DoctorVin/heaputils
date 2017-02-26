package heaputil

/* For a given index into a slice or array acting as a binary heap,
   calculate the index of the parent node. We are guaranteed of not
   being presented an index less than 3. */
func getParentIndex(cIdx int) int {
	var pIdx int
	if cIdx%2 == 1 {
		pIdx = (cIdx - 1) / 2
	} else {
		pIdx = (cIdx - 2) / 2
	}
	return pIdx
}

func siftDown(h []int, parent int) {
	sliceLen := len(h)
	leftChild := (2 * parent) + 1
	for sliceLen > leftChild {
		rightChild := leftChild + 1
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
			leftChild = (2 * parent) + 1
		}
	}
}

//MaxIntHeapify - take a slice of integers and turn it into a max heap
func MaxIntHeapify(h []int) {
	sliceLen := len(h)
	if sliceLen < 2 {
		// Either empty or a single element, this is a valid "heap"
		return
	}
	if sliceLen == 2 {
		if h[0] < h[1] {
			h[0], h[1] = h[1], h[0]
		}
		return
	}

	lastIndex := sliceLen - 1
	for idx := getParentIndex(lastIndex); -1 < idx; idx-- {
		siftDown(h, idx)
	}
}

//MaxIntHeapPush - add a value to a MaxIntHeap. The returned slice will be a
//heap regardless of whether the initial parameter slice was one at the start.
func MaxIntHeapPush(s []int, v int) []int {
	h := append(s, v)
	MaxIntHeapify(h)
	return h
}

//MaxIntHeapPop - take the max value from the heap and re-heapify. If the
//initial slice isn't a heap to start, you will get garbage out.
func MaxIntHeapPop(s []int) (int, []int) {
	cpy := make([]int, len(s), cap(s))
	copy(cpy, s)
	maxVal := cpy[0]
	lastIndex := len(cpy) - 1
	cpy[0] = cpy[lastIndex]
	cpy = cpy[:lastIndex]
	MaxIntHeapify(cpy)
	return maxVal, cpy
}

//MaxIntHeapSort - take a slice of integers and sort it in place, max value last
func MaxIntHeapSort(s []int) {
	MaxIntHeapify(s) // just in case it's not a heap when we get it.
	for idx := len(s) - 1; idx > 0; idx-- {
		s[0], s[idx] = s[idx], s[0]
		MaxIntHeapify(s[:idx])
	}
}
