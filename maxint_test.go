package heaputil

import "testing"

type testData struct {
    testCase string
    initial []int
    final []int
    success bool
}

var tests = []testData {
    {"Single Element", []int{0}, []int{0}, true},
    {"Two Elements Swap", []int{6, 27}, []int{27, 6}, true},
    {"Two Elements No Swap", []int{6, 4}, []int{6, 4}, true},
    {"Three Elements - 1", []int{1, 2, 3}, []int{3, 2, 1}, true},
    {"Three Elements - 2", []int{1, 3, 2}, []int{3, 1, 2}, true},
    {"Three Elements - 3", []int{3, 1, 2}, []int{3, 1, 2}, true},
    {"Seven Elements - 1", []int{2,3,18,7,4,1,5}, []int{18,7,5,3,4,1,2}, true},
}

func TestEmptySlice(t *testing.T){
    mySlice := make([]int, 0)
    success := MaxIntHeapify(mySlice)
    if success {
        t.Error("Heapify didn't fail with an empty slice.")
    }
}

func sliceEquality(lhs, rhs []int) bool {
    lenEq := false
    valEq := true
    if (len(lhs) == len(rhs)) {
        lenEq = true
        for i, _ := range lhs {
            if lhs[i] != rhs[i] {
                valEq = false
                break
            }
        }
    }
    return lenEq && valEq
}

func TestHeapify(t *testing.T){
    for _, test := range tests {
        t.Logf("Running %s", test.testCase)
        testSlice := make([]int, len(test.initial))
        copy(testSlice, test.initial)
        res := MaxIntHeapify(testSlice)
        if test.success != res {
            t.Errorf("Test case %s returns an error.", test.testCase)
        }
        if ! sliceEquality(testSlice, test.final) {
            t.Errorf("Test case %s:\n\tExpected: %v\n\tActual: %v")
        }
    }
}

