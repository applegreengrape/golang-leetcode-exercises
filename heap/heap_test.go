package heap

import (
	"testing"
	"fmt"
)

func TestMinHeap(t *testing.T){
	min := &IntHeap{6,3,5,3,5,2,8}

	h := minHeap(*min, 1)
	
	fmt.Println(h)

	if (h[0] != 1){
		t.Fatalf("err")
	}
}

func TestMaxHeapFunc(t *testing.T){
	max := &MaxHeap{[]int{3,5,46,75,46546,3423,3424}}
	
	h := maxHeap(*max, 5555)

	fmt.Println(h)

	if (h.IntHeap[0] != 46546 ){
		t.Fatalf("err")
	}
}