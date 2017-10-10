package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
	values := []int{5,2,4,1,3}
	BubbleSort(values)
	for i:=0;i<len(values)-2;i++{
		if values[i+1]<values[i] {
			t.Error("BubbleSort failed. Got:",values,"Expected: 1,2,3,4,5")
		}
	}
}
