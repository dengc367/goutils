package utils

import (
	"fmt"
	"testing"
	"time"
)

// TestTime test the time.Since
func TestTime(t *testing.T) {
	now := time.Now()
	et := time.Since(now)
	fmt.Println(fmt.Sprintf("elapsed time: %v", et))

}
func TestRange(t *testing.T) {
	defer func() {
		fmt.Println("defer end")

	}()
	defer func() {
		fmt.Println("defer end 2")
	}()
	a := []int{3, 4}
	for i, e := range a {
		fmt.Println(fmt.Sprintf("i, %d, e: %d", i, e))
	}

	b := 3
	for i := 0; i < b; i++ {
		println(i)
	}
}

func arrayArray() (c [][]int) {
	var a [][]int
	b := []int{3, 4, 5}
	a = append(a, b)
	c = append(c, b)
	println(fmt.Sprintf("a: %v , c: %v", a, c))
	return
}

func TestArrayArray(t *testing.T) {
	arrayArray()
	a := true
	if a == true {
		println(fmt.Sprintf("a==true?: %v", a == true))
	} else {
		println(fmt.Sprintf("a=false?: %v", a == false))
	}
}

func TestSlice(t *testing.T) {
	a := []int{3, 4, 5}
	max := len(a)
	if max > 2 {
		max = 2
	}
	println(fmt.Sprintf("arr: %v", a[:max]), MinInt(len(a), 20))
}
