package graphalgo

import (
	"container/list"
	"fmt"
)

func queue(){
	vals := []int{1,3,5,2,4,3}
	q := list.New()

	for _, v := range vals {
		q.PushBack(v)
	}

	fmt.Println(q.Front().Value)

	for val := q.Front(); val != nil; val = val.Next(){
		fmt.Println(val.Value)
	}

}