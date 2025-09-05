package p0875_koko_eating_bananas

import (
	"fmt"
	"testing"
)

func TestUsolve(t *testing.T) {
	//piles := []int{3, 6, 7, 11}
	//h := 8
	//ans := minEatingSpeed(piles, h)
	//fmt.Println(ans)
	//piles = []int{30, 11, 23, 4, 20}
	//h = 5
	//ans = minEatingSpeed(piles, h)
	//fmt.Println(ans)
	piles := []int{312884470}
	h := 312884469
	ans := minEatingSpeed(piles, h)
	fmt.Println(ans)
	// Add test cases here
	// Example:
	// result := Usolve(/* input */)
	// expected := /* expected */
	// if result != expected {
	//     t.Errorf("Expected %v, got %v", expected, result)
	// }
}
