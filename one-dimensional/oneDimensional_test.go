package one_dimensional

import (
	"fmt"
	"testing"
)

func Test_nextNum(t *testing.T) {
	fmt.Println(getNextNum(1,100))
	fmt.Println(getNextNum(99,100))
	fmt.Println(getNextNum(100,100))
}

func Test_beforeNum(t *testing.T) {
	fmt.Println(getBeforeNum(1,100))
	fmt.Println(getBeforeNum(99,100))
	fmt.Println(getBeforeNum(0,100))
}

func Test_bytePlaceNum(t *testing.T) {
	fmt.Println(bytePlaceNum(1,0))
}

func Test_initSpace(t *testing.T) {
	space := make([]bool, 100)
	initSpace(&space,10)
	fmt.Println(space)
}