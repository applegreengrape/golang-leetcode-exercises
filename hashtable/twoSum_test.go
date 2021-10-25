package hashtable

import (
	"testing"
	"reflect"
	"log"
)

var data = []int{2,5,5,11}
var r1= []int{1,2}
var r2= []int{2,1}

func Test(t *testing.T) {
	res := TwoSum1(data, 10)
	if (!reflect.DeepEqual(res, r1) && !reflect.DeepEqual(res, r2)){
		t.Fatalf("error")
	}

	res = TwoSum2(data, 10)

	if (!reflect.DeepEqual(res, r1) && !reflect.DeepEqual(res, r2)){
		t.Fatalf("error")
	}

}

func BenchmarkTwoSum1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        res := TwoSum1(data, 10)
		if (!reflect.DeepEqual(res, r1) && !reflect.DeepEqual(res, r2)){
			log.Fatalf("error")
		}
    }
}

func BenchmarkTwoSum2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        res := TwoSum2(data, 10)
		if (!reflect.DeepEqual(res, r1) && !reflect.DeepEqual(res, r2)){
			log.Fatalf("error")
		}
    }
}