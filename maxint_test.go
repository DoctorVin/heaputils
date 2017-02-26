package heaputil

import (
	"testing"

	"reflect"
)

type testData struct {
	testCase string
	initial  []int
	final    []int
}

var heapifyTests = []testData{
	{"Empty Slice", []int{}, []int{}},
	{"Single Element", []int{0}, []int{0}},
	{"Two Elements Swap", []int{6, 27}, []int{27, 6}},
	{"Two Elements No Swap", []int{6, 4}, []int{6, 4}},
	{"Three Elements - 1", []int{1, 2, 3}, []int{3, 2, 1}},
	{"Three Elements - 2", []int{1, 3, 2}, []int{3, 1, 2}},
	{"Three Elements - 3", []int{3, 1, 2}, []int{3, 1, 2}},
	{"Seven Elements - 1", []int{2, 3, 18, 7, 4, 1, 5}, []int{18, 7, 5, 3, 4, 1, 2}},
	{"Eight Elements - 1", []int{6, 8, 14, 10, 18, 27, 5, 13}, []int{27, 18, 14, 13, 8, 6, 5, 10}},
}

func TestMaxIntHeapify(t *testing.T) {
	for _, test := range heapifyTests {
		t.Logf("Running %s", test.testCase)
		testSlice := make([]int, len(test.initial))
		copy(testSlice, test.initial)
		MaxIntHeapify(testSlice)
		if !reflect.DeepEqual(testSlice, test.final) {
			t.Errorf("%s:\n\tExpected: %v\n\tActual: %v",
				test.testCase, test.final, testSlice)
		}
	}
}

var pushTests = []struct {
	testCase string
	pushVal  int
	initial  []int
	final    []int
}{
	{"Push append", 6, []int{27, 18, 14, 13, 8, 6, 5, 10}, []int{27, 18, 14, 13, 8, 6, 5, 10, 6}},
	{"Push swap 1 level", 16, []int{27, 18, 14, 13, 8, 6, 5, 10}, []int{27, 18, 14, 16, 8, 6, 5, 10, 13}},
	{"Push swap 2 levels", 22, []int{27, 18, 14, 13, 8, 6, 5, 10}, []int{27, 22, 14, 18, 8, 6, 5, 10, 13}},
	{"Push swap new max", 33, []int{27, 18, 14, 13, 8, 6, 5, 10}, []int{33, 27, 14, 18, 8, 6, 5, 10, 13}},
}

func TestMaxIntHeapPush(t *testing.T) {
	for _, test := range pushTests {
		t.Logf("Running %s", test.testCase)
		result := MaxIntHeapPush(test.initial, test.pushVal)
		if !reflect.DeepEqual(result, test.final) {
			t.Errorf("%s:\n\tExpected: %v\n\tActual: %v",
				test.testCase, test.final, result)
		}
	}
}

var popTests = []struct {
	testCase string
	initial  []int
	popVal   int
	final    []int
}{
	{"Pop max - 1", []int{27, 18, 14, 13, 8, 6, 5, 10}, 27, []int{18, 13, 14, 10, 8, 6, 5}},
	{"Pop max - 2", []int{18, 7, 5, 3, 4, 1, 2}, 18, []int{7, 4, 5, 3, 2, 1}},
}

func TestMaxIntHeapPop(t *testing.T) {
	for _, test := range popTests {
		t.Logf("Running %s", test.testCase)
		pop, heap := MaxIntHeapPop(test.initial)
		if pop != test.popVal || !reflect.DeepEqual(heap, test.final) {
			t.Errorf("%s:\n\tExpected: %d / %v\n\tActual: %d / %v",
				test.testCase, test.popVal, test.final, pop, heap)
		}
	}
}

var sortTests = []struct {
	testCase string
	initial  []int
	final    []int
}{
	{"Sort 1", []int{18, 7, 5, 3, 4, 1, 2}, []int{1, 2, 3, 4, 5, 7, 18}},
	{"sort 2", []int{27, 18, 14, 13, 8, 6, 5, 10}, []int{5, 6, 8, 10, 13, 14, 18, 27}},
}

func TestMaxIntHeapSort(t *testing.T) {
	for _, test := range sortTests {
		t.Logf("Running %s", test.testCase)
		testSlice := make([]int, len(test.initial))
		copy(testSlice, test.initial)
		MaxIntHeapSort(testSlice)
		if !reflect.DeepEqual(testSlice, test.final) {
			t.Errorf("%s:\n\tExpected: %v\n\tActual: %v",
				test.testCase, test.final, testSlice)
		}
	}
}
