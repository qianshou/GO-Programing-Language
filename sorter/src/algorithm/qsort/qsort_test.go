package qsort

import "testing"

func TestQuickSort(t *testing.T) {
	values := []int{5,2,4,1,3}
	QuickSort(values)
	for i:=0;i<len(values)-2;i++{
		if values[i+1]<values[i] {
			t.Error("QuickSort failed. Got:",values,"Expected: 1,2,3,4,5")
		}
	}
}
