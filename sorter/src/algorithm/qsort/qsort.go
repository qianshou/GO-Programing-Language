package qsort

func Partition(values []int,left int,right int) int {
	pivokey := values[left]
	for left<right {
		for left<right && values[right]>=pivokey {
			right--
		}
		values[left] = values[right]
		for left<right && values[left]<=pivokey {
			left++
		}
		values[right] = values[left]
	}
	values[left] = pivokey
	return  left
}

func qSort(values []int,left int,right int){
	if left<right {
		pivotloc := Partition(values,left,right)
		qSort(values,left,pivotloc-1)
		qSort(values,pivotloc+1,right)
	}
}

func QuickSort(values []int)  {
	qSort(values,0,len(values)-1)
}
